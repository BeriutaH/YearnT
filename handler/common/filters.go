package common

// UserQueryableFields 定义用户账号允许作为查询条件的字段（支持模糊查询）
var UserQueryableFields = map[string]bool{
	"username":  true,
	"email":     true,
	"real_name": true,
}

// UserSensitiveFields 定义用户账号对外查询时需要排除的敏感字段
var UserSensitiveFields = []string{"password"}
