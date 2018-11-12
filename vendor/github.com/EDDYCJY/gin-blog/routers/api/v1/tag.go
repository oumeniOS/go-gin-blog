package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/EDDYCJY/gin-blog/pkg/e"
	"github.com/EDDYCJY/gin-blog/models"
	"github.com/EDDYCJY/gin-blog/pkg/util"
	"github.com/EDDYCJY/gin-blog/pkg/setting"
	"net/http"
	"github.com/astaxie/beego/validation"
)

//获取多个文章标签
func GetTags(context *gin.Context) {
	name := context.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS
	data["list"] = models.GetTags(util.GetPage(context), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorMsg(code),
		"data": data,
		"maps": maps,
	})
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "TAG标题"
// @Param state query int false "TAG状态"
// @Param created_by query int false "创建人"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddTag(context *gin.Context) {
	name := context.Query("name")
	state := com.StrTo(context.DefaultQuery("state", "0")).MustInt()
	createdBy := context.Query("createdBy")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "createdBy").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "createdBy").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			success := models.AddTag(name, state, createdBy)
			if !success {
				println("models.AddTag err")
			}
		} else {
			code = e.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			println(err.Key, err.Message)
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorMsg(code),
		"data": make(map[string]string),
	})
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
	id := com.StrTo(context.Query("id")).MustInt()
	name := context.Query("name")
	state := com.StrTo(context.Query("state")).MustInt()
	modifiedBy := context.Query("modified_by")

	valid := validation.Validation{}

	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "motified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagById(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIT_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			println(err.Key, err.Message)
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorMsg(code),
		"data": make(map[string]string),
	})
}

//@Summary 删除标签
//@Produce json
//@Param id query int true "ID"
//@Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
//@Router /api/v1/tags/{id} [delete]
func DeleteTag(context *gin.Context) {
	id := com.StrTo(context.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id 不能为空")
	valid.Min(id, 1, "id").Message("id不能小于1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.DeleteTagById(id) {
			code = e.SUCCESS
		} else {
			code = e.ERROR
		}
	} else {
		for _, err := range valid.Errors {
			println(err.Key, err.Message)
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorMsg(code),
		"data": make(map[string]string),
	})
}
