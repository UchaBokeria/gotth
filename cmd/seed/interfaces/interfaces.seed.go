package interfaces

import (
	"main/cmd/seed/news"
	"main/server/common/storage"
	"main/server/model"
)

var Interfaces = []model.Interface{
	{
		Ver:  1,
		Name: "Configuration",
		Slug: "configuration",
	},
} 

func Populate() {
	for _, row := range Interfaces { 
		storage.DB.Create(&row)
			
		storage.DB.
			Model(&row).
			Association("News").
			Append(&news.Seed[1], &news.Seed[2], &news.Seed[3], &news.Seed[4])
	}
}