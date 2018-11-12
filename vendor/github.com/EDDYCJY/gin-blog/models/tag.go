package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{})(tags []Tag)  {
	//查询条件。偏移量。页面大小。查询结果
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{})(count int)  {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func EditTag(id int, maps map[string]interface{}) bool {
	var tag Tag
	db.Where("id=?",id).First(&tag)
	db.Model(&tag).Update(maps)
	return true
}

//是否已存在标签
func ExistTagByName(name string) bool  {
	var tag Tag
	db.Select("id").Where("name=?",name).First(&tag)
	if tag.ID > 0{
		return true
	}
	return false
}

//是否已存在标签
func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id=?",id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

//删除标签
func DeleteTagById(id int) bool {
	db.Where("id=?",id).Delete(&Tag{})
	return true
}

//添加标签
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:name,
		State:state,
		CreatedBy:createdBy,
	})
	return true
}

func (tag *Tag)BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn",time.Now().Unix())
	return nil
}

func (tag *Tag)BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn",time.Now().Unix())
	return nil
}













































