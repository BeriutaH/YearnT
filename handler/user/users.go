package user

import (
	"Yearn-go/config"
	"Yearn-go/consts"
	"Yearn-go/handler/common"
	"Yearn-go/model"
	"Yearn-go/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

var userActionMap = map[string]common.HandlerFunc{
	"create": CreateUser,
	"edit":   EditUser,
	"reset":  ResetPwdUser,
	"policy": EditPayloadUser,
}

// GetUserInfo 获取全部的用户列表
func GetUserInfo(g *gin.Context) {
	var userList []model.CoreAccount
	if err := config.DB.Omit("password").Find(&userList).Error; err != nil {
		utils.Fail(g, consts.ErrGetUser)
		return
	}
	utils.Ok(g, userList)
}

// SelectUserInfo 分页查询
func SelectUserInfo(g *gin.Context) {
	common.HandlePaging[[]model.CoreAccount](g, common.UserSensitiveFields, common.UserQueryableFields)
}

// ManageUserCreateOrEdit 用户 添加，修改，重置密码
func ManageUserCreateOrEdit(g *gin.Context) {
	common.ActionDispatcher(g, userActionMap)
}

func DeleteUserById(g *gin.Context) {
	var uId IdType
	if err := g.ShouldBindJSON(&uId); err != nil {
		utils.Fail(g, fmt.Sprint(consts.ErrParamInvalid, err))
		return
	}
	// 物理删除
	if err := config.DB.Unscoped().Delete(&model.CoreAccount{}, uId.ID).Error; err != nil {
		utils.Fail(g, fmt.Sprint(consts.ErrOperate, err))
		return
	}
	utils.Ok(g, consts.MsgDeleteSuccess)
}
