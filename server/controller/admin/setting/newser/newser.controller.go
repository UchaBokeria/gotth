package newser

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

func Newser(ctx *controller.Context) error {
	var Body NewserDto

	if err := ctx.Bind(&Body); err != nil {
		return ctx.String(http.StatusBadRequest, "Parameters Binding Problem: " + err.Error())
	}

	var New model.News
	ID, _ := strconv.Atoi(Body.ID)

	match := storage.DB.Last(&New, uint(ID))
	if match.Error != nil {
		fmt.Print("No Faq as ", uint(ID))
		return ctx.String(http.StatusNotFound, "")
	}

	TypeID, _ := strconv.Atoi(Body.TypeID)
	Public := true
	if Body.Public == "false" { Public = false }
	
	Parameters := map[string]interface{}{
		"Title": Body.Title,
		"Body": Body.Body,
		"Public": Public,
		"Url": Body.Url,
		"TypeID": TypeID,
	}

	file, err := ctx.FormFile("Thumbnail")
	if file != nil && err == nil {
		Upload := uploader.File(file)
		if !Upload.Success {
			fmt.Print(Upload.Message) 
			// return ctx.String(http.StatusBadRequest, Upload.Message)
		} else {
			Parameters["ThumbnailID"] = Upload.ID
		}
		
	}

	result := storage.DB.Model(New).Updates(&Parameters)
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Newz []model.News
	storage.DB.Order("created_at desc").Preload("Thumbnail").Find(&Newz)
	return ctx.Html(view.Newser(Newz))
}

func NewserNew(ctx *controller.Context) error {
	var Body NewserDto

	if err := ctx.Bind(&Body); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}

	TypeID, _ := strconv.Atoi(Body.TypeID)
	
	Public := true
	if Body.Public == "false" { Public = false }

	Parameters := model.News{
		Title: Body.Title,
		Body: Body.Body,
		Public: Public,
		Url: Body.Url,
		TypeID: TypeID,
	}

	file, err := ctx.FormFile("Thumbnail")
	if file != nil && err == nil {
		Upload := uploader.File(file)
		if !Upload.Success {
			fmt.Print(Upload.Message) 
			return ctx.String(http.StatusBadRequest, Upload.Message)
		}
		
		Parameters.ThumbnailID = Upload.ID
	} else {
		fmt.Print(err.Error()) 
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	result := storage.DB.Create(&Parameters)
						 
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Newz []model.News
	storage.DB.Order("created_at desc").Preload("Thumbnail").Find(&Newz)
	return ctx.Html(view.Newser(Newz))
}

func NewserRemove(ctx *controller.Context) error {
	var New model.News
	var ID struct{ ID string `param:"id"`}

	if err := ctx.Bind(&ID); err != nil {
		return ctx.String(http.StatusBadRequest, "Parameters Binding Problem: " + err.Error())
	}

	storage.DB.First(&New, ID)
	result := storage.DB.Delete(&New)
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Newz []model.News
	storage.DB.Order("created_at desc").Preload("Thumbnail").Find(&Newz)
	return ctx.Html(view.Newser(Newz))
}
