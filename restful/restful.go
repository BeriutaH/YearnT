package restful

// todo 统一封装Gin的Restful风格的路径

import "github.com/gin-gonic/gin"

type RestfulAPI struct {
	Get    gin.HandlerFunc
	Post   gin.HandlerFunc
	Put    gin.HandlerFunc
	Delete gin.HandlerFunc
}

func Restful(r *gin.RouterGroup, path string, api RestfulAPI) {

	group := r.Group(path)
	if api.Get != nil {
		group.GET("", api.Get)
	}
	if api.Post != nil {
		group.POST("", api.Post)
	}
	if api.Put != nil {
		group.PUT("", api.Put)
	}
	if api.Delete != nil {
		group.DELETE("", api.Delete)
	}
}
