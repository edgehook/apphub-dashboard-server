package router

import (
	"net/http"
	"path"

	v1 "github.com/edgehook/apphub-dashboard-server/webserver/api/v1"
	"github.com/edgehook/apphub-dashboard-server/webserver/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	// r.LoadHTMLGlob("frontend/index.html")
	// r.Static("/static", "./static")
	// r.Static("/assets", "./assets")
	// r.Static("/monacoeditorwork", "./monacoeditorwork")
	// r.StaticFile("./favicon.ico", "./favicon.ico")

	var (
		vueAssetsRoutePath = "./frontend" // 前端编译出来的 dist 所在路径
	)
	r.StaticFile("/", path.Join(vueAssetsRoutePath, "index.html"))             // 指定资源文件 url.  127.0.0.1/ 这种
	r.StaticFile("/favicon.ico", path.Join(vueAssetsRoutePath, "favicon.ico")) // 127.0.0.1/favicon.ico

	r.StaticFS("/assets", http.Dir(path.Join(vueAssetsRoutePath, "assets"))) // 以 assets 为前缀的 url
	r.StaticFS("/monacoeditorwork", http.Dir(path.Join(vueAssetsRoutePath, "monacoeditorwork")))
	r.StaticFS("/static", http.Dir(path.Join(vueAssetsRoutePath, "static")))
	//r.POST("/auth", api.GetAuth)
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/v1")
	//apiv1.Use(jwt.JWT())
	screenApi := apiv1.Group("screen")
	screenApi.GET("", v1.GetScreens)
	screenApi.GET("/image", v1.GetScreenImagesByScreenIdAndType)
	screenApi.GET("/:id", v1.GetScreenById)
	screenApi.GET("/isAside", v1.GetScreensByIsAside)
	screenApi.POST("", v1.AddScreen)
	screenApi.POST("/image", v1.AddScreenImage)
	screenApi.POST("/:id/copy", v1.CopyScreen)
	screenApi.PUT("/:id", v1.UpdateScreen)
	screenApi.DELETE("/:id", v1.DeleteScreen)
	return r
}
