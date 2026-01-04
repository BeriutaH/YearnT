package user

import (
	"Yearn-go/restful"
)

func SuperUserApi() restful.RouteGroup {
	return restful.RouteGroup{
		Get:    GetUserInfo,
		Post:   ManageUserCreateOrEdit,
		Put:    SelectUserInfo,
		Delete: InterfaceTestF,
	}
}
