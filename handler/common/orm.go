package common

import (
	"encoding/json"
	"gorm.io/gorm"
)

// AccordingToField 通用条件单个查询
func AccordingToField(fieldName, keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if keyword == "" {
			return db
		}
		return db.Where(fieldName+" LIKE ?", "%"+keyword+"%")
	}
}

// ApplyFilters 模糊查询
func ApplyFilters(allowedFields map[string]bool, filters map[string]string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for field, value := range filters {
			if allowedFields[field] && value != "" {
				db = AccordingToField(field, value)(db)
			}
		}
		return db
	}
}

// QmiFilters 排除查询
func QmiFilters(omitFields []string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(omitFields) > 0 {
			return db.Omit(omitFields...)
		}
		return db
	}
}

func EmptyGroup() []byte {
	group, _ := json.Marshal([]string{})
	return group
}
