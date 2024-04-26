package news

import (
	"main/build/view"
	"main/server/common/controller"
)

func index(ctx *controller.Context) error {
	return ctx.Html(view.News())
}

func detail(ctx *controller.Context) error {
	return ctx.Html(view.NewsDetails())
}