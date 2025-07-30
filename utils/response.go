package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 通用响应结构体
type Response struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 数据内容
}

// Resp 通用响应函数，可传 code、msg、data，自动处理默认值
func Resp(c *gin.Context, code int, msg string, data interface{}) {
	if code == 0 {
		code = 200
	}
	if msg == "" {
		msg = "success"
	}
	if data == nil {
		data = gin.H{}
	}
	c.JSON(http.StatusOK, Response{Code: code, Message: msg, Data: data})
}

// Ok 成功响应，带 data
func Ok(c *gin.Context, data interface{}) {
	Resp(c, 200, "success", data)
}

// Fail 失败响应，带 msg 和可选 code
func Fail(c *gin.Context, msg string, code ...int) {
	cod := 400 // 默认失败状态码
	if len(code) > 0 {
		cod = code[0]
	}
	Resp(c, cod, msg, nil)
}
