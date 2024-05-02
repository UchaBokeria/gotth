package landing

import (
	"gorm.io/gorm"

	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
)

func index(ctx *controller.Context) error {
	var Interface model.Interface

	storage.DB.
	Preload("Contact").
	Preload("News.Thumbnail").
	Preload("Reasons.Icon").
		Preload("SlideShow", func(db *gorm.DB) *gorm.DB {
			return db.Order("interface_slide_shows.index ASC").Preload("Pic")
		}).
		Last(&Interface)

	return ctx.Html(view.Landing(Interface))
}

func subscribe(ctx *controller.Context) error {
	return ctx.Html(view.Terms(""))
}