package news

import (
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
)

func index(ctx *controller.Context) error {
	var News []model.News
	var Types []model.News_types

	var Where = &model.News{Public: true}

	if ctx.QueryParam("type") != "" {
		var filter Filters
		ctx.Bind(&filter)
		Where.TypeID = filter.Type
	}

	storage.DB.Find(&Types)
	storage.DB.Where(Where).Preload("Thumbnail").Find(&News)

	return ctx.Html(view.News(News, Types, ctx.QueryParam("type")))
}

func detail(ctx *controller.Context) error {
	return ctx.Html(view.NewsDetails())
}