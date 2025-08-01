package service

import (
	"Yearn-go/model"
)

// AutoMigrateAll 数据库迁移模型配置
func AutoMigrateAll() []interface{} {
	return []interface{}{
		&model.CoreAccount{},
		&model.CoreGlobalConfiguration{},
		&model.CoreSqlRecord{},
		&model.CoreSqlOrder{},
		&model.CoreRollback{},
		&model.CoreDataSource{},
		&model.CoreGrained{},
		&model.CoreRoleGroup{},
		&model.CoreQueryOrder{},
		&model.CoreQueryRecord{},
		&model.CoreAutoTask{},
		&model.CoreWorkflowTpl{},
		&model.CoreWorkflowDetail{},
		&model.CoreOrderComment{},
		&model.CoreRules{},
		&model.CoreTotalTickets{},
	}
}
