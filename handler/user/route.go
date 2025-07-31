package user

import (
	"Yearn-go/restful"
)

func SuperUserApi() restful.RestfulAPI {
	return restful.RestfulAPI{
		Get:    GetUserInfo,
		Post:   ManageUserCreateOrEdit,
		Put:    SelectUserInfo,
		Delete: InterfaceTestF,
	}
}
