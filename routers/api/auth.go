package api

import (
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"github.com/oumeniOS/go-gin-blog/pkg/e"
	"github.com/oumeniOS/go-gin-blog/models"
	"github.com/oumeniOS/go-gin-blog/pkg/util"
	"net/http"
	"github.com/oumeniOS/go-gin-blog/pkg/logging"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}


func GetAuth(c *gin.Context)  {
	username := c.Query("username")
	password := c.Query("password")
	valid := validation.Validation{}

	a := auth{Username:username,Password:password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ok {
		isExist := models.CheckAuth(username,password)
		if isExist {
			token, err := util.GenerateToken(username,password)
			if err != nil{
				code = e.ERROR_AUTH_TOKEN
			}else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors{
			logging.Info(err.Key,err.Value)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetErrorMsg(code),
		"data":data,
	})
}

func AddUser(c *gin.Context)  {
	username := c.Query("username")
	password := c.Query("password")

	a := auth{
		Username:username,
		Password:password,
	}
	valid := validation.Validation{}

	ok , _ := valid.Valid(&a)
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username,password)
		if isExist {
			code = e.ERROR_AUTH_EXISTED
		} else {
			models.AddAuth(username,password)
			code = e.SUCCESS
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key,err.Value)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetErrorMsg(code),
		"data":make(map[string]interface{}),
	})
}




































