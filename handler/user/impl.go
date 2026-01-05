package user

import (
	"Yearn-go/config"
	"Yearn-go/consts"
	"Yearn-go/factory"
	"Yearn-go/handler/common"
	"Yearn-go/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var validate = validator.New()

// BaseUser 用户共用
type BaseUser struct {
	Username   string `json:"username"`
	Department string `json:"department"`
	RealName   string `json:"real_name"`
	Email      string `json:"email"`
}
type PwdType struct {
	Password string `json:"password" binding:"required"`
}

type CreateUserRequest struct {
	BaseUser
	PwdType
}
type IdType struct {
	ID int `json:"id" binding:"required"`
}

type EditUserRequest struct {
	IdType
	BaseUser
	IsRecorder uint `json:"is_recorder"`
}

type ChPwd struct {
	IdType
	PwdType
}

func checkEmailFormat(email string) error {
	if err := validate.Var(email, "email"); err != nil {
		return errors.New("email " + consts.ErrFormat)
	}
	return nil
}

func validateCreateUser(u CreateUserRequest) error {
	checkMap := map[string]string{
		u.Username:   "username",
		u.Password:   "password",
		u.Department: "department",
		u.RealName:   "real_name",
		u.Email:      "email",
	}

	for val, field := range checkMap {
		if val == "" {
			return fmt.Errorf("%s is required", field)
		}
	}

	return checkEmailFormat(u.Email)
}

func validateEditUser(u EditUserRequest) error {
	if u.ID <= 0 {
		return errors.New("ID错误")
	}
	if u.Email != "" {
		return checkEmailFormat(u.Email)
	}
	return nil
}

func CreateUser(g *gin.Context) (bool, string) {
	var u CreateUserRequest
	if err := g.ShouldBindJSON(&u); err != nil {
		return false, consts.ErrParamInvalid + ": " + err.Error()
	}
	if err := validateCreateUser(u); err != nil {
		return false, consts.ErrParamInvalid + ": " + err.Error()
	}

	// 判断是否重名
	var exists bool
	config.DB.Model(&model.CoreAccount{}).
		Select("count(*) > 0").Where("username = ?", u.Username).Find(&exists)
	if exists {
		return false, consts.UserMsg + consts.ErrUserExists
	}

	// 加密密码
	u.Password = factory.DjangoEncrypt(u.Password, string(factory.GetRandom()))
	var user model.CoreAccount
	if err := copier.Copy(&user, &u); err != nil {
		return false, consts.ErrOperate + ": " + err.Error()
	}
	user.IsRecorder = 2

	// 添加数据库,提交事务
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		if err := tx.Create(&model.CoreGrained{
			UserId: user.ID,
			Group:  common.EmptyGroup(),
		}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return false, fmt.Sprintf("%v: %v", consts.ErrOperate, err)
	}

	return true, consts.MsgUserCreate
}

func EditUser(g *gin.Context) (bool, string) {
	var u EditUserRequest
	if err := g.ShouldBindJSON(&u); err != nil {
		return false, consts.ErrParamInvalid + ": " + err.Error()
	}
	if err := validateEditUser(u); err != nil {
		return false, consts.ErrParamInvalid + ": " + err.Error()
	}

	// 前端传哪个值，改哪个值
	m := common.RemoveZeroValues(common.StructToMap(u))

	// 判断只能更改的字段
	if len(m) > 0 {
		if err := config.DB.Model(model.CoreAccount{}).Where("id = ?", u.ID).Updates(m).Error; err != nil {
			return false, consts.ErrOperate + ": " + err.Error()
		}
	}
	return true, consts.MsgUserUpdate
}

func ResetPwdUser(g *gin.Context) (bool, string) {
	var u ChPwd
	if err := g.ShouldBindJSON(&u); err != nil {
		return false, consts.ErrParamInvalid + ": " + err.Error()
	}
	if err := config.DB.Model(model.CoreAccount{}).Where("id = ?", u.ID).
		Updates(model.CoreAccount{Password: factory.DjangoEncrypt(u.Password, string(factory.GetRandom()))}).Error; err != nil {
		return false, consts.ErrOperate + ": " + err.Error()
	}
	return true, consts.MsgUserUpdate
}

func EditPayloadUser(g *gin.Context) (bool, string) {
	var u ChPwd
	if err := g.ShouldBindJSON(&u); err != nil {
		return false, consts.ErrParamInvalid + ": " + err.Error()
	}
	if err := config.DB.Model(model.CoreAccount{}).Where("id = ?", u.ID).
		Updates(model.CoreAccount{Password: factory.DjangoEncrypt(u.Password, string(factory.GetRandom()))}).Error; err != nil {
		return false, consts.ErrOperate + ": " + err.Error()
	}
	return true, consts.UserMsg + consts.MsgUpdateSuccess
}
