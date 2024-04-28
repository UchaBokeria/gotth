package branches

import (
	"main/server/common/storage"
	"main/server/model"
)

var City = []model.Cities{
	{
		Name: "თბილისი",
		Slug: "tbilisi",
	},
	{
		Name: "ბათუმი",
		Slug: "batumi",
	},
	{
		Name: "ქუთაისი",
		Slug: "quteisi",
	},
	{
		Name: "ყაზბეგი/სტეფანწმინდა",
		Slug: "yazbegi",
	},
	{
		Name: "მესტია",
		Slug: "mestia",
	},
	{
		Name: "თელავი",
		Slug: "telavi",
	},
}

func Cities() {
	for _, row := range City { storage.DB.Create(&row) }
}