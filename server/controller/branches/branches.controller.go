package branches

import (
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
)

func index(ctx *controller.Context) error {
	var Branches []model.Branches
	storage.DB.Preload("City").Find(&Branches)
	return ctx.Html(view.Branches(Branches))
}