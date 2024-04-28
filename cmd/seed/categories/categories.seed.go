package categories

import (
	"main/server/common/storage"
	"main/server/model"
)

var Seed = []model.Categories{
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
	{
		IconID: 18,
		Name: "მეურნეობა",
	},
	{
		IconID: 19,
		Name: "გადამცემთა კოლოფის ღერძი",
	},
	{
		IconID: 20,
		Name: "მოვლა / გაწმენდა",
	},
	{
		IconID: 21,
		Name: "ნადირობა / იახტინგი",
	},
	{
		IconID: 22,
		Name: "მსუბუქი თვითმფრინავები",
	},
	{
		IconID: 23,
		Name: "სპეციალობები",
	},
}

func Populate() {
	for _, row := range Seed { storage.DB.Create(&row) }
}