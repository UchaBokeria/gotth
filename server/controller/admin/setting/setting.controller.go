package setting

import (
	"main/build/view"
	"main/server/common/controller"
)

func index(ctx *controller.Context) error {

	return ctx.Html(view.Setting())
}