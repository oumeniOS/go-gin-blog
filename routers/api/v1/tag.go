package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/oumeniOS/go-gin-blog/pkg/e"
	"github.com/oumeniOS/go-gin-blog/models"
			"net/http"
	"github.com/astaxie/beego/validation"
	"github.com/oumeniOS/go-gin-blog/pkg/app"
	"github.com/oumeniOS/go-gin-blog/service/tag_service"
	"github.com/oumeniOS/go-gin-blog/pkg/export"
	"github.com/oumeniOS/go-gin-blog/pkg/logging"
	"github.com/oumeniOS/go-gin-blog/pkg/setting"
)

//获取多个文章标签
func GetTags(context *gin.Context) {
	var (
		data       map[string]interface{} = make(map[string]interface{})
		list []*models.Tag
		tagService *tag_service.Tag
		appG       app.Gin
		name string
		state int = -1
		err error
		totalCount int
	)

	appG.C = context
	name = context.Query("name")
	if name != "" {
		tagService.Name = name
	}

	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		tagService.State = state
	}

	list, err = tagService.GetAll()
	if err != nil{
		appG.Response(http.StatusOK,e.ERROR_GET_TAG_FAIL,nil)
		return
	}
	data["list"] = list


	totalCount, err = tagService.TotalCount()
	if err != nil{
		appG.Response(http.StatusOK,e.ERROR_GET_TAG_FAIL,nil)
		return
	}
	data["total"] = totalCount

	appG.Response(http.StatusOK,e.SUCCESS,data)
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "TAG标题"
// @Param state query int false "TAG状态"
// @Param created_by query int false "创建人"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddTag(context *gin.Context) {

	var(
		name string
		state int = -1
		createdBy string
		appG app.Gin = app.Gin{context}
		serviceTag *tag_service.Tag
		err error
	)

	name = context.PostForm("name")
	state = com.StrTo(context.DefaultQuery("state", "0")).MustInt()
	createdBy = context.Query("createdBy")
	serviceTag.Name = name
	serviceTag.State = state
	serviceTag.CreatedBy = createdBy

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "createdBy").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "createdBy").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	if valid.HasErrors() {
		appG.Response(http.StatusOK,e.INVALID_PARAMS,nil)
		return
	}

	//排重
	exist,_ := serviceTag.Exist()
	if exist == true{
		appG.Response(http.StatusOK,e.ERROR_EXIST_TAG,nil)
		return
	}

	err = serviceTag.Add()
	if err != nil{
		appG.Response(http.StatusOK,e.ERROR_ADD_TAG_FAIL,nil)
		return
	}

	appG.Response(http.StatusOK,e.SUCCESS,make(map[string]interface{}))
}

// @Summary 修改文章标签
// @Produce  json
// @Param id query int true "ID"
// @Param name query string true "ID"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [put]
func EditTag(context *gin.Context) {

	var(
		id int
		name string
		state int = -1
		modifiedBy string
		appG app.Gin = app.Gin{context}
		serviceTag = tag_service.Tag{}
	)

	id = com.StrTo(context.Query("id")).MustInt()
	name = context.Query("name")
	state = com.StrTo(context.Query("state")).MustInt()
	modifiedBy = context.Query("modified_by")
	serviceTag.ID = id
	serviceTag.Name = name
	serviceTag.State = state
	serviceTag.ModifiedBy = modifiedBy


	valid := validation.Validation{}

	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "motified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长100字符")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK,e.INVALID_PARAMS,nil)
		return
	}

	//排重
	exist,_ := serviceTag.Exist()
	if exist == true{
		appG.Response(http.StatusOK,e.ERROR_EXIST_TAG,nil)
		return
	}

	err := serviceTag.Edit()
	if err != nil{
		appG.Response(http.StatusOK,e.ERROR_EDIT_TAG_FAIL,nil)
		return
	}

	appG.Response(http.StatusOK,e.SUCCESS,make(map[string]interface{}))
}

//@Summary 删除标签
//@Produce json
//@Param id query int true "ID"
//@Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
//@Router /api/v1/tags/{id} [delete]
func DeleteTag(context *gin.Context) {

	var(
		id int
		appG app.Gin = app.Gin{context}
		serviceTag tag_service.Tag
	)

	id = com.StrTo(context.Query("id")).MustInt()
	serviceTag.ID = id

	valid := validation.Validation{}
	valid.Required(id, "id").Message("id 不能为空")
	valid.Min(id, 1, "id").Message("id不能小于1")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK,e.INVALID_PARAMS,nil)
		return
	}

	err := serviceTag.Delete()
	if err != nil{
		appG.Response(http.StatusOK,e.ERROR_DEL_TAG_FAIL,nil)
		return
	}

	appG.Response(http.StatusOK,e.SUCCESS,make(map[string]string))
}


func ExportTag(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	tagService := tag_service.Tag{
		Name:  name,
		State: state,
		PageSize:setting.AppSetting.PageSize,
	}

	filename, err := tagService.Export()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_EXPORT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"export_url":      export.GetExcelFullUrl(filename),
		"export_save_url": export.GetExcelPath() + filename,
	})
}

func ImportTag(c *gin.Context) {
	appG := app.Gin{C: c}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	tagService := tag_service.Tag{}
	err = tagService.Import(file)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusOK, e.ERROR_IMPORT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}