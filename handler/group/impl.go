package group

import "Yearn-go/model"

type policy struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	model.PermissionList
}
