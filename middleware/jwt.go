package middleware

import (
	"Yearn-go/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

type JwtConfig struct {
	GetKey     string
	AuthScheme string
	SigningKey string
}

var DefaultJwtConfig = JwtConfig{
	GetKey:     "Authorization",
	SigningKey: "algorithmHS256",
	AuthScheme: "Bearer ",
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader(DefaultJwtConfig.GetKey)
		if tokenStr == "" || !strings.HasPrefix(tokenStr, DefaultJwtConfig.AuthScheme) {
			utils.Fail(c, "缺少或格式错误的Token")
			c.Abort()
			return
		}

		// 去除前缀
		tokenStr = strings.TrimPrefix(tokenStr, DefaultJwtConfig.AuthScheme)

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// 验证签名方法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("意外的签名方法: %v", token.Header["alg"])
			}
			return []byte(DefaultJwtConfig.SigningKey), nil
		})
		if err != nil || !token.Valid {
			utils.Fail(c, "无效的Token")
			c.Abort()
			return
		}

		// 将 Claims 写入 Gin 上下文（方便后续使用）
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("claims", claims)
		}

		c.Next()
	}
}
