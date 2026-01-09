package db

import (
	"Yearn-go/config"
	"Yearn-go/consts"
	"Yearn-go/factory"
	"Yearn-go/model"
	"Yearn-go/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CommonDBPost struct {
	Encrypt       bool                 `json:"encrypt"`
	DB            model.CoreDataSource `json:"db"`
	ExcludeDbList []string             `json:"exclude_db_list"`
	WordList      []string             `json:"word_list"`
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

	if err := config.DB.Model(&req).
		Where("id = ?", req.ID).
		Select("*").
		Omit(qmi...).
		Updates(&req).Error; err != nil {
		return false, fmt.Sprint(consts.ErrOperate, ": ", err)
	}

	// 根据group_id 更新,删除，CoreRoleSourcePrivilege的source_id
	if req.IsQuery == 0 {
		// 变为只写，将QuerySource列表里把这个删掉，清理掉所有类型为 "query" (只读) 的权限记录
		config.DB.Where("source_id = ? AND type = ?", req.SourceId, "query").
			Delete(&model.CoreRoleSourcePrivilege{})
	}
	if req.IsQuery == 1 {
		// 变为只读，将DDL 和 DML 列表里把这个删掉，清理掉所有类型为 "ddl" 或 "dml" (写权限) 的记录
		config.DB.Where("source_id = ? AND type IN ?", req.SourceId, []string{"ddl", "dml"}).
			Delete(&model.CoreRoleSourcePrivilege{})
	}

	return true, consts.MsgUpdateSuccess
}
