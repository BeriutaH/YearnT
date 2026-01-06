package user

import (
	"Yearn-go/config"
	"Yearn-go/consts"
	"Yearn-go/handler/common"
	"Yearn-go/model"
	"Yearn-go/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	var req common.QueryRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		utils.Fail(g, consts.ErrParamInvalid+": "+err.Error())
		return
	}

	p := new(common.PageList[[]model.CoreAccount])
	p.ToPageInfo(req.PageInfo).Paging().Query(
		common.QmiFilters(common.UserSensitiveFields),
		common.ApplyFilters(common.UserQueryableFields, req.Filters),
	)
	utils.Ok(g, p.ToMessage())
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
	userId := uId.ID

	// 假设 userId 是要删除的用户 ID
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 彻底删除关联的 CoreGrained 数据
		if err := tx.Unscoped().Where("user_id = ?", userId).Delete(&model.CoreGrained{}).Error; err != nil {
			return err
		}

		// 2. 彻底删除用户主体
		result := tx.Unscoped().Delete(&model.CoreAccount{}, userId)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	// 统一处理事务结果
	if err != nil {
		utils.Fail(g, fmt.Sprint(consts.ErrOperate, err))
		return
	}
	utils.Ok(g, consts.MsgDeleteSuccess)
}
