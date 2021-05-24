package controllers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type EchoController struct {
}

type ReqInfo struct {
	Headers    map[string]string `json:"headers"`
	Method     string            `json:"method"`
	Params     gin.Params        `json:"params"`
	RemoteAddr string            `json:"remote_addr"`
	Form       map[string]string `json:"form"`
	Data       string            `json:"data"`
}

func (t *EchoController) Index(c *gin.Context) {
	c.Request.ParseMultipartForm(32 << 20)
	r := ReqInfo{
		Headers:    map[string]string{},
		Form:       map[string]string{},
		Params:     c.Params,
		RemoteAddr: c.Request.RemoteAddr,
		Method:     c.Request.Method,
	}
	for k, v := range c.Request.Header {
		r.Headers[k] = v[0]
	}
	for k, v := range c.Request.Form {
		r.Form[k] = v[0]
	}
	if body, err := ioutil.ReadAll(c.Request.Body); err != nil {
		println(err.Error())
	} else {
		r.Data = string(body)
	}

	c.JSON(http.StatusOK, r)
}
