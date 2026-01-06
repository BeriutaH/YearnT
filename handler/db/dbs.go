package db

import (
	"Yearn-go/handler/common"
	"Yearn-go/utils"
	"github.com/gin-gonic/gin"
)

var sourceActionMap = map[string]common.HandlerFunc{
	"create": SuperCreateSource,
	//"edit":   EditUser,
	//"reset":  ResetPwdUser,
	//"policy": EditPayloadUser,
}

// ManageDBCreateOrEdit 创建 修改 数据源
func ManageDBCreateOrEdit(g *gin.Context) {
	common.ActionDispatcher(g, sourceActionMap)
}

func SuperDeleteSource(g *gin.Context) {
	utils.Ok(g, nil)

}

func SuperFetchSource(g *gin.Context) {
	utils.Ok(g, nil)

}
