package terms

import (
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
)

func index(ctx *controller.Context) error {
	var About model.Interface_about
	result := storage.DB.Last(&About)

	if result.Error != nil {
		return ctx.Html(view.ErrorPage())
	}

	return ctx.Html(view.Terms(About.Terms))
}