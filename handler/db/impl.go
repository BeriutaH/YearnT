package db

import (
	"Yearn-go/config"
	"Yearn-go/consts"
	"Yearn-go/factory"
	"Yearn-go/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		return false, consts.ErrParamInvalid + ": " + err.Error()
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
	source.SourceId, source.ID = uuid.New().String(), 0

	if err := config.DB.Create(&source).Error; err != nil {
		return false, consts.ErrOperate + ": " + err.Error()
	}
	return true, consts.MsgCreateSuccess
}
