package models

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{})(tags []*Tag, err error)  {
	//查询条件。偏移量。页面大小。查询结果
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return tags, nil
}

func GetTagTotal(maps interface{})(count int, err error)  {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return count, nil
}

func EditTag(id int, maps map[string]interface{}) error {
	var tag Tag
	db.Where("id=?",id).First(&tag)
	db.Model(&tag).Update(maps)
	return nil
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
func ExistTagById(id int) (bool, error) {
	var tag Tag
	db.Select("id").Where("id=?",id).First(&tag)
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

//删除标签
func DeleteTagById(id int) error {
	db.Where("id=?",id).Delete(&Tag{})
	return nil
}

//添加标签
func AddTag(name string, state int, createdBy string) error {
	db.Create(&Tag{
		Name:name,
		State:state,
		CreatedBy:createdBy,
	})
	return nil
}


//硬删除标签
func ClearAllTag() bool  {

	db.Unscoped().Where("deleted_at != ?", 0).Delete(&Tag{})
	return true
}












































