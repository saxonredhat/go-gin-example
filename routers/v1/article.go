package v1

import (
    "fmt"
    "strings"
    "net/http"

    "github.com/saxonredhat/go-gin-example/models"
    "github.com/saxonredhat/go-gin-example/pkg/e"
    "github.com/saxonredhat/go-gin-example/pkg/setting"
    "github.com/saxonredhat/go-gin-example/pkg/util"
    "github.com/saxonredhat/go-gin-example/pkg/logging"

    "github.com/gin-gonic/gin"
    "github.com/astaxie/beego/validation"
    "github.com/unknwon/com"
)

// @Summary 获取文章列表 
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [get]
func GetArticles(c *gin.Context) {
    name := c.Query("name")

    maps := make(map[string]interface{})
    data := make(map[string]interface{})

    if name != "" {
        maps["name"] = name
    }

    var state int = -1
    if arg := c.Query("state"); arg != "" {
        state = com.StrTo(arg).MustInt()
        maps["state"] = state
    }

    code := e.SUCCESS

    data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
    data["total"] = models.GetArticleTotal(maps)

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}

// @Summary 新增文章
// @Produce  json
// @Param tag_id query int true "TagID"
// @Param created_by query string true "CreatedBy"
// @Param title query string true "Title"
// @Param desc query string false "Desc"
// @Param content query string false "Content"
// @Param state query int false "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [post]
func AddArticle(c *gin.Context) {
    //获取参数
    tagId := c.Query("tag_id")
    title := c.Query("title")
    desc := c.Query("desc")
    content := c.Query("content")
    createdBy := c.Query("created_by")

    //校验参数
    valid := validation.Validation{}

    tagIdInt := -1
    if tagId != ""  {
        tagIdInt = com.StrTo(tagId).MustInt()
    }

    var state int = -1
    if arg := c.Query("state"); arg != "" {
        state = com.StrTo(arg).MustInt()
        valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
    }

    valid.Required(tagId, "tagId").Message("TagID不能为空")
    valid.Required(title, "title").Message("title不能为空")
    valid.MaxSize(title, 100, "name").Message("名称最长为100字符")
    valid.Required(createdBy, "created_by").Message("创建人不能为空")
    valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")

    code := e.INVALID_PARAMS
    msg := ""
    if ! valid.HasErrors() {
        data := make(map[string]interface{})
        data["tag_id"] = tagIdInt 
        data["title"] = title
        data["desc"] = desc 
        data["content"] = content
        data["created_by"] = createdBy 
        data["state"] = state
        //判断tag_id是否存在
        if models.ExistTagByID(tagIdInt){
            if models.AddArticle(data){
                code = e.SUCCESS
            }else{
                code = e.ERROR_DML
            }
        }else{
            code = e.ERROR_NOT_EXIST_TAG
        }
        msg = e.GetMsg(code)
    }else{
        var errors strings.Builder
        errors.WriteString(fmt.Sprintf("%s,",e.GetMsg(code)))
        for idx, err := range valid.Errors {
            //打印日志
            logging.Info(" controller: %s: %s",err.Key, err.Message)
            if idx == len(valid.Errors)-1{
                errors.WriteString(fmt.Sprintf("%s: %s。",err.Key,err.Message))
            }else{
                errors.WriteString(fmt.Sprintf("%s: %s，",err.Key,err.Message))
            }
        }
        msg = errors.String() 
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : msg,
        "data" : make(map[string]string),
    })
}

// @Summary 修改文章
// @Produce  json
// @Param modified_by query string true "ModifiedBy"
// @Param tag_id query int false "TagID"
// @Param title query string false "Title"
// @Param desc query string false "Desc"
// @Param content query string false "Content"
// @Param state query int false "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles/{id} [put]
func EditArticle(c *gin.Context) {
    //获取参数
    id := com.StrTo(c.Param("id")).MustInt()
    tagId := c.Query("tag_id")
    title := c.Query("title")
    desc := c.Query("desc")
    content := c.Query("content")
    modifiedBy := c.Query("modified_by")

    //校验参数
    valid := validation.Validation{}
    var state int = -1
    if arg := c.Query("state"); arg != "" {
        state = com.StrTo(arg).MustInt()
        valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
    }
    tagIdInt := -1
    if tagId != ""  {
        tagIdInt = com.StrTo(tagId).MustInt()
    }

    valid.Required(id, "id").Message("ID不能为空")
    valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
    valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
    valid.MaxSize(title, 100, "title").Message("title最长为100字符")
    valid.MaxSize(desc, 255, "desc").Message("desc最长为255字符")

    //处理请求参数
    code := e.INVALID_PARAMS
    msg := ""
    data := make(map[string]interface{})
    if ! valid.HasErrors() {
        if models.ExistArticleByID(id) {
            if models.ExistTagByID(tagIdInt){
                data["tag_id"] = tagIdInt
                data["modified_by"] = modifiedBy
                if title != "" {
                    data["title"] = title
                }
                if desc != "" {
                    data["desc"] = desc 
                }
                if content != "" {
                    data["content"] = content
                }
                if state != -1 {
                    data["state"] = state
                }
                if models.EditArticle(id, data){
                    code = e.SUCCESS
                }else{
                    code = e.ERROR_DML
                }
            }else{
                code = e.ERROR_NOT_EXIST_TAG
            }
        } else {
            code = e.ERROR_NOT_EXIST_ARTICLE
        }
        msg = e.GetMsg(code)
    }else{
        var errors strings.Builder
        errors.WriteString(fmt.Sprintf("%s,",e.GetMsg(code)))
        for idx, err := range valid.Errors {
            if idx == len(valid.Errors)-1{
                errors.WriteString(fmt.Sprintf("%s: %s.",err.Key,err.Message))
            }else{
                errors.WriteString(fmt.Sprintf("%s: %s,",err.Key,err.Message))
            }
        }
        msg = errors.String() 
    }

    logging.Info(fmt.Sprintf(" controller: %d %s %v", code, msg, data))
    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : msg,
        "data" : data,
    })
}

// @Summary 删除文章
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles/{id} [delete]
func DeleteArticle(c *gin.Context) {
    //获取参数
    id := c.Param("id")
    //校验参数
    valid := validation.Validation{}
    idInt := -1
    if id != "" {
        idInt = com.StrTo(id).MustInt()
    }
    valid.Required(id, "id").Message("ID不能为空")
    //数据库校验
    var code int
    if models.ExistArticleByID(idInt){
        if models.DeleteArticle(idInt){
           code = e.SUCCESS
        }else{
           code = e.ERROR_DML 
        }
    }else{
        code = e.ERROR_NOT_EXIST_ARTICLE
    }
    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : make(map[string]interface{}),
    })
}
