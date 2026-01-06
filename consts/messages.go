package consts

const (
	// 基础原子
	MsgSuccess  = "成功"
	MsgFal      = "失败"
	UserMsg     = "用户"
	ExistsMsg   = "已存在"
	NotFoundMsg = "不存在"

	// 成功系列 (利用拼接)
	MsgRegisterSuccess = "注册" + MsgSuccess
	MsgUpdateSuccess   = "更新" + MsgSuccess
	MsgDeleteSuccess   = "删除" + MsgSuccess
	MsgCreateSuccess   = "创建" + MsgSuccess

	// 错误系列 - 系统与参数
	ErrParamInvalid  = "参数错误"
	ErrFormat        = "格式错误"
	ErrInvalidInput  = "无效输入"
	ErrOperate       = "操作" + MsgFal
	ErrEncryptFailed = "加密" + MsgFal

	// 错误系列 - 权限与Token
	ErrUnauthorized          = "非法越权操作"
	ErrMissingOrInvalidToken = "Token缺少或格式错误"
	ErrInvalidToken          = "无效的Token"

	// 错误系列 - 用户业务
	MsgUserUpdate      = UserMsg + MsgUpdateSuccess
	MsgUserCreate      = UserMsg + MsgCreateSuccess
	ErrUserNotFound    = UserMsg + "不存在"
	ErrInvalidPassword = "密码错误"
	ErrGetUser         = "获取" + UserMsg + MsgFal

	// 错误系列 - 审批流业务
	ErrFlowInvalid = "关联审批流无效"
)
