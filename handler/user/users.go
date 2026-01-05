package user

import (
	"Yearn-go/config"
	"Yearn-go/consts"
	"Yearn-go/handler/common"
	"Yearn-go/model"
	"Yearn-go/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

// ProcessorUser 处理器类型
type ProcessorUser func(*gin.Context) (bool, string)

var actionMap = map[string]ProcessorUser{
	"add":    CreateUser,
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
	action := g.Query("action")
	// 操作映射
	handler, ok := actionMap[action]
	if !ok {
		utils.Fail(g, "不支持的操作类型: "+action)
		return
	}

	success, msg := handler(g)
	utils.HandleResult(g, success, msg)
}

func DeleteUserById(g *gin.Context) {
	var uId IdType
	if err := g.ShouldBindJSON(&uId); err != nil {
		utils.Fail(g, consts.ErrParamInvalid+": "+err.Error())
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
		// 打印具体错误以便调试，返回给前端通用失败提示
		log.Printf("Delete user error: %v", err)
		utils.Fail(g, consts.ErrOperate)
		return
	}
	utils.Ok(g, nil)
}
