package products

import (
	"main/cmd/seed/news"
	"main/server/common/storage"
	"main/server/model"
)

var Approvals = []model.Product_approvals{
	{
		ProductsID: 1,
		Name: "",
		Slug: "",
	},
} 

func Approval() {
	for _, row := range Seed { 
		storage.DB.Create(&row)
			
		storage.DB.
			Model(&row).
			Association("News").
			Append(&news.Seed[1], &news.Seed[2], &news.Seed[3], &news.Seed[4])
	}
}