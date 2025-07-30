package service

import (
	"Yearn-go/models"
)

// AutoMigrateAll 数据库迁移模型配置
func AutoMigrateAll() []interface{} {
	return []interface{}{
		&models.CoreAccount{},
		&models.CoreGlobalConfiguration{},
		&models.CoreSqlRecord{},
		&models.CoreSqlOrder{},
		&models.CoreRollback{},
		&models.CoreDataSource{},
		&models.CoreGrained{},
		&models.CoreRoleGroup{},
		&models.CoreQueryOrder{},
		&models.CoreQueryRecord{},
		&models.CoreAutoTask{},
		&models.CoreWorkflowTpl{},
		&models.CoreWorkflowDetail{},
		&models.CoreOrderComment{},
		&models.CoreRules{},
		&models.CoreTotalTickets{},
	}
}
