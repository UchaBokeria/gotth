package products

import (
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
	"strconv"
)

func index(ctx *controller.Context) error {
	var Filters FiltersQuery
	var Category model.Categories
	var Categories []model.Categories

	if ctx.QueryParam("category") == "" { return ctx.Html(view.Wildcard()) }

	ctx.Bind(&Filters)
	ID, _ := strconv.Atoi(Filters.Category)

	storage.DB.Find(&Categories, &model.Categories{Public: true})
	storage.DB.Preload("Filters.Options").
			   Where(&model.Categories{Public: true}).
			   First(&Category, ID)


	return ctx.Html(view.ProductLayout(Category, Categories))
}

func list(ctx *controller.Context) error {
	var Filters FiltersQuery
	var Products []model.Products

	if ctx.QueryParam("category") == "" {
		return ctx.Html(view.Wildcard())
	}

	ctx.Bind(&Filters)
	NextPage := strconv.Itoa(ctx.Page() + 1)
	CategoryID, _ := strconv.Atoi(Filters.Category)

	storage.DB.Scopes(storage.Paginate(ctx)).
			   Where(&model.Products{CategoryID: CategoryID, Public: true}).
			   Preload("Thumbnail").
			   Preload("Packing").
			   Preload("Approvals").
			   Preload("Properties").
			   Preload("Specifications").
			   Find(&Products)

	if len(Products) == 0 { return ctx.String(404, "") }
	return ctx.Html(view.Products(NextPage, Filters.Category, Products))
}

func detail(ctx *controller.Context) error {
	var Filters ProductDetail
	var Product model.Products

	ctx.Bind(&Filters)
	ID, _ := strconv.Atoi(Filters.ID)

	storage.DB.Preload("Thumbnail").
			   Preload("Packing").
			   Preload("Approvals").
			   Preload("Properties").
			   Preload("Specifications").
			   Find(&Product, ID)

	return ctx.Html(view.ProductDetail(Product))
}