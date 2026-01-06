package flow

import "Yearn-go/restful"

func TplRestApis() restful.RouteGroup {
	return restful.RouteGroup{
		Get:    TplGetAPis,
		Post:   TplPostSourceTemplate,
		Put:    EditSourceTemplateInfo,
		Delete: DeleteSourceTemplateInfo,
	}
}
