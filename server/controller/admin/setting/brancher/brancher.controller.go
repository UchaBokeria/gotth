package brancher

import (
	"fmt"
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
	"net/http"
	"strconv"
)

func Brancher(ctx *controller.Context) error {
	var Body BrancherDto

	if err := ctx.Bind(&Body); err != nil {
		return ctx.String(http.StatusBadRequest, "Parameters Binding Problem: " + err.Error())
	}

	var branch model.Branches
	ID, _ := strconv.Atoi(Body.ID)

	match := storage.DB.Last(&branch, uint(ID))
	if match.Error != nil {
		fmt.Print("No branch as ", uint(ID))
		return ctx.String(http.StatusNotFound, "")
	}
	DistrictID, _ := strconv.Atoi(Body.DistrictID)
	result := storage.DB.Model(branch).Updates(&model.Branches{
		Map: Body.Map,
		Name: Body.Name,
		PhoneNumber: Body.PhoneNumber,
		DistrictID: DistrictID,
	})

	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	return refresh(ctx)
}

func BrancherNew(ctx *controller.Context) error {
	var Body BrancherDto

	if err := ctx.Bind(&Body); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}

	DistrictID, _ := strconv.Atoi(Body.DistrictID)
	result := storage.DB.Create(&model.Branches{
		Map: Body.Map,
		Name: Body.Name,
		PhoneNumber: Body.PhoneNumber,
		DistrictID: DistrictID,
	})
						 
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	return refresh(ctx)
}

func BrancherRemove(ctx *controller.Context) error {
	var branch model.Branches
	var ID struct{ ID string `param:"id"`}

	if err := ctx.Bind(&ID); err != nil {
		return ctx.String(http.StatusBadRequest, "Parameters Binding Problem: " + err.Error())
	}

	storage.DB.First(&branch, ID)
	result := storage.DB.Delete(&branch)
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	return refresh(ctx)
}

func refresh(ctx *controller.Context) error {
	var Cities []model.Cities
	storage.DB.Find(&Cities)

	var branchs []model.Branches
	storage.DB.Order("created_at desc").Preload("District.City").Find(&branchs)
	return ctx.Html(view.Brancher(branchs, Cities))
}

func Districts(ctx *controller.Context) error {
	var Params struct{ DistrictDom string `param:"district"`; CityID string `param:"id"`; Default string `param:"default"`;}

	if err := ctx.Bind(&Params); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}

	CityID, _ := strconv.Atoi(Params.CityID)

	var Districts []model.Districts
	storage.DB.Where(model.Districts{CityID: CityID}).Find(&Districts)

	var District model.Districts
	ID, _ := strconv.Atoi(Params.Default)
	storage.DB.Last(&District, ID)
	
	var placeholder string = District.Display_name
	if placeholder == "" {
		placeholder = "უბანი"
	}

	return ctx.Html(view.LoadDistricts(Districts, placeholder, Params.DistrictDom, Params.Default))
}
