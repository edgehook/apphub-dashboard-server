package router

import (
	"net/http"
	"os"
	"path"

	v1 "github.com/edgehook/apphub-dashboard-server/webserver/api/v1"
	"github.com/edgehook/apphub-dashboard-server/webserver/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	if gin.Mode() == gin.DebugMode {
		r.Use(gin.Logger(), gin.Recovery())
	} else {
		r.Use(gin.Recovery())
	}
	r.Use(middlewares.Cors())
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
	var (
		pwd, _             = os.Getwd()
		vueAssetsRoutePath = pwd + string(os.PathSeparator) + "frontend" // 前端编译出来的 dist 所在路径
	)
	r.StaticFile("/", path.Join(vueAssetsRoutePath, "index.html"))             // 指定资源文件 url.  127.0.0.1/ 这种
	r.StaticFile("/favicon.ico", path.Join(vueAssetsRoutePath, "favicon.ico")) // 127.0.0.1/favicon.ico

	r.StaticFS("/assets", http.Dir(path.Join(vueAssetsRoutePath, "assets"))) // 以 assets 为前缀的 url
	r.StaticFS("/monacoeditorwork", http.Dir(path.Join(vueAssetsRoutePath, "monacoeditorwork")))
	r.StaticFS("/static", http.Dir(path.Join(vueAssetsRoutePath, "static")))
	return r
}
