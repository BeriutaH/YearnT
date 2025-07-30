package controllers

import (
	"Yearn-go/config"
	"Yearn-go/consts"
	"Yearn-go/middleware"
	"Yearn-go/models"
	"Yearn-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

func UserRegister(c *gin.Context) {

	if !consts.GloRegister {
		utils.Fail(c, consts.ErrRegisterDisabled)
		return
	}
	u := new(models.CoreAccount)
	if err := c.ShouldBindJSON(&u); err != nil {
		utils.Fail(c, consts.ErrInvalidInput)
		return
	}

	hashed, _ := utils.HashPassword(u.Password) // 加密密码
	user := models.CoreAccount{Username: u.Username, Password: hashed}
	if err := config.DB.Create(&user).Error; err != nil {
		utils.Fail(c, consts.ErrUserExists)
		return
	}

	utils.Ok(c, "注册成功")
}

func Login(c *gin.Context) {
	u := new(models.CoreAccount)
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效输入"})
		return
	}

	var user models.CoreAccount
	if err := config.DB.Where("username = ?", u.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无此用户"})
		return
	}

	if !utils.CheckPassword(user.Password, u.Password) {
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
