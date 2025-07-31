package routers

import (
	"Yearn-go/controllers"
	"Yearn-go/handler/user"
	"Yearn-go/middleware"
	"Yearn-go/restful"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.UserRegister)
	r.POST("/login", controllers.Login)

	// 下面的请求都开启JWT
	auth := r.Group("/api", middleware.JWTAuth())
	manager := auth.Group("/manage", middleware.SuperManageGroup())

	restful.Restful(manager, "user", user.SuperUserApi())

	return r
}
