package user

import (
	"Yearn-go/restful"
)

func SuperUserApi() restful.RestfulAPI {
	return restful.RestfulAPI{
		Put:    InterfaceTestT,
		Post:   InterfaceTes3tS,
		Delete: InterfaceTestF,
		Get:    InterfaceTestO,
	}
}
