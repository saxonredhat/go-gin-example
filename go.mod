module github.com/saxonredhat/go-gin-example

go 1.16

require (
	github.com/astaxie/beego v1.12.3
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.7.7
	github.com/go-ini/ini v1.66.4
	github.com/go-openapi/spec v0.20.6 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/swaggo/gin-swagger v1.4.3 // indirect
	github.com/swaggo/swag v1.8.2 // indirect
	github.com/ugorji/go v1.2.7 // indirect
	github.com/unknwon/com v1.0.1
	github.com/urfave/cli/v2 v2.8.0 // indirect
	golang.org/x/crypto v0.0.0-20220518034528-6f7dac969898 // indirect
	golang.org/x/net v0.0.0-20220520000938-2e3eb7b945c2 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/saxonredhat/go-gin-example/conf => /home/gopath/src/github.com/saxonredhat/go-gin-example/pkg/conf
	github.com/saxonredhat/go-gin-example/middleware => /home/gopath/src/github.com/saxonredhat/go-gin-example/middleware
	github.com/saxonredhat/go-gin-example/models => /home/gopath/src/github.com/saxonredhat/go-gin-example/models
	github.com/saxonredhat/go-gin-example/pkg/setting => /home/gopath/src/github.com/saxonredhat/go-gin-example/pkg/setting
	github.com/saxonredhat/go-gin-example/routers => /home/gopath/src/github.com/saxonredhat/go-gin-example/routers
//	github.com/saxonredhat/go-gin-example/routers/v2 => /home/gopath/src/github.com/saxonredhat/go-gin-example/routers/v2
)
