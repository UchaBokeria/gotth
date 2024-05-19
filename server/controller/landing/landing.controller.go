package landing

import (
	"log"

	"gorm.io/gorm"

	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
	mailer "main/server/service/mail"
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
	Result, _ := mailer.Send(mailer.Config{
		To: "ucha2bokeria@gmail.com",
		Subject: "New Subsribe",
		Body: "მადლობა გამოწერისთვის, იხილეთ სიახლეები ჩვენს ვებ გვერდზე და მიიღეთ ექსკლუზიური სიახლეები ელ ფოსტის საშუალებით",
	})

	if Result {
		log.Print("Sent")
	}
	return ctx.Html(view.Subscribe())
}