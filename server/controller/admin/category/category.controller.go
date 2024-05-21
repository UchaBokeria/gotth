package category

import (
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
)

func index(ctx *controller.Context) error {
	var Categories []model.Categories
	storage.DB.Order("created_at desc").Preload("Icon").Find(&Categories)
	return ctx.Html(view.Category(Categories))
}

func create(ctx *controller.Context) error {
	var Categories []model.Categories
	storage.DB.Preload("Icon").Find(&Categories)
	return ctx.Html(view.Category(Categories))
}

func update(ctx *controller.Context) error {

	var Categories []model.Categories
	storage.DB.Order("created_at desc").Preload("Icon").Find(&Categories)
	return ctx.Html(view.Category(Categories))
}

func remove(ctx *controller.Context) error {
	var Categories []model.Categories
	storage.DB.Order("created_at desc").Preload("Icon").Find(&Categories)
	return ctx.Html(view.Category(Categories))
}
