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
			g.Next() // 验证通过，继续执行后续中间件和处理函数
			return
		}
		utils.Fail(g, consts.ErrUnauthorized, http.StatusForbidden)
		g.Abort() // 停止后续中间件或处理函数
		return    // 确保函数立即返回，不继续执行
	}
}
