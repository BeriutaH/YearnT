package routers

import (
	"Yearn-go/apis"
	"Yearn-go/controllers"
	"Yearn-go/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// 下面的请求都开启JWT
	auth := r.Group("/api", middleware.JWTAuth())
	auth.GET("/me", apis.Me)

	return r
}
