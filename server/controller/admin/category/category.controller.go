package category

import (
	"fmt"
	"main/build/view"
	"main/server/common/controller"
	uploader "main/server/common/helpers"
	"main/server/common/storage"
	"main/server/model"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
)

func index(ctx *controller.Context) error {
	var Categories []model.Categories
	storage.DB.Order("created_at desc").Preload("Icon").Find(&Categories)
	return ctx.Html(view.Category(Categories))
}

func indexByID(ctx *controller.Context) error {
	var Categorie model.Categories
	var ID struct{ ID string `param:"id"`}

	if err := ctx.Bind(&ID); err != nil {
		return ctx.String(http.StatusBadRequest, "Parameters Binding Problem: " + err.Error())
	}

	result := storage.DB.Preload("Icon").First(&Categorie, ID)
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	return ctx.Html(view.UpdateCategory(Categorie))
}

func CategoryNew(ctx *controller.Context) error {
	var Body CategoryDto

	if err := ctx.Bind(&Body); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}

	Public := true
	if Body.Public == "false" { Public = false }
	Parameters := model.Categories{Name: Body.Name, Public: Public}

	file, err := ctx.FormFile("Icon")
	if file != nil && err == nil {
		Upload := uploader.File(file)
		if !Upload.Success {
			fmt.Print(Upload.Message) 
			return ctx.String(http.StatusBadRequest, Upload.Message)
		}
		Parameters.IconID = Upload.ID
	} else {
		fmt.Print(err.Error()) 
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	result := storage.DB.Create(&Parameters)
						 
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Categories []model.Categories
	storage.DB.Order("created_at desc").Preload("Icon").Find(&Categories)
	return ctx.Html(view.Category(Categories))
}

func CategoryUpdate(ctx *controller.Context) error {
	var Body CategoryDto

	if err := ctx.Bind(&Body); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}
	var Categorie model.Categories
	ID, _ := strconv.Atoi(Body.ID)

	match := storage.DB.Last(&Categorie, uint(ID))
	if match.Error != nil {
		fmt.Print("No Faq as ", uint(ID))
		return ctx.String(http.StatusNotFound, "")
	}

	Public := true
	if Body.Public == "false" { Public = false }
	Parameters := map[string]interface{}{"Name": Body.Name, "Public": Public}

	file, err := ctx.FormFile("Icon")
	if file != nil && err == nil {
		Upload := uploader.File(file)
		if !Upload.Success {
			fmt.Print(Upload.Message) 
			return ctx.String(http.StatusBadRequest, Upload.Message)
		}
		Parameters["IconID"] = Upload.ID
	}

	result := storage.DB.Model(Categorie).Updates(&Parameters)
						 
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Categories []model.Categories
	storage.DB.Order("created_at desc").Preload("Icon").Find(&Categories)
	return ctx.Html(view.Category(Categories))
}

func CategoryStatus(ctx *controller.Context) error {
	var Categorie model.Categories
	var ID struct{ ID string `param:"id"`}

	if err := ctx.Bind(&ID); err != nil {
		return ctx.String(http.StatusBadRequest, "Parameters Binding Problem: " + err.Error())
	}

	storage.DB.First(&Categorie, ID)
	result := storage.DB.Model(Categorie).Updates(&map[string]interface{}{
		"Public": !Categorie.Public,
	})

	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Categories []model.Categories
	storage.DB.Order("created_at desc").Preload("Icon").Find(&Categories)
	return ctx.Html(view.Category(Categories))
}

func CategoryRemove(ctx *controller.Context) error {
	var Categorie model.Categories
	var ID struct{ ID string `param:"id"`}

	if err := ctx.Bind(&ID); err != nil {
		return ctx.String(http.StatusBadRequest, "Parameters Binding Problem: " + err.Error())
	}

	storage.DB.First(&Categorie, ID)
	result := storage.DB.Delete(&Categorie)
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	return ctx.Html(templ.NopComponent)
}
