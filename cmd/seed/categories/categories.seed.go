package categories

import (
	"main/server/common/storage"
	"main/server/model"
)

var Seed = []model.Categories{
	{
		IconID: 15,
		Public: true,
		Name: "ავტომობილი",
		Slug: "Automobile",
		Filters: Automobile,
	},
	{
		IconID: 16,
		Public: true,
		Name: "მოტო / კვად /კარტინგი",
		Slug: "Moto / quad / Karting",
		Filters: MotoQuadKarting,
	},
	{
		IconID: 17,
		Public: true,
		Name: "ტრანსპორტი / მძიმე ტექნიკა",
		Slug: "Transport / Heavy equipment",
		Filters: TransportHeavyEquipment,
	},
	{
		IconID: 18,
		Public: true,
		Name: "მეურნეობა",
		Slug: "Farming",
		Filters: Farming,
	},
	{
		IconID: 19,
		Public: true,
		Name: "გადამცემთა კოლოფის ღერძი",
		Slug: "Gearboxes / beam axles",
		Filters: GearboxesBeamAxles,
	},
	{
		IconID: 20,
		Public: false,
		Name: "მოვლა / გაწმენდა",
		Slug: "Upkeep and cleaning",
		Filters: UpkeepAndCleaning,
	},
	{
		IconID: 21,
		Public: true,
		Name: "ნადირობა / იახტინგი",
		Slug: "Sailing / Yachting",
		Filters: SailingAndYachting,
	},
	{
		IconID: 22,
		Public: false,
		Name: "მსუბუქი თვითმფრინავები",
		Slug: "Light planes",
		Filters: LightPlanes,
	},
	{
		IconID: 23,
		Public: true,
		Name: "სპეციალობები",
		Slug: "Specialities",
		Filters: Specialities,
	},
}

func Populate() {
	for _, row := range Seed { storage.DB.Create(&row) }
}