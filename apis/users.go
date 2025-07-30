package apis

import (
	"Yearn-go/utils"
	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	utils.Ok(c, "验证成功")
}
