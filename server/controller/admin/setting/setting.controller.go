package setting

import (
	"fmt"
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
	"net/http"

	"gorm.io/gorm"
)

func index(ctx *controller.Context) error {
	var Interface model.Interface
	var Current struct{ Tab string `param:"tab"`}

	if err := ctx.Bind(&Current); err != nil {
		return ctx.String(http.StatusBadRequest, "Parameters Binding Problem: " + err.Error())
	}

	if Current.Tab == "" { Current.Tab = "Contacter" }

	result := storage.DB.
		Preload("SlideShow", func(db *gorm.DB) *gorm.DB {
			return db.Order("interface_slide_shows.index ASC").Preload("Pic")
		}).
		Preload("News", func(db *gorm.DB) *gorm.DB {
			return db.Order("news.created_at desc").Preload("Thumbnail")
		}).
		Preload("Reasons", func(db *gorm.DB) *gorm.DB {
			return db.Order("interface_reasons.created_at desc").Preload("Icon")
		}).
		Preload("About").
		Preload("Contact").
		Preload("SocialMedia.Icon").
		Last(&Interface)
	
	if result.Error != nil {
		fmt.Print(result.Error.Error())
		return ctx.Html(view.ErrorPage())
	}

	var Faq []model.Faq
	storage.DB.Order("created_at desc").Find(&Faq)

	var Newz []model.News
	storage.DB.Order("created_at desc").Preload("Thumbnail").Find(&Newz)

	var Cities []model.Cities
	storage.DB.Find(&Cities)

	var Branches []model.Branches
	storage.DB.Order("created_at desc").Preload("District.City").Find(&Branches)

	return ctx.Html(view.Setting(Interface, Faq, Newz, Branches, Cities, Current.Tab))
}
