package factory

import (
	"encoding/json"
	"gorm.io/gorm"
)

type QueryRequest struct {
	Page     int               `json:"page" binding:"required"`
	PageSize int               `json:"page_size" binding:"required"`
	Filters  map[string]string `json:"filters"` // 多字段模糊查询 {key:value}
}

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
func (r *QueryRequest) ApplyFilters(db *gorm.DB, allowedFields map[string]bool) *gorm.DB {
	for field, value := range r.Filters {
		if allowedFields[field] && value != "" {
			db = AccordingToField(field, value)(db)
		}
	}
	return db
}

type PageResult[T any] struct {
	Items []T   `json:"items"` // 当前页的数据列表
	Total int64 `json:"total"` // 符合条件的总记录数
	Page  int   `json:"page"`  //当前页码
	Size  int   `json:"size"`  // 每页条数
}

// Paginate 分页器
func Paginate[T any](page, pageSize int, db *gorm.DB) (PageResult[T], error) {
	// 参数：
	//   - page: 当前页码，从 1 开始。
	//   - pageSize: 每页记录数，最大限制为 100。
	// 返回：一个 Gorm 的函数，用于 Offset 和 Limit 分页。
	var result PageResult[T]
	if page < 1 {
		page = 1
	}
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	// 查询总数
	if err := db.Count(&result.Total).Error; err != nil {
		return result, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	var items []T
	if err := db.Offset(offset).Limit(pageSize).Find(&items).Error; err != nil {
		return result, err
	}
	result.Items = items
	result.Page = page
	result.Size = pageSize

	return result, nil
}

func EmptyGroup() []byte {
	group, _ := json.Marshal([]string{})
	return group
}
