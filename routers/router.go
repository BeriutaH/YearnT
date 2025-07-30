package routers

import (
	"Yearn-go/controllers"
	"Yearn-go/handler"
	"Yearn-go/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.UserRegister)
	r.POST("/login", controllers.Login)

	// 下面的请求都开启JWT
	auth := r.Group("/api", middleware.JWTAuth())
	auth.GET("/me", handler.Me)

	return r
}
