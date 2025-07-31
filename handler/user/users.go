package user

import (
	"Yearn-go/config"
	"Yearn-go/consts"
	"Yearn-go/factory"
	"Yearn-go/models"
	"Yearn-go/utils"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(g *gin.Context) {
	var userList []models.CoreAccount
	if err := config.DB.Omit("password").Find(&userList).Error; err != nil {
		utils.Fail(g, consts.ErrGetUser)
		return
	}
	utils.Ok(g, userList)
}

func SelectUserInfo(g *gin.Context) {
	var req factory.QueryRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		utils.Fail(g, consts.ErrParamInvalid+": "+err.Error())
		return
	}
	// 初始化查询
	db := config.DB.Model(&models.CoreAccount{}).Omit("password")
	allowedFields := map[string]bool{
		"username":  true,
		"email":     true,
		"real_name": true,
	}
	db = req.ApplyFilters(db, allowedFields)
	// 分页
	info, _ := factory.Paginate[models.CoreAccount](req.Page, req.PageSize, db)
	utils.Ok(g, info)
}

func ManageUserCreateOrEdit(g *gin.Context) {
	var success bool
	var msg string
	// 获取参数，判断操作类型
	switch g.Query(consts.UrlOp) {
	case "add":
		success, msg = CreateUser(g)
	case "edit":
		success, msg = CreateUser(g)
	}
	utils.HandleResult(g, success, msg)
}

func InterfaceTestF(g *gin.Context) {
	idUser := g.Param(consts.UrlOp)
	println(idUser)
	utils.Ok(g, "DELETE验证成功")
}
