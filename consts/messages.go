package consts

const (
	MsgSuccess               = "成功"
	MsgFal                   = "失败"
	UserMsg                  = "用户"
	MsgRegisterSuccess       = "注册" + MsgSuccess
	MsgUpdateSuccess         = "更新" + MsgSuccess
	MsgDeleteSuccess         = "删除" + MsgSuccess
	MsgCreateSuccess         = "创建" + MsgSuccess
	ErrUserNotFound          = "不存在"
	ErrUserExists            = "已存在"
	ErrUnauthorized          = "非法越权操作"
	ErrInvalidPassword       = "密码错误"
	ErrInvalidInput          = "无效输入"
	ErrMissingOrInvalidToken = "Token缺少或格式错误"
	ErrInvalidToken          = "无效的Token"
	ErrParamInvalid          = "参数错误"
	ErrFormat                = "格式错误"
	ErrOperate               = "操作" + MsgFal
	ErrGetUser               = "获取" + UserMsg + MsgFal
)
