package controllers

import (
	"Yearn-go/config"
	"Yearn-go/consts"
	"Yearn-go/factory"
	"Yearn-go/middleware"
	"Yearn-go/models"
	"Yearn-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	hashed, _ := factory.HashPassword(u.Password) // 加密密码
	user := models.CoreAccount{Username: u.Username, Password: hashed}
	if err := config.DB.Create(&user).Error; err != nil {
		utils.Fail(c, consts.ErrUserExists)
		return
	}

	utils.Ok(c, consts.MsgRegisterSuccess)
}

func Login(c *gin.Context) {
	u := new(models.CoreAccount)
	if err := c.ShouldBindJSON(&u); err != nil {
		utils.Fail(c, consts.ErrInvalidInput)
		return
	}

	var user models.CoreAccount
	if err := config.DB.Where("username = ?", u.Username).First(&user).Error; err != nil {
		utils.Fail(c, consts.ErrUserNotFound)
		return
	}

	if !factory.CheckPassword(user.Password, u.Password) {
		utils.Fail(c, consts.ErrInvalidPassword)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     "",
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // 3 天后过期
	})

	tokenString, _ := token.SignedString(middleware.DefaultJwtConfig.SigningKey)
	utils.Ok(c, gin.H{"token": tokenString})
}
