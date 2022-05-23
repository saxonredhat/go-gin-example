package log 

import (
    "fmt"
    "github.com/gin-gonic/gin"

    "github.com/saxonredhat/go-gin-example/pkg/logging" 
)

func LOG() gin.HandlerFunc {
    return func(c *gin.Context) {
        reqId := c.Request.Header.Get("X-Request-Id")
        msg := fmt.Sprintf("reqid:%s method:%s url:%s proto:%s", reqId, c.Request.Method, c.Request.URL, c.Request.Proto)
        logging.Info(msg)
        c.Next()
    }
}
