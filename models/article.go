package models 

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "time"
)

type Article struct{
    Model
    TagID int `json:"tag_id" gorm:"index"`
    Tag Tag `json:"tag"`
    Title string `json:"title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
    CreatedBy string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    DeletedOn int `json:"deleted_on"`
    State int `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedOn",time.Now().Unix()) 
    return nil 
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("ModifiedOn",time.Now().Unix()) 
    return nil 
}

func GetArticleTotal(maps interface{})(count int){
    db.Model(&Article{}).Where(maps).Count(&count)
    return
}

func GetArticles(pageNum int, pageSize int,maps interface{})(articles []Article){
    db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
    return
}

func AddArticle(data map[string]interface{}) bool{
    err := db.Create(&Article{
        TagID : data["tag_id"].(int),
        Title : data["title"].(string),
        Desc : data["desc"].(string),
        Content : data["content"].(string),
        CreatedBy : data["created_by"].(string),
        State : data["state"].(int),
    }).Error
    fmt.Printf("%+v\n", err)
    if err != nil {
        return false 
    }
    return true
}

func EditArticle(id int, data interface{}) bool{
    if err := db.Model(&Article{}).Where("id = ?", id).Updates(data).Error; err != nil {
        return false 
    }
    return true
}

func ExistArticleByID(id int) bool{
    //查询数据库是否存在
    var article Article
    db.Select("id").Where("id = ?",id).First(&article)
    if article.ID > 0 {
        return true
    }
    return false 
}

func DeleteArticle(id int) bool{
    if err := db.Where("id = ?", id).Delete(&Article{}).Error; err != nil {
        return false
    }
    return true
}

