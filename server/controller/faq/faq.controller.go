package faq

import (
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
)

func index(ctx *controller.Context) error {
	var Faq []model.Faq
	storage.DB.Find(&Faq)
	return ctx.Html(view.Faq(Faq))
}