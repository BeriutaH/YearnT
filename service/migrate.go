package service

import (
	"Yearn-go/models"
)

// AutoMigrateAll 数据库迁移模型配置
func AutoMigrateAll() []interface{} {
	return []interface{}{
		&models.User{},
		// 未来可以添加更多模型
		// &models.Role{},
	}
}
