package factory

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	RealName string `json:"role"`
}

func (h *Token) JwtParse(g *gin.Context) *Token {
	claimsRaw, _ := g.Get("claims")
	claims, _ := claimsRaw.(jwt.MapClaims)
	h.Username = claims["username"].(string)
	h.RealName = claims["role"].(string)
	uid, _ := claims["user_id"].(float64)
	h.UserID = int(uid)
	fmt.Println("111", claims)
	return h
}
