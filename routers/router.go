package routers

import (
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"

    "github.com/saxonredhat/go-gin-example/pkg/setting"
    "github.com/saxonredhat/go-gin-example/routers/api"
    "github.com/saxonredhat/go-gin-example/routers/v1"
    "github.com/saxonredhat/go-gin-example/middleware/jwt"
    "github.com/saxonredhat/go-gin-example/middleware/log"
    "github.com/saxonredhat/go-gin-example/middleware/reqid"
    _ "github.com/saxonredhat/go-gin-example/docs" 
)

func InitRouter() *gin.Engine {
    r := gin.New()

    r.Use(reqid.REQID())
    r.Use(gin.Logger())
    r.Use(log.LOG())

    r.Use(gin.Recovery())

    gin.SetMode(setting.RunMode)

    r.GET("/auth", api.GetAuth)
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    apiv1 := r.Group("/api/v1")
    //apiv1.Use(log.LOG())
    apiv1.Use(jwt.JWT())
    {
        //获取标签列表
        apiv1.GET("/tags", v1.GetTags)
        //新建标签
        apiv1.POST("/tags", v1.AddTag)
        //更新指定标签
        apiv1.PUT("/tags/:id", v1.EditTag)
        //删除指定标签
        apiv1.DELETE("/tags/:id", v1.DeleteTag)
        //获取文章列表
        apiv1.GET("/articles", v1.GetArticles)
        //新建文章
        apiv1.POST("/articles", v1.AddArticle)
        //更新指定文章
        apiv1.PUT("/articles/:id", v1.EditArticle)
        //删除指定文章
        apiv1.DELETE("/articles/:id", v1.DeleteArticle)
    }


    r.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "test",
        })
    })

    fmt.Printf("\nfmt\n")
    return r
}
