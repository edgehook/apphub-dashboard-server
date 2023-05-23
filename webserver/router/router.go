package router

import (
	v1 "github.com/edgehook/ithings/webserver/api/v1"
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
		apiv1.GET("/project", v1.GetProjects)
		apiv1.GET("/project/data/:id", v1.GetProjectDataById)
		apiv1.POST("/project/:id", v1.AddProject)
		apiv1.PUT("/project/data/:id", v1.SaveProjectData)
		apiv1.PUT("/project/:id", v1.UpdateProject)
	}

	return r
}
