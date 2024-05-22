package categories

import (
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
)

func index(ctx *controller.Context) error {
	var Categories []model.Categories
	storage.DB.Preload("Icon").Where(&model.Categories{Public: true}).Find(&Categories)
	return ctx.Html(view.Categories(Categories))
}