package middleware

import (
	"Yearn-go/consts"
	"Yearn-go/factory"
	"Yearn-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuperManageGroup() gin.HandlerFunc {
	return func(g *gin.Context) {
		// 暂时先用户名是admin才行，后期变成角色为admin
		token := new(factory.Token).JwtParse(g)
		// 获取用户名
		if token.UserID == 1 {
			return
		}
		utils.Fail(g, consts.ErrUnauthorized, http.StatusForbidden)
		g.Abort()
	}
}
