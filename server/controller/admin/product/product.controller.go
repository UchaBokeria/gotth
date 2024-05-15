package product

import (
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
)

func index(ctx *controller.Context) error {
	var Products []model.Products

	storage.DB.Preload("TechnicalSheet").
			   Preload("Thumbnail").
			   Preload("Packing").
			   Preload("Approvals").
			   Preload("Specifications").
			   Find(&Products)
	return ctx.Html(view.Product(Products))
}