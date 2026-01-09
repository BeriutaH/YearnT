package db

import (
	"Yearn-go/config"
	"Yearn-go/consts"
	"Yearn-go/factory"
	"Yearn-go/model"
	"Yearn-go/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommonDBPost struct {
	Encrypt       bool                 `json:"encrypt"`
	DB            model.CoreDataSource `json:"db"`
	ExcludeDbList []string             `json:"exclude_db_list"`
	WordList      []string             `json:"word_list"`
}

// updateSourceWithAtomic 数据源编辑及其连锁反应的事务逻辑
func updateSourceWithAtomic(req model.CoreDataSource, qmi []string) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {

		// 更新数据源主表 (强制更新零值并排除黑名单)
		if err := tx.Model(&req).
			Where("id = ?", req.ID).
			Select("*").
			Omit(qmi...).
			Updates(&req).Error; err != nil {
			return err
		}

		// 同步更新待处理工单的负责人
		if err := tx.Model(&model.CoreQueryOrder{}).
			Where("status = ? AND source_id = ?", 1, req.SourceId).
			Update("assigned", req.Principal).Error; err != nil {
			return err
		}

		// 动态权限清理，根据group_id 更新,删除，CoreRoleSourcePrivilege的source_id
		return cleanObsoletePrivileges(tx, req.SourceId, req.IsQuery)
	})
}

// cleanObsoletePrivileges 负责清理不合规的权限
func cleanObsoletePrivileges(tx *gorm.DB, sourceId string, isQuery int) error {
	query := tx.Where("source_id = ?", sourceId)

	switch isQuery {
	case 0: // 变为只写 -> 删掉读权限
		return query.Where("type = ?", "query").Delete(&model.CoreRoleSourcePrivilege{}).Error
	case 1: // 变为只读 -> 删掉写权限
		return query.Where("type IN ?", []string{"ddl", "dml"}).Delete(&model.CoreRoleSourcePrivilege{}).Error
	default:
		return nil
	}
}

// SuperCreateSource 添加数据库源信息
func SuperCreateSource(g *gin.Context) (bool, string) {
	var u CommonDBPost
	if err := g.ShouldBindJSON(&u); err != nil {
		return false, fmt.Sprint(consts.ErrParamInvalid, ": ", err)
	}

	source := u.DB
	// 判断 FlowID 是否存在
	if config.DB.Model(&model.CoreWorkflowTpl{}).Where("id = ?", source.FlowID).Limit(1).
		Find(&struct{ ID uint }{}).RowsAffected == 0 {
		return false, consts.ErrFlowInvalid
	}
	//  校验密码
	if source.Password = factory.Encrypt(config.Cfg.General.SecretKey, source.Password); source.Password == "" {
		return false, consts.ErrEncryptFailed
	}
	source.SourceId, source.ID = utils.GetUUID32(), 0

	if err := config.DB.Create(&source).Error; err != nil {
		return false, fmt.Sprint(consts.ErrOperate, ": ", err)
	}
	return true, consts.MsgCreateSuccess
}

// SuperEditSource 更新数据库源信息
func SuperEditSource(g *gin.Context) (bool, string) {
	// 除了密码，其余都携带最新信息，只有用户修改密码才会更新密码
	var req model.CoreDataSource
	if err := g.ShouldBindJSON(&req); err != nil {
		return false, fmt.Sprint(consts.ErrParamInvalid, ": ", err)
	}
	qmi := []string{"id", "source_id"}
	if req.Password != "" {
		req.Password = factory.Encrypt(config.Cfg.General.SecretKey, req.Password)
	} else {
		qmi = append(qmi, "password")
	}

	// 执行数据库更新操作
	if err := updateSourceWithAtomic(req, qmi); err != nil {
		return false, fmt.Sprint(consts.ErrOperate, ": ", err)
	}

	return true, consts.MsgUpdateSuccess
}
