package news

import (
	"main/server/common/storage"
	"main/server/model"
)

var Type = []model.News_types{
	{
		Name: "ფოტო",
		Slug: "photo",
	},
	{
		Name: "ვიდეო",
		Slug: "video",
	},
	{
		Name: "ბლოგი",
		Slug: "blog",
	},
} 

func Types() {
	for _, row := range Type { storage.DB.Create(&row) }
}