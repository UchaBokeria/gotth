package branches

import (
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
	"net/http"
	"strconv"
)

func index(ctx *controller.Context) error {
	var Query BranchesDto
	var Citie = "ქალაქი"
	var District = "უბანი"
	BranchesWhere := model.Branches{}
	DistrictsWhere := model.Districts{}

	if err := ctx.Bind(&Query); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if Query.Citie != "" {
		ID, _ := strconv.Atoi(Query.Citie)
		var Citier model.Cities
		storage.DB.Last(&Citier, ID)
		Citie = Citier.Display_name
		DistrictsWhere.CityID = ID
		BranchesWhere.District.CityID = ID
	}

	if Query.District != "" {
		ID, _ := strconv.Atoi(Query.District)
		var Districter model.Districts
		storage.DB.Last(&Districter, ID)
		District = Districter.Display_name
		BranchesWhere.DistrictID = ID
	}

	var Branches []model.Branches
	storage.DB.Where(&BranchesWhere).Preload("District.City").Preload("Shifts").Find(&Branches)

	var Cities []model.Cities
	storage.DB.Find(&Cities)

	var Districts []model.Districts
	storage.DB.Where(&DistrictsWhere).Find(&Districts)

	return ctx.Html(view.Branches(Branches, Cities, Districts, Citie, Query.Citie, District, Query.District))
}
