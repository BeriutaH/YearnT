package common

import (
	"Yearn-go/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"reflect"
)

// HandlerFunc 定义了分发函数签名：返回成功标识和错误消息
type HandlerFunc func(*gin.Context) (bool, string)

// ActionDispatcher 通用的请求分发器
func ActionDispatcher(g *gin.Context, mapper map[string]HandlerFunc) {
	// mapper: 存储 action 字符串与对应处理函数的映射
	action := g.Query("action")

	// 查找是否存在对应的操作
	handler, ok := mapper[action]
	if !ok {
		utils.Fail(g, "不支持的操作类型: "+action)
		return
	}
	// 执行业务逻辑
	success, msg := handler(g)
	// 统一处理结果（
	utils.HandleResult(g, success, msg)
}

// StructToMap 结构体转map
func StructToMap(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	b, err := json.Marshal(s)
	if err != nil {
		return m
	}
	_ = json.Unmarshal(b, &m)
	return m
}

// RemoveZeroValues 去掉零值 跟 id
func RemoveZeroValues(m map[string]interface{}) map[string]interface{} {
	filtered := make(map[string]interface{})
	for k, v := range m {
		if k == "id" {
			continue
		}
		if !isZero(reflect.ValueOf(v)) {
			filtered[k] = v
		}
	}
	return filtered
}

func isZero(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
