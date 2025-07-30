package controllers

import (
	"Yearn-go/config"
	"Yearn-go/middleware"
	"Yearn-go/models"
	"Yearn-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var jwtKey = []byte("secret")

func Register(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效输入"})
		return
	}

	hashed, _ := utils.HashPassword(req.Password)
	user := models.User{Username: req.Username, Password: hashed}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "用户已存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

func Login(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效输入"})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无此用户"})
		return
	}

	if !utils.CheckPassword(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, _ := token.SignedString(middleware.DefaultJwtConfig.SigningKey)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
