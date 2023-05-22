package router

import (
	"github.com/edgehook/ithings/webserver/api/v1"
	"github.com/edgehook/ithings/webserver/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	//r.POST("/auth", api.GetAuth)
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/deviceModel/bypage", v1.GetModelsByPage)
		apiv1.POST("/deviceModel", v1.AddModels)
	}

	return r
}
