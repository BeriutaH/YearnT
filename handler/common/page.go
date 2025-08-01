package common

import (
	"Yearn-go/config"
	"gorm.io/gorm"
)

type PageInfo struct {
	Page     int `json:"page" binding:"required"`
	PageSize int `json:"page_size" binding:"required"`
}

type QueryRequest struct {
	PageInfo
	Filters map[string]string `json:"filters"` // 多字段模糊查询 {key:value}
}

type PageResult struct {
	Items    any   `json:"items"`     // 当前页的数据列表
	Total    int64 `json:"total"`     // 符合条件的总记录数
	Page     int   `json:"page"`      // 当前页码
	PageSize int   `json:"page_size"` // 每页条数
}

func SuccessPayload(data any, total int64, pageInfo PageInfo) PageResult {
	return PageResult{
		Items:    data,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}
}

type PageList[T any] struct {
	PageInfo
	Total int64 `json:"total"`
	Data  T     `json:"data"`
	order string
}

func (p *PageList[T]) ToMessage() PageResult {
	return SuccessPayload(p.Data, p.Total, p.PageInfo)
}

func (p *PageList[T]) Paging() *PageList[T] {
	if p.Page < 1 {
		p.Page = 1
	}
	switch {
	case p.PageSize > 100:
		p.PageSize = 100
	case p.PageSize <= 0:
		p.PageSize = 10
	}
	return p
}

func (p *PageList[T]) ToPageInfo(pageInfo PageInfo) *PageList[T] {
	p.PageInfo = pageInfo
	return p
}

func (p *PageList[T]) OrderBy(order string) *PageList[T] {
	p.order = order
	return p
}

func (p *PageList[T]) Query(scopes ...func(*gorm.DB) *gorm.DB) *PageList[T] {
	if p.order == "" {
		p.order = "id desc"
	}
	offset := (p.Page - 1) * p.PageSize
	config.DB.Model(p.Data).Scopes(scopes...).Count(&p.Total).Order(p.order).Offset(offset).Limit(p.PageSize).Find(&p.Data)
	return p
}
