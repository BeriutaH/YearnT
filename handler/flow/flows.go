package flow

import (
	"Yearn-go/config"
	"Yearn-go/consts"
	"Yearn-go/model"
	"Yearn-go/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func TplGetAPis(g *gin.Context) {
	utils.Ok(g, nil)

}

// TplPostSourceTemplate 添加，修改审批流程
func TplPostSourceTemplate(g *gin.Context) {
	var u tplTypes
	if err := g.ShouldBindJSON(&u); err != nil {
		utils.Fail(g, consts.ErrParamInvalid+": "+err.Error())
		return
	}
	step, _ := json.Marshal(u.Steps)

	tpl := model.CoreWorkflowTpl{
		Source: u.Source,
		Steps:  step,
	}

	if u.ID != -1 {
		// 检查当前流程名下的工单
		if undo := checkFlowOrderCompletion(); len(undo) > 0 {
			utils.Fail(g, fmt.Sprintf("无法更新，存在未完成工单: %s", strings.Join(undo, ",")))
			return
		}

		if err := config.DB.Model(&model.CoreWorkflowTpl{}).
			Where("id = ?", u.ID).
			Select("Source", "Steps").
			Updates(&tpl).Error; err != nil {
			utils.Fail(g, consts.ErrOperate+": "+err.Error())
			return
		}
	} else {
		// 添加
		if err := config.DB.Create(&tpl).Error; err != nil {
			utils.Fail(g, consts.ErrOperate+": "+err.Error())
			return
		}
	}
	utils.Ok(g, nil)
}

func EditSourceTemplateInfo(g *gin.Context) {
	utils.Ok(g, nil)

}

func DeleteSourceTemplateInfo(g *gin.Context) {
	utils.Ok(g, nil)

}
