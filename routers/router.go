package routers

import (
	"Yearn-go/controllers"
	"Yearn-go/handler/db"
	"Yearn-go/handler/flow"
	"Yearn-go/handler/group"
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

	// 用户相关
	restful.Restful(manager, "user", user.SuperUserApi())
	// 权限组相关
	restful.Restful(manager, "policy", group.GroupsApis())
	// 数据库链接相关
	restful.Restful(manager, "db", db.ManageDbApis())
	// 审批相关
	restful.Restful(manager, "tpl", flow.TplRestApis())

	return r
}
