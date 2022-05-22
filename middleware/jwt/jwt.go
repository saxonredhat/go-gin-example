package jwt

import (
    "time"
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/saxonredhat/go-gin-example/pkg/util"
    "github.com/saxonredhat/go-gin-example/pkg/e"
)

func JWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        var code int
        data := make(map[string]interface{})

        code = e.SUCCESS
        token := c.Query("token")
        if token == "" {
            code = e.ERROR_AUTH_NOT_EXIST_TOKEN
        } else {
            claims, err := util.ParseToken(token)
            if err != nil {
                code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
            } else if time.Now().Unix() > claims.ExpiresAt {
                code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
            }
        }

        if code != e.SUCCESS {
            c.JSON(http.StatusUnauthorized, gin.H{
                "code" : code,
                "msg" : e.GetMsg(code),
                "data" : data,
            })

            c.Abort()
            return
        }

        c.Next()
    }
}
