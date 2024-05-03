package categories

import (
	"main/server/common/storage"
	"main/server/model"
)

var Filters = []model.Category_filters{
	{
		IconID: 15,
		Name: "ავტომობილი",
	},
	{
		IconID: 16,
		Name: "მოტო / კვად /კარტინგი",
	},
	{
		IconID: 17,
		Name: "ტრანსპორტი / მძიმე ტექნიკა",
	},
}

func Filter() {
	for _, row := range Seed { storage.DB.Create(&row) }
}