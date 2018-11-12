package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/oumeniOS/go-gin-blog/pkg/e"
	"github.com/oumeniOS/go-gin-blog/models"
	"net/http"
	"github.com/oumeniOS/go-gin-blog/pkg/util"
	"github.com/oumeniOS/go-gin-blog/pkg/setting"
	"github.com/oumeniOS/go-gin-blog/pkg/logging"
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
		valid.Min(tagId,1,"tag_id").Message("标签ID必须大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors(){
		data["list"] = models.ArticleList(util.GetPage(c),setting.PageSize,maps)
		data["total"] = models.GetArticleTotal(maps)
		code = e.SUCCESS
	} else {
		for _, err := range valid.Errors{
			logging.Info(err.Key,err.Error())
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetErrorMsg(code),
		"data":data,
	})
}

//获取指定文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()

	valid := validation.Validation{}
	valid.Required(id, "id").Message("Id 不能为空")
	valid.Min(id, 1, "id").Message("id 不能小于1")

	code := e.INVALID_PARAMS
	var data interface{}

	if ! valid.HasErrors() {
		if models.IsArticleExistById(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIT_ARTICLE
		}

	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Value)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorMsg(code),
		"data": data,
	})

}

//新建文章
func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	content := c.Query("content")
	desc := c.Query("desc")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.Query("state")).MustInt()

	valid := validation.Validation{}

	valid.Required(tagId,"tag_id").Message("标签ID不能问空")
	valid.Min(tagId,1,"tag_id").Message("标签ID最小为1")
	valid.Required(title,"title").Message("文章标题不能为空")
	valid.Required(content,"content").Message("文章内容不能为空")
	valid.Required(desc,"desc").Message("文章描述不能为空")
	valid.Required(createdBy,"created_by").Message("文章创建人不能为空")
	valid.Range(state,0,1,"state").Message("文章状态只能为0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistTagById(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["content"] = content
			data["desc"] = desc
			data["created_by"] = createdBy
			data["state"] = state
			models.NewArticle(data)
			code = e.SUCCESS
		}else {
			code = e.ERROR_NOT_EXIT_TAG
		}
	} else {
		for _, err := range valid.Errors{
			logging.Info(err.Key,err.Value)
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetErrorMsg(code),
		"data":make(map[string]interface{}),
	})
}

//更新指定文章
func EditArticle(c *gin.Context) {

	maps := make(map[string]interface{})
	id := com.StrTo(c.Query("id")).MustInt()

	title := c.Query("title")
	content := c.Query("content")
	desc := c.Query("desc")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}
	valid.Required(id,"id").Message("文章ID不能为空")
	valid.Min(id,1,"id").Message("文章ID不能小于1")
	valid.Required(modifiedBy,"modified_by").Message("文章修改人不能为空")
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
	if arg := c.Query("state");arg != ""{
		state := com.StrTo(arg).MustInt()
		valid.Range(state,0,1,"state").Message("文章状态只能为0或1")
		maps["state"] = state
	}
	code := e.INVALID_PARAMS
	if !valid.HasErrors(){
		if models.IsArticleExistById(id){
			models.EditArticle(id,maps)
			code= e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIT_ARTICLE
		}
	}else {
		for _, err := range valid.Errors{
			logging.Info(err.Key,err.Value)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetErrorMsg(code),
		"data":make(map[string]interface{}),
	})
}

//删除指定文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Required(id,"id").Message("文章ID不能为空")
	valid.Min(id,0,"id").Message("文章ID不能小于1")

	code :=e.INVALID_PARAMS
	if !valid.HasErrors(){
		if models.IsArticleExistById(id){
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIT_ARTICLE
		}
	}else {
		for _, err := range valid.Errors{
			logging.Info(err.Key,err.Value)
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetErrorMsg(code),
		"data":make(map[string]interface{}),
	})

}





































