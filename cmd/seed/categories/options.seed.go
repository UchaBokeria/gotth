package categories

import (
	"main/server/common/storage"
	"main/server/model"
)

var Options = []model.Category_filters_option{
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

func Option() {
	for _, row := range Seed { storage.DB.Create(&row) }
}