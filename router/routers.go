package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"studyGIN/controllers"
	"time"
)

func Router() *gin.Engine {

	r := gin.Default()

	// 配置更详细的 CORS 中间件
	config := cors.Config{
		AllowOrigins:     []string{"https://jy.eweda.cn"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "x-app"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(config))

	test := r.Group("/ceshi")
	{
		test.GET("/test", controllers.UserController{}.GetTest)
		test.GET("/get/:id", controllers.UserController{}.GetUserInfo)
		test.POST("/post", controllers.UserController{}.PostUserInfo)
		test.PUT("/put", func(context *gin.Context) {
			context.String(http.StatusOK, "put")
		})
		test.DELETE("/delete", func(context *gin.Context) {
			context.String(http.StatusOK, "delete")
		})
	}
	link := r.Group("/link")
	{
		link.GET("/ces", controllers.UserController{}.GetLinkTime)
		link.GET("/content", controllers.UserController{}.GetLinkContent)
		link.GET("/ws", controllers.UserController{}.GetLinkWS)
	}
	export := r.Group("/export")
	{
		export.GET("", controllers.UserController{}.ExportTable)
		export.GET("/contrast", controllers.UserController{}.ExportAdd)
	}

	return r
}
