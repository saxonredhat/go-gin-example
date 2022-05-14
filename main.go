package main

import (
    "fmt"
    "net/http"
    //"github.com/gin-gonic/gin"
    "github.com/saxonredhat/go-gin-example/pkg/setting"
    "github.com/saxonredhat/go-gin-example/routers"
)

func main() {
    router := routers.InitRouter()
    s := &http.Server{
        Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
        Handler:        router,
        ReadTimeout:    setting.ReadTimeout,
        WriteTimeout:   setting.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
    }

    fmt.Printf("main")
    fmt.Printf("main2")
    fmt.Printf("main3")
    fmt.Printf("main4")
    fmt.Printf("main5")
    fmt.Printf("test")
    fmt.Printf("test2")
    s.ListenAndServe()
}
