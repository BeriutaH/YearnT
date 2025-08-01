package user

import (
	"Yearn-go/config"
	"Yearn-go/consts"
	"Yearn-go/handler/common"
	"Yearn-go/model"
	"Yearn-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ActionUserBase struct {
	Action string `json:"action" binding:"required,oneof=add edit reset"`
}

func GetUserInfo(g *gin.Context) {
	var userList []model.CoreAccount
	if err := config.DB.Omit("password").Find(&userList).Error; err != nil {
		utils.Fail(g, consts.ErrGetUser)
		return
	}
	utils.Ok(g, userList)
}

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

func ManageUserCreateOrEdit(g *gin.Context) {
	var action ActionUserBase
	if err := g.ShouldBindBodyWith(&action, binding.JSON); err != nil {
		utils.Fail(g, consts.ErrParamInvalid+": "+err.Error())
		return
	}
	var success bool
	var msg string
	// 获取参数，判断操作类型
	switch action.Action {
	case "add":
		success, msg = CreateUser(g)
	case "edit":
		success, msg = EditUser(g)
	case "reset":
		success, msg = ResetPwdUser(g)
	case "policy":
		success, msg = EditPayloadUser(g)

	}

	utils.HandleResult(g, success, msg)
}

func InterfaceTestF(g *gin.Context) {
	idUser := g.Param(consts.UrlOp)
	println(idUser)
	utils.Ok(g, "DELETE验证成功")
}
