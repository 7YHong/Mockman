package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"mockman/di"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func RespLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		reqInfo := getReqInfo(c)
		c.Set("reqInfo", &reqInfo)
		reqJson, _ := json.MarshalIndent(reqInfo, "", "    ")

		c.Next()

		di.Zap().Infof("Send HTTP response, req uri: %s, \nreqInfo: %v, \nresp code: %v, \nrespbody: %v",
			c.Request.RequestURI, string(reqJson), c.Writer.Status(), blw.body.String())
	}
}
