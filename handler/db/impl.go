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
func SuperEditSource(g *gin.Context) (bool, string) {
	// 除了密码，其余都携带最新信息，只有用户修改密码才会更新密码
	var req model.CoreDataSource
	if err := g.ShouldBindJSON(&req); err != nil {
		return false, fmt.Sprint(consts.ErrParamInvalid, ": ", err)
	}
	var qmi []string
	if req.Password == "" {

		qmi = []string{"id", "source_id", "password"}
	} else {

		qmi = []string{"id", "source_id"}
		req.Password = factory.Encrypt(config.Cfg.General.SecretKey, req.Password)
	}

	if err := config.DB.Model(&req).
		Where("id = ?", req.ID).
		Select("*").
		Omit(qmi...).
		Updates(&req).Error; err != nil {
		return false, fmt.Sprint(consts.ErrOperate, ": ", err)
	}
	return true, consts.MsgCreateSuccess
}
