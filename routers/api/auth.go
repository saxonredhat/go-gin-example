package api 
import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/saxonredhat/go-gin-example/pkg/e"
    "github.com/saxonredhat/go-gin-example/pkg/util"
    "github.com/saxonredhat/go-gin-example/pkg/logging"
    "github.com/saxonredhat/go-gin-example/models"
    "github.com/astaxie/beego/validation"
    _ "github.com/unknwon/com"
)

func GetAuth(c *gin.Context){
    //获取参数
    username := c.Query("username")
    password := c.Query("password")
    //校验参数
    valid := validation.Validation{}
    valid.Required(username, "username").Message("username不能为空")
    valid.Required(password, "password").Message("password不能为空")
    //校验用户名和密码
    var code int
    data := make(map[string]interface{})
    if ! valid.HasErrors(){
        if models.CheckAuth(username, password){
            //生成token
            token, err := util.GenerateToken(username, password)
            fmt.Printf("%v",err)
            if err != nil {
                code = e.ERROR_AUTH_TOKEN
            }else{
                data["token"] = token
                code = e.SUCCESS
            }
        }else{
            code = e.ERROR_AUTH
        }
    }else{
        code = e.INVALID_PARAMS
    }
    //返回json
    msg := e.GetMsg(code)
    logging.Info(fmt.Sprintf(" controller: %d %s %v", code, msg, data))
    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : msg,
        "data" : data,
    })
}
