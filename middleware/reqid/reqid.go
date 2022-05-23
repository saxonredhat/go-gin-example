package reqid 

import (
    "github.com/gin-gonic/gin"
    "github.com/gofrs/uuid"
)


func REQID() gin.HandlerFunc {
    return func(c *gin.Context){
        reqId := c.Request.Header.Get("X-Request-Id")
        if reqId == "" {
            u, err := uuid.NewV4()
            if err == nil {
                reqId = u.String()
                c.Request.Header.Set("X-Request-Id", reqId)
            }
        }
        c.Header("X-Request-Id", reqId)
        c.Next()
    }
}
