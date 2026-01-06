package group

import (
	"Yearn-go/config"
	"Yearn-go/consts"
	"Yearn-go/handler/common"
	"Yearn-go/model"
	"Yearn-go/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// validateSourceIds 负责真正的数据库校验
func validateSourceIds(db *gorm.DB, p model.PermissionList) error {
	sourceIds := utils.GetUniqueIds(p.DDLSource, p.DMLSource, p.QuerySource)
	if len(sourceIds) == 0 {
		return nil
	}

	var count int64
	if err := db.Model(&model.CoreDataSource{}).Where("source_id IN ?", sourceIds).Count(&count).
		Error; err != nil {
		return err
	}

	if count != int64(len(sourceIds)) {
		return errors.New("包含无效或不存在的数据源 ID")
	}
	return nil
}

func SuperGroupUpdate(g *gin.Context) {
	var p policy
	if err := g.ShouldBind(&p); err != nil {
		utils.Fail(g, fmt.Sprint(consts.ErrParamInvalid, ": ", err))
		return
	}

	// 去重校验数据源ID是否有效
	if err := validateSourceIds(config.DB, p.PermissionList); err != nil {
		utils.Fail(g, err.Error())
		return
	}

	per, _ := json.Marshal(p.PermissionList)

	rg := model.CoreRoleGroup{
		Name:        p.Name,
		Permissions: per,
	}

	if p.ID != 0 {
		if err := config.DB.Model(&model.CoreRoleGroup{}).
			Scopes(common.AccordingToIDEqual(p.ID)).
			Updates(&rg).Error; err != nil {
			utils.Fail(g, fmt.Sprint(consts.ErrOperate, ": ", err))
			return
		}
	} else {
		rg.GroupId = utils.GetUUID32()
		if err := config.DB.Save(&rg).Error; err != nil {
			utils.Fail(g, fmt.Sprint(consts.ErrOperate, ": ", err))
			return
		}
	}

	utils.Ok(g, nil)

}

func SuperClearUserRule(g *gin.Context) {
	utils.Ok(g, nil)

}

func SuperGroup(g *gin.Context) {
	common.HandlePaging[[]model.CoreRoleGroup](g, nil, common.PolicyQueryableFields)

}
