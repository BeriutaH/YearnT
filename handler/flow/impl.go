package flow

import (
	"Yearn-go/config"
	"Yearn-go/model"
)

type Tpl struct {
	Desc    string   `json:"desc"`
	Auditor []string `json:"auditor"`
	Type    int      `json:"type"`
}

type tplTypes struct {
	Steps    []Tpl  `json:"steps"`
	Source   string `json:"source"`
	ID       int    `json:"id"`
	Relevant int    `json:"relevant"`
}

// checkFlowOrderCompletion 检查是否有正在执行的工单审批
func checkFlowOrderCompletion() []string {
	var workIds []string
	config.DB.Model(model.CoreSqlOrder{}).
		Select("work_id").
		Where("`status` =?", 2).
		Pluck("work_id", &workIds)
	return workIds
}
