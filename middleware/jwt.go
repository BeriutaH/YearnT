package middleware

import (
	"Yearn-go/consts"
	"Yearn-go/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

type JwtConfig struct {
	GetKey     string
	AuthScheme string
	SigningKey []byte
}

var DefaultJwtConfig = JwtConfig{
	GetKey:     "Authorization",
	SigningKey: []byte("frvWy3w9y9lSx+ZXlvGTC7wheS4b9P8CUekGL7vQR0Q"),
	AuthScheme: "Bearer ",
}

func JWTAuth() gin.HandlerFunc {
	return func(g *gin.Context) {
		tokenStr := g.GetHeader(DefaultJwtConfig.GetKey)
		if tokenStr == "" || !strings.HasPrefix(tokenStr, DefaultJwtConfig.AuthScheme) {
			utils.Fail(g, consts.ErrMissingOrInvalidToken)
			g.Abort()
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
			utils.Fail(g, consts.ErrInvalidToken)
			g.Abort() // 停止后续中间件或处理函数
			return
		}

		// 将claims存入context，供后续中间件使用
		if claimsI, ok := token.Claims.(jwt.MapClaims); ok {
			g.Set("claims", claimsI)
		}

		g.Next()
	}
}
