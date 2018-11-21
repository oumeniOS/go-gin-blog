package models

type Article struct {
	Model

	TagID int `json:"tag_id" grom:"index"`
	Tag   Tag `json:"tag"`

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
	//DeletedAt  string    `json:"deleted_at"`
}

/*
获取文章列表：GET("/articles")
获取指定文章：POST("/articles/:id")
新建文章：POST("/articles")
更新指定文章：PUT("/articles/:id")
删除指定文章：DELETE("/articles/:id")
*/

//文章是否存在
func IsArticleExistById(id int) (bool, error) {
	var article Article
	db.Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetArticleTotal(maps interface{}) (count int, err error) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return count, nil
}

//获取文章列表
func ArticleList(pageNum int, pageSize int, maps map[string]interface{}) (articles []*Article, err error) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return articles, nil
}

//获取指定文章
func GetArticle(id int) (article *Article, err error) {

	db.Preload("Tag").Unscoped().Where("id=6").First(&article)
	return article,nil
}

//更新指定文章
func EditArticle(id int, maps map[string]interface{}) error {
	var artical Article
	db.Where("id=?", id).First(&artical)
	db.Model(&artical).Update(maps)
	//db.Model(&Article{}).Where("id=?", id).Updates(maps)
	return nil
}

//删除指定文章  软删除
func DeleteArticle(id int) error {
	db.Where("id=?", id).Delete(&Article{})
	return nil
}

//新建文章
func AddArticle(maps map[string]interface{}) error {
	db.Create(&Article{
		TagID:         maps["tag_id"].(int),
		Title:         maps["title"].(string),
		Desc:          maps["desc"].(string),
		Content:       maps["content"].(string),
		CoverImageUrl: maps["cover_image_url"].(string),
		CreatedBy:     maps["created_by"].(string),
		State:         maps["state"].(int),
	})
	return nil
}

//硬删除文章
func ClearAllArtical() bool {
	db.Unscoped().Where("deleted_at != ?", 0).Delete(&Article{})
	return true
}
