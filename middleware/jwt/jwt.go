package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/oumeniOS/go-gin-blog/pkg/e"
	"github.com/oumeniOS/go-gin-blog/pkg/util"
	"time"
	"net/http"
	"github.com/oumeniOS/go-gin-blog/pkg/logging"
)

func JWT()gin.HandlerFunc  {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == ""{
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil{
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				logging.Info(err)
			} else if time.Now().Unix() > claims.ExpiresAt{
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				logging.Info(err)
			}
		}
		if code != e.SUCCESS{
			c.JSON(http.StatusUnauthorized,gin.H{
				"code":code,
				"msg":e.GetErrorMsg(code),
				"data":data,
			})
			c.Abort()
			return
		}
		c.Next()

	}
}




















































