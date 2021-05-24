package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type ReqInfo struct {
	Headers    map[string]string `json:"headers"`
	Method     string            `json:"method"`
	Params     gin.Params        `json:"params"`
	RemoteAddr string            `json:"remote_addr"`
	Form       map[string]string `json:"form"`
	Data       string            `json:"data"`
}

func getReqInfo(c *gin.Context) ReqInfo {
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
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	r.Data = string(bodyBytes)

	return r
}
