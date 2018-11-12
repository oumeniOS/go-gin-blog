package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model

	TagID int `json:"tag_id" grom:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

/*
获取文章列表：GET("/articles")
获取指定文章：POST("/articles/:id")
新建文章：POST("/articles")
更新指定文章：PUT("/articles/:id")
删除指定文章：DELETE("/articles/:id")
*/

func BeforeCreate(scope gorm.Scope) error {
	scope.SetColumn("CreatedOn",time.Now().Unix())
	return nil
}

func BeforeUpdate(scope gorm.Scope) error {
	scope.SetColumn("ModifiedOn",time.Now().Unix())
	return nil
}

//文章是否存在
func IsArticleExistById(id int) bool {
	var article Article
	db.Select("id").Where("id=?",id).First(&article)
	if article.ID > 0 {
		return true
	}

	return false
}

func GetArticleTotal(maps interface{})(count int)  {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

//获取文章列表
func ArticleList(pageNum int, pageSize int, maps map[string]interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

//获取指定文章
func GetArticle(id int) (article Article) {
	//db.Preload("Tag").Where("id=?",id).First(&article)
	db.Where("id=?",id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

//更新指定文章
func EditArticle(id int, maps map[string]interface{}) bool {
	db.Model(&Article{}).Where("id=?",id).Updates(maps)
	return true
}

//删除指定文章
func DeleteArticle(id int) bool {
	db.Delete(&Article{},"id=?",id)
	return true
}

//新建文章
func NewArticle(maps map[string]interface{}) bool {
	db.Create(&Article{
		TagID:maps["tag_id"].(int),
		Title:maps["title"].(string),
		Desc:maps["desc"].(string),
		Content:maps["content"].(string),
		CreatedBy:maps["created_by"].(string),
		State:maps["state"].(int),
	})
	return true
}















































