package product

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

func findProducts(ctx *controller.Context) ([]model.Products, []model.Categories) {
	var Products []model.Products
	
	storage.DB.Scopes(storage.Paginate(ctx)).
				Order("created_at desc").
				Preload("Category").
				Preload("Thumbnail").
				Preload("Packing").
				Preload("Approvals").
				Preload("Properties").
				Preload("Specifications").
				Find(&Products)

	var Categories []model.Categories

	storage.DB.Find(&Categories)
	return Products, Categories
}

func index(ctx *controller.Context) error {

	return ctx.Html(view.Product(findProducts(ctx)))
}

func indexByID(ctx *controller.Context) error {
	var Productie model.Products
	var ID struct{ ID string `param:"id"`}

	if err := ctx.Bind(&ID); err != nil {
		return ctx.String(http.StatusBadRequest, "Parameters Binding Problem: " + err.Error())
	}

	result := storage.DB.Preload("Thumbnail").First(&Productie, ID)
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}


	return ctx.Html(view.UpdateProducts(Productie))
}

func ProductsNew(ctx *controller.Context) error {
	var Body ProductDto

	if err := ctx.Bind(&Body); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}

	Public := true
	if Body.Public == "false" { Public = false }
	CategoryID, _ := strconv.Atoi(Body.CategoryID)

	Parameters := model.Products{
		Public: Public,
		CategoryID: CategoryID,
		Name: Body.Name,
		DescriptionHtml: Body.DescriptionHtml,
		TechnicalSheetUrl: Body.TechnicalSheetUrl,
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

	return ctx.Html(view.Product(findProducts(ctx)))
}

func ProductsUpdate(ctx *controller.Context) error {
	var Body ProductDto

	if err := ctx.Bind(&Body); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}
	var Productie model.Products
	ID, _ := strconv.Atoi(Body.ID)

	match := storage.DB.Last(&Productie, uint(ID))
	if match.Error != nil {
		fmt.Print("No Faq as ", uint(ID))
		return ctx.String(http.StatusNotFound, "")
	}

	Public := true
	if Body.Public == "false" { Public = false }
	Parameters := map[string]interface{}{"Name": Body.Name, "Public": Public}

	file, err := ctx.FormFile("Thumbnail")
	if file != nil && err == nil {
		Upload := uploader.File(file)
		if !Upload.Success {
			fmt.Print(Upload.Message) 
			return ctx.String(http.StatusBadRequest, Upload.Message)
		}
		Parameters["ThumbnailID"] = Upload.ID
	}

	result := storage.DB.Model(Productie).Updates(&Parameters)
						 
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	return ctx.Html(view.Product(findProducts(ctx)))
}

func ProductsStatus(ctx *controller.Context) error {
	var Productie model.Products
	var ID struct{ ID string `param:"id"`}

	if err := ctx.Bind(&ID); err != nil {
		return ctx.String(http.StatusBadRequest, "Parameters Binding Problem: " + err.Error())
	}

	storage.DB.First(&Productie, ID)
	result := storage.DB.Model(Productie).Updates(&map[string]interface{}{
		"Public": !Productie.Public,
	})

	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}


	return ctx.Html(view.Product(findProducts(ctx)))
}

func ProductsRemove(ctx *controller.Context) error {
	var Productie model.Products
	var ID struct{ ID string `param:"id"`}

	if err := ctx.Bind(&ID); err != nil {
		return ctx.String(http.StatusBadRequest, "Parameters Binding Problem: " + err.Error())
	}

	storage.DB.First(&Productie, ID)
	result := storage.DB.Delete(&Productie)
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	return ctx.Html(templ.NopComponent)
}
