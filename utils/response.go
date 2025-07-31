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
func Resp(g *gin.Context, code int, msg string, data interface{}) {
	if code == 0 {
		code = 200
	}
	if msg == "" {
		msg = "success"
	}
	if data == nil {
		data = gin.H{}
	}
	g.JSON(http.StatusOK, Response{Code: code, Message: msg, Data: data})
}

// Ok 成功响应，带 data
func Ok(g *gin.Context, data interface{}) {
	Resp(g, 0, "", data)
}

// Fail 失败响应，带 msg 和可选 code
func Fail(g *gin.Context, msg string, code ...int) {
	cod := 400 // 默认失败状态码
	if len(code) > 0 {
		cod = code[0]
	}
	Resp(g, cod, msg, nil)
}
