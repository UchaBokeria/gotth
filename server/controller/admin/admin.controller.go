package admin

import (
	"main/build/view"
	"main/server/common/controller"
	"net/http"

	"github.com/a-h/templ"
)

func index(ctx *controller.Context) error {
	return ctx.Html(view.Admin(templ.NopComponent))
}

func upload(ctx *controller.Context) error {
	return ctx.String(http.StatusOK, "")
}