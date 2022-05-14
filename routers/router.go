package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/saxonredhat/go-gin-example/pkg/setting"
    "github.com/saxonredhat/go-gin-example/routers/v2"
)

func InitRouter() *gin.Engine {
    r := gin.New()

    r.Use(gin.Logger())

    r.Use(gin.Recovery())

    gin.SetMode(setting.RunMode)

    apiv2 := r.Group("/api/v2")
    {
        //获取标签列表
        apiv2.GET("/tags", v2.GetTags)
        //新建标签
        apiv2.POST("/tags", v2.AddTag)
        //更新指定标签
        apiv2.PUT("/tags/:id", v2.EditTag)
        //删除指定标签
        apiv2.DELETE("/tags/:id", v2.DeleteTag)
    }


    r.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "test",
        })
    })

    return r
}
