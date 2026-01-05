package group

import (
	"Yearn-go/restful"
)

func GroupsApis() restful.RouteGroup {
	return restful.RouteGroup{
		Post:   SuperGroupUpdate,
		Put:    SuperGroup,
		Delete: SuperClearUserRule,
	}
}
