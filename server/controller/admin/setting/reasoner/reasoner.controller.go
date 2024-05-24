package reasoner

import (
	"fmt"
	"main/build/view"
	"main/server/common/controller"
	uploader "main/server/common/helpers"
	"main/server/common/storage"
	"main/server/model"
	"net/http"
	"strconv"
)

func Reasoner(ctx *controller.Context) error {
	var Body ReasonerDto

	if err := ctx.Bind(&Body); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}

	ID, _ := strconv.Atoi(Body.ID)
	var slide model.Interface_reasons

	match := storage.DB.Last(&slide, uint(ID))
	if match.Error != nil {
		fmt.Print("No Reason as ", uint(ID))
		return ctx.String(http.StatusNotFound, "")
	}

	Parameters := model.Interface_reasons{
		Name: Body.Name,
		Title: Body.Title,
		Desc: Body.Desc,
		Url: Body.Url,
	}

	file, err := ctx.FormFile("Icon")
	if file != nil && err == nil {
		Upload := uploader.File(file)
		if !Upload.Success {
			fmt.Print(Upload.Message) 
			return ctx.String(http.StatusBadRequest, Upload.Message)
		}
		
		Parameters.IconID = &Upload.ID
	}

	result := storage.DB.Model(slide).Updates(&Parameters)
	if result.Error != nil { return ctx.String(http.StatusBadRequest, "") }

	var Reasons []model.Interface_reasons
	storage.DB.Order("interface_reasons.created_at desc").Preload("Icon").Find(&Reasons)

	return ctx.Html(view.Reasoners(Reasons))
}

func ReasonerNew(ctx *controller.Context) error {
	var Body ReasonerDto

	if err := ctx.Bind(&Body); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}

	var Parameters = model.Interface_reasons{
			InterfaceID: 1,
			Name: Body.Name,
			// Slug: Body.Name,
			Title: Body.Title,
			Desc: Body.Desc,
			Url: Body.Url,
		}

	file, err := ctx.FormFile("Icon")
	if file != nil && err == nil {	
		Upload := uploader.File(file)
		if !Upload.Success {
			fmt.Print("File did not upload -> " + Upload.Message) 
			// return ctx.String(http.StatusBadRequest, Upload.Message)
		} else {
			Parameters.IconID = &Upload.ID
		}
	}

	result := storage.DB.Create(&Parameters)
						 
	if result.Error != nil { return ctx.String(http.StatusBadRequest, result.Error.Error()) }

	var Reasons []model.Interface_reasons
	storage.DB.Order("interface_reasons.created_at desc").Preload("Icon").Find(&Reasons)
	return ctx.Html(view.Reasoners(Reasons))
}

func ReasonerRemove(ctx *controller.Context) error {
	var Reason model.Interface_reasons
	var ID struct{ ID string `param:"id"`}

	if err := ctx.Bind(&ID); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}

	storage.DB.First(&Reason, ID)
	storage.DB.Delete(&Reason)

	var Reasons []model.Interface_reasons
	storage.DB.Order("interface_reasons.created_at desc").Preload("Icon").Find(&Reasons)
	return ctx.Html(view.Reasoners(Reasons))
}
