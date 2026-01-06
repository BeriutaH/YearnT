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

// GetUniqueIds 负责从嵌套结构中提取并去重 UUID
func GetUniqueIds(lists ...[]string) []string {
	idMap := make(map[string]struct{})

	for _, list := range lists {
		for _, id := range list {
			if id != "" {
				idMap[id] = struct{}{}
			}
		}
	}

	return utils.MapKeys(idMap)
}

// ValidateSourceIds 负责真正的数据库校验
func ValidateSourceIds(db *gorm.DB, p model.PermissionList) error {
	sourceIds := GetUniqueIds(p.DDLSource, p.DMLSource, p.QuerySource)
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

	// 去重校验group_id是否有效
	if err := ValidateSourceIds(config.DB, p.PermissionList); err != nil {
		utils.Fail(g, err.Error())
		return
	}

	per, _ := json.Marshal(p.PermissionList)

	groupPolicy := model.CoreRoleGroup{
		Name:        p.Name,
		Permissions: per,
	}

	if p.ID != 0 {
		if err := config.DB.Model(&model.CoreRoleGroup{}).
			Scopes(common.AccordingToIDEqual(p.ID)).
			Updates(&groupPolicy).Error; err != nil {
			utils.Fail(g, fmt.Sprint(consts.ErrOperate, ": ", err))
			return
		}
	} else {
		groupPolicy.GroupId = utils.GetUUID32()
		if err := config.DB.Save(&groupPolicy).Error; err != nil {
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
	utils.Ok(g, nil)

}
