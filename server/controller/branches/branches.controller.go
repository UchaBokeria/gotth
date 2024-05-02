package branches

import (
	"log"

	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
)

func index(ctx *controller.Context) error {
	var Branches []model.Branches
	storage.DB.Preload("City").Preload("Shifts").Find(&Branches)

	for _, v := range Branches {
		log.Printf(" a %T \n", v)
	}
	return ctx.Html(view.Branches(Branches))
}