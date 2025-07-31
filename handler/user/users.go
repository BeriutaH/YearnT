package user

import (
	"Yearn-go/utils"
	"github.com/gin-gonic/gin"
)

func InterfaceTestO(c *gin.Context) {
	utils.Ok(c, "Get验证成功")
}

func InterfaceTestT(c *gin.Context) {
	utils.Ok(c, "PUT验证成功")
}

func InterfaceTes3tS(c *gin.Context) {
	utils.Ok(c, "POST验证成功")
}

func InterfaceTestF(c *gin.Context) {
	idUser := c.Param("ya")
	println(idUser)
	utils.Ok(c, "DELETE验证成功")
}
