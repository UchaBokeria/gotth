package news

import (
	"strconv"

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
	storage.DB.
		Order("news.created_at desc").
		Where(Where).
		Preload("Thumbnail").
		Find(&News)

	return ctx.Html(view.News(News, Types, ctx.QueryParam("type")))
}

func detail(ctx *controller.Context) error {
	var News model.News
	
	var Where = &model.News{Public: true}

	if ctx.Param("ID") != "" {
		var filter Detail
		ctx.Bind(&filter)
		num, err := strconv.ParseUint(filter.ID, 10, 0)
		if err == nil {
			Where.ID = uint(num)
		}
	}

	storage.DB.Where(Where).Preload("Thumbnail").Last(&News)

	return ctx.Html(view.NewsDetails(News))
}