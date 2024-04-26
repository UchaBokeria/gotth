package interfaces

import (
	"main/server/common/storage"
	"main/server/model"
)

func Init() {
	for _, row := range []model.Interface{
		{
			Ver:  1,
			Name: "Landing",
			Slug: "landing",
		},
    } { storage.DB.Create(&row) }
}