package app

import (
	"github.com/gin-gonic/gin"
	"github.com/oumeniOS/go-gin-blog/pkg/e"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{})  {
	g.C.JSON(httpCode,gin.H{
		"code":errCode,
		"msg":e.GetErrorMsg(errCode),
		"data":data,
	})
	return
}














