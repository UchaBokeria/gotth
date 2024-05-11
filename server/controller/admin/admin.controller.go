package admin

import (
	"main/build/view"
	"main/server/common/controller"

	"github.com/a-h/templ"
)

func index(ctx *controller.Context) error {
	return ctx.Html(view.Admin(templ.NopComponent))
}