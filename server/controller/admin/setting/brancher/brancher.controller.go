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

	var branchs []model.Branches
	storage.DB.Order("created_at desc").Preload("District.City").Find(&branchs)
	return ctx.Html(view.Brancher(branchs, FindCities(), FindDistricts()))
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

	var branchs []model.Branches
	storage.DB.Order("created_at desc").Preload("District.City").Find(&branchs)
	return ctx.Html(view.Brancher(branchs, FindCities(), FindDistricts()))
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

	var branchs []model.Branches
	storage.DB.Order("created_at desc").Preload("District.City").Find(&branchs)
	return ctx.Html(view.Brancher(branchs, FindCities(), FindDistricts()))
}

func FindCities() []model.Cities{
	var Cities []model.Cities
	storage.DB.Find(&Cities)

	return Cities
}

func FindDistricts() []model.Districts{
	var Districts []model.Districts
	storage.DB.Find(&Districts)

	return Districts
}
