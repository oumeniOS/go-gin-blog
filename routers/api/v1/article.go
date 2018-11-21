package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/oumeniOS/go-gin-blog/pkg/e"
		"net/http"
		"github.com/oumeniOS/go-gin-blog/pkg/app"
	"github.com/oumeniOS/go-gin-blog/service/article_service"
	"github.com/oumeniOS/go-gin-blog/service/tag_service"
)

/*
获取文章列表：GET("/articles")
获取指定文章：POST("/articles/:id")
新建文章：POST("/articles")
更新指定文章：PUT("/articles/:id")
删除指定文章：DELETE("/articles/:id")
*/

//获取文章列表
func GetArticleList(c *gin.Context) {
	//筛选条件只允许有state、tagid
	//查询参数变量
	appG := app.Gin{c}
	maps := make(map[string]interface{})
	//存放结果的变量
	data := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	//state范围限定
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许为0或1")
	}

	//tadid范围限定
	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{
		State: state,
		TagID: tagId,
	}

	list, err := articleService.GetAll()
	if err != nil{
		appG.Response(http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}
	data["list"] = list

	total, err := articleService.Count()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}
	data["total"] = total

	appG.Response(http.StatusOK,e.SUCCESS,data)
}

//获取指定文章
func GetArticle(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Query("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id 不能小于1")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIT_ARTICLE, nil)
		return
	}

	article, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, article)
}

//新建文章
func AddArticle(c *gin.Context) {
	appG := app.Gin{c}
	var title string = ""
	title = c.Request.URL.Query().Get("title")
	title1 := c.Query("title")
	println(title1)
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title = c.Query("title") //("title")
	content := c.Query("content")
	desc := c.Query("desc")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.Query("state")).MustInt()
	coverImageUrl:= c.Query("cover_image_url")

	valid := validation.Validation{}

	valid.Required(tagId, "tag_id").Message("标签ID不能问空")
	valid.Min(tagId, 1, "tag_id").Message("标签ID最小为1")
	valid.Required(title, "title").Message("文章标题不能为空")
	valid.Required(content, "content").Message("文章内容不能为空")
	valid.Required(desc, "desc").Message("文章描述不能为空")
	valid.Required(createdBy, "created_by").Message("文章创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("文章状态只能为0或1")

	articleService := article_service.Article{
		Title:title,
		TagID:tagId,
		Content:content,
		CreatedBy:createdBy,
		Desc:desc,
		CoverImageUrl:coverImageUrl,
		State:state,
	}

	tagService := tag_service.Tag{ID:tagId}
	exists, err := tagService.Exist()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_CHECK_EXIST_TAG_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIT_TAG, nil)
		return
	}

	err = articleService.Add()
	if err != nil {
		appG.Response(http.StatusOK,e.ERROR_ADD_ARTICLE_FAIL,nil)
		return
	}

	appG.Response(http.StatusOK,e.SUCCESS,make(map[string]interface{}))
}

//更新指定文章
func EditArticle(c *gin.Context) {

	appG := app.Gin{c}
	maps := make(map[string]interface{})
	id := com.StrTo(c.Query("id")).MustInt()

	title := c.Query("title")
	content := c.Query("content")
	desc := c.Query("desc")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("文章ID不能为空")
	valid.Min(id, 1, "id").Message("文章ID不能小于1")
	valid.Required(modifiedBy, "modified_by").Message("文章修改人不能为空")
	maps["id"] = id
	maps["modified_by"] = modifiedBy

	if title != "" {
		maps["title"] = title
	}
	if content != "" {
		maps["content"] = content
	}
	if desc != "" {
		maps["desc"] = desc
	}
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("文章状态只能为0或1")
		maps["state"] = state
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK,e.INVALID_PARAMS,nil)
		return
	}

	articleService := article_service.Article{
		ID:id,
		Title:title,
		Content:content,
		Desc:desc,
		ModifiedBy:modifiedBy,
		State:state,
	}

	exist, err := articleService.ExistByID()
	if err != nil{
		appG.Response(http.StatusOK,e.ERROR_CHECK_EXIST_ARTICLE_FAIL,nil)
		return
	}

	if !exist{
		appG.Response(http.StatusOK,e.ERROR_NOT_EXIT_ARTICLE,nil)
		return
	}

	err = articleService.Edit()
	if err != nil{
		appG.Response(http.StatusOK,e.ERROR_EDIT_ARTICLE_FAIL,nil)
		return
	}

	appG.Response(http.StatusOK,e.SUCCESS,make(map[string]interface{}))
}

//删除指定文章
func DeleteArticle(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Required(id, "id").Message("文章ID不能为空")
	valid.Min(id, 0, "id").Message("文章ID不能小于1")

	if valid.HasErrors(){
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK,e.INVALID_PARAMS,nil)
		return
	}

	articleService := article_service.Article{
		ID:id,
	}

	exist, err := articleService.ExistByID()
	if err != nil{
		appG.Response(http.StatusOK,e.ERROR_CHECK_EXIST_ARTICLE_FAIL,nil)
		return
	}

	if !exist{
		appG.Response(http.StatusOK,e.ERROR_NOT_EXIT_ARTICLE,nil)
		return
	}

	err = articleService.Delete()
	if err != nil {
		appG.Response(http.StatusOK,e.ERROR_DEL_ARTICLE_FAIL,nil)
		return
	}
	appG.Response(http.StatusOK,e.SUCCESS,make(map[string]interface{}))
}
