module github.com/saxonredhat/go-gin-example

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-ini/ini v1.66.4
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/ugorji/go v1.2.7 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20220511200225-c6db032c6c88 // indirect
	golang.org/x/sys v0.0.0-20220503163025-988cb79eb6c6 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/saxonredhat/go-gin-example/conf => /home/gopath/src/github.com/saxonredhat/go-gin-example/pkg/conf
	github.com/saxonredhat/go-gin-example/middleware => /home/gopath/src/github.com/saxonredhat/go-gin-example/middleware
	github.com/saxonredhat/go-gin-example/models => /home/gopath/src/github.com/saxonredhat/go-gin-example/models
	github.com/saxonredhat/go-gin-example/pkg/setting => /home/gopath/src/github.com/saxonredhat/go-gin-example/pkg/setting
	github.com/saxonredhat/go-gin-example/routers => /home/gopath/src/github.com/saxonredhat/go-gin-example/routers
//	github.com/saxonredhat/go-gin-example/routers/v1 => /home/gopath/src/github.com/saxonredhat/go-gin-example/routers/v1
)
