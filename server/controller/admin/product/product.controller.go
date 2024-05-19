package product

import (
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
)

func index(ctx *controller.Context) error {
	var Products []model.Products

	storage.DB.Preload("Thumbnail").
			   Preload("Packing").
			   Preload("Approvals").
			   Preload("Properties").
			   Preload("Specifications").
			   Find(&Products)
	return ctx.Html(view.Product(Products))
}