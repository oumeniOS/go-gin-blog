package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/EDDYCJY/gin-blog/pkg/setting"
	"github.com/EDDYCJY/gin-blog/routers/api/v1"
	"github.com/EDDYCJY/gin-blog/routers/api"
	"github.com/EDDYCJY/gin-blog/middleware/jwt"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	 _"github.com/EDDYCJY/gin-blog/docs"
)

func InitRouters() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	//注册路由

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/auth", api.GetAuth)
	router.POST("/auth", api.AddUser)
	apiv1 := router.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签了；列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticleList)
		//获取指定的文章
		apiv1.POST("/article:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/article", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/article:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/article:id", v1.DeleteArticle)

	}
	return router
}