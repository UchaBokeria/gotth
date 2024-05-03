package products

import (
	"main/server/common/storage"
	"main/server/model"
)

var Seed = []model.Products{
	{
		Name: "",
		Slug: "",
		Description: "",
		TechnicalSheetID: 1,
		ThumbnailID: 1,
	},
}

func Populate() {
	for _, row := range Seed { storage.DB.Create(&row) }
}