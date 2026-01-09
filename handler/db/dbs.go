package db

import (
	"Yearn-go/config"
	"Yearn-go/consts"
	"Yearn-go/handler/common"
	"Yearn-go/model"
	"Yearn-go/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

var sourceActionMap = map[string]common.HandlerFunc{
	"create": SuperCreateSource,
	"edit":   SuperEditSource,
}

// ManageDBCreateOrEdit 创建 修改 数据源
func ManageDBCreateOrEdit(g *gin.Context) {
	common.ActionDispatcher(g, sourceActionMap)
}

func SuperDeleteSource(g *gin.Context) {
	sourceId := g.Query("source_id")
	if sourceId == "" {
		utils.Fail(g, "未指定 source_id")
		return
	}

	if err := config.DB.Where("source_id = ?", sourceId).Delete(&model.CoreDataSource{SourceId: sourceId}).Error; err != nil {
		utils.Fail(g, fmt.Sprint(consts.ErrOperate, ": ", err))
		return
	}
	utils.Ok(g, nil)
}

func SuperFetchSource(g *gin.Context) {
	utils.Ok(g, nil)

}
