package model

type PermissionList struct {
	DDLSource   []string `json:"ddl_source"`
	DMLSource   []string `json:"dml_source"`
	QuerySource []string `json:"query_source"`
}
