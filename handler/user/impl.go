package user

import (
	"Yearn-go/config"
	"Yearn-go/consts"
	"Yearn-go/factory"
	"Yearn-go/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type CommonUser struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Department string `json:"department" binding:"required"`
	RealName   string `json:"real_name" binding:"required"`
	Email      string `json:"email" binding:"required,email"` // email 格式校验
}

func CreateUser(g *gin.Context) (bool, string) {
	var u CommonUser
	if err := g.ShouldBindJSON(&u); err != nil {
		return false, consts.ErrParamInvalid + ": " + err.Error()
	}
	// 判断是否重名
	var unique models.CoreAccount
	if err := config.DB.Where("username = ?", u.Username).Select("username").First(&unique).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, consts.ErrUserExists
	}
	// 加密密码
	u.Password = factory.DjangoEncrypt(u.Password, string(factory.GetRandom()))
	var user models.CoreAccount
	if err := copier.Copy(&user, &u); err != nil {
		return false, consts.ErrOperate + ": " + err.Error()
	}
	user.IsRecorder = 2

	// 添加数据库
	config.DB.Create(&user)
	config.DB.Create(&models.CoreGrained{Username: u.Username, Group: factory.EmptyGroup()})

	return true, "用户" + consts.MsgCreateSuccess
}
