package db

import (
	"Yearn-go/restful"
)

func ManageDbApis() restful.RouteGroup {
	return restful.RouteGroup{
		Post:   ManageDBCreateOrEdit,
		Delete: SuperDeleteSource,
		Put:    SuperFetchSource,
	}
}
