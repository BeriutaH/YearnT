package user

import (
	"Yearn-go/restful"
)

func SuperUserApi() restful.RestfulAPI {
	return restful.RestfulAPI{
		Get:    InterfaceTestO,
		Post:   ManageUserCreateOrEdit,
		Put:    InterfaceTestT,
		Delete: InterfaceTestF,
	}
}
