package categories

import (
	"main/server/common/storage"
	"main/server/model"
)

var Seed = []model.Categories{
	{
		IconID: 15,
		Name: "ავტომობილი",
		Slug: "Automobile",
		Filters: Automobile,
	},
	{
		IconID: 16,
		Name: "მოტო / კვად /კარტინგი",
		Slug: "Moto / quad / Karting",
		Filters: MotoQuadKarting,
	},
	{
		IconID: 17,
		Name: "ტრანსპორტი / მძიმე ტექნიკა",
		Slug: "Transport / Heavy equipment",
		Filters: TransportHeavyEquipment,
	},
	{
		IconID: 18,
		Name: "მეურნეობა",
		Slug: "Farming",
		Filters: Farming,
	},
	{
		IconID: 19,
		Name: "გადამცემთა კოლოფის ღერძი",
		Slug: "Gearboxes / beam axles",
		Filters: GearboxesBeamAxles,
	},
	{
		IconID: 20,
		Name: "მოვლა / გაწმენდა",
		Slug: "Upkeep and cleaning",
		Filters: UpkeepAndCleaning,
	},
	{
		IconID: 21,
		Name: "ნადირობა / იახტინგი",
		Slug: "Sailing / Yachting",
		Filters: SailingAndYachting,
	},
	{
		IconID: 22,
		Name: "მსუბუქი თვითმფრინავები",
		Slug: "Light planes",
		Filters: LightPlanes,
	},
	{
		IconID: 23,
		Name: "სპეციალობები",
		Slug: "Specialities",
		Filters: Specialities,
	},
}

func Populate() {
	for _, row := range Seed { storage.DB.Create(&row) }
}