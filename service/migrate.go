package service

import (
	"Yearn-go/model"
)

// AutoMigrateAll 数据库迁移模型配置
func AutoMigrateAll() []interface{} {
	return []interface{}{

		// 第一阶段：基础表（主表），必须先存在
		&model.CoreAccount{},
		&model.CoreRoleGroup{},
		&model.CoreDataSource{},

		// 第二阶段：关联表（中间表/外键表），依赖于上面的主表
		&model.CoreGrained{},
		&model.CoreGlobalConfiguration{},
		&model.CoreSqlRecord{},
		&model.CoreSqlOrder{},
		&model.CoreRollback{},
		&model.CoreQueryOrder{},
		&model.CoreQueryRecord{},
		&model.CoreAutoTask{},
		&model.CoreWorkflowTpl{},
		&model.CoreWorkflowDetail{},
		&model.CoreOrderComment{},
		&model.CoreRules{},
		&model.CoreTotalTickets{},
		&model.CoreRoleSourcePrivilege{},
	}
}
