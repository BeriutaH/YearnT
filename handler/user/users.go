package user

import (
	"Yearn-go/utils"
	"github.com/gin-gonic/gin"
)

func InterfaceTestO(g *gin.Context) {
	utils.Ok(g, "Get验证成功")
}

func InterfaceTestT(g *gin.Context) {
	utils.Ok(g, "PUT验证成功")
}

func ManageUserCreateOrEdit(g *gin.Context) {
	var success bool
	var msg string
	// 获取参数，判断操作类型
	switch g.Query("op") {
	case "add":
		success, msg = CreateUser(g)
	}

	utils.HandleResult(g, success, msg)
}

func InterfaceTestF(g *gin.Context) {
	idUser := g.Param("ya")
	println(idUser)
	utils.Ok(g, "DELETE验证成功")
}
