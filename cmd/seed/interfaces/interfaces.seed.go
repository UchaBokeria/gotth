package interfaces

import (
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
	for _, row := range Interfaces { storage.DB.Create(&row) }
}