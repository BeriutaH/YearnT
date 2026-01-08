package group

import (
	"Yearn-go/config"
	"Yearn-go/consts"
	"Yearn-go/handler/common"
	"Yearn-go/model"
	"Yearn-go/utils"
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

func updateGroupPrivileges(tx *gorm.DB, req policy, gid string) error {
	var list []model.CoreRoleSourcePrivilege

	// 1. 构造列表
	addToList := func(sources []string, pType string) {
		for _, id := range sources {
			list = append(list, model.CoreRoleSourcePrivilege{
				GroupId:  gid,
				SourceId: id,
				Type:     pType,
			})
		}
	}
	addToList(req.DDLSource, "ddl")
	addToList(req.DMLSource, "dml")
	addToList(req.QuerySource, "query")

	// 删除旧权限记录
	if err := tx.Where("group_id = ?", gid).Delete(&model.CoreRoleSourcePrivilege{}).Error; err != nil {
		return err
	}

	// 插入新权限记录
	if len(list) > 0 {
		if err := tx.Create(&list).Error; err != nil {
			return err
		}
	}
	return nil
}

func SuperGroup(g *gin.Context) {
	common.HandlePaging[[]model.CoreRoleGroup](g, nil, common.PolicyQueryableFields)

}

// SuperGroupUpdate 创建，修改 权限组对应的数据库源
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

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		var gid string
		rg := model.CoreRoleGroup{Name: p.Name}

		if p.ID != 0 {
			// 更新模式：先获取现有的 GroupId
			var existing model.CoreRoleGroup
			if err := tx.First(&existing, p.ID).Error; err != nil {
				return err
			}
			gid = existing.GroupId
			// 更新组名
			if err := tx.Model(&existing).Updates(&rg).Error; err != nil {
				return err
			}
		} else {
			// 新增模式：生成新的 GroupId
			gid = utils.GetUUID32()
			rg.GroupId = gid
			if err := tx.Create(&rg).Error; err != nil {
				return err
			}
		}

		// 更新最新权限
		return updateGroupPrivileges(tx, p, gid)
	})

	if err != nil {
		utils.Fail(g, fmt.Sprint(consts.ErrOperate, ": ", err))
		return
	}

	utils.Ok(g, nil)

}

// SuperClearUserRule 删除权限以及对应的权限组
func SuperClearUserRule(g *gin.Context) {

	gId := g.Query("group_id")
	if gId == "" {
		utils.Fail(g, "删除失败：未指定 group_id")
		return
	}

	if err := config.DB.Unscoped().Where("group_id = ?", gId).Delete(&model.CoreRoleGroup{}).Error; err != nil {
		utils.Fail(g, fmt.Sprint(consts.ErrOperate, ": ", err))
		return
	}
	utils.Ok(g, nil)

}
