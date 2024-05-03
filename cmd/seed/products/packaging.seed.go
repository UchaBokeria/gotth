package products

import (
	"main/cmd/seed/news"
	"main/server/common/storage"
	"main/server/model"
)

var Packagings = []model.Product_packaging{
	{
		ProductsID: 1,
		Name: "",
		Slug: "",
		Reference: "",
		Conditioning: "",
		Cardboard: "",
	},
} 

func Packaging() {
	for _, row := range Seed { 
		storage.DB.Create(&row)
			
		storage.DB.
			Model(&row).
			Association("News").
			Append(&news.Seed[1], &news.Seed[2], &news.Seed[3], &news.Seed[4])
	}
}