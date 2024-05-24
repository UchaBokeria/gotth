package abouter

import (
	"fmt"
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
	"net/http"
)

func Abouter(ctx *controller.Context) error {
	var Body AbouterDto

	if err := ctx.Bind(&Body); err != nil {
		return ctx.String(http.StatusBadRequest, "Parameters Binding Problem: " + err.Error())
	}

	var About model.Interface_about

	match := storage.DB.Last(&About, uint(1))
	if match.Error != nil {
		fmt.Print("No About as ", uint(1))
		return ctx.String(http.StatusNotFound, "")
	}

	result := storage.DB.Model(About).Updates(&model.Interface_about{Body: Body.Content})

	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Abouts model.Interface_about
	storage.DB.Last(&Abouts)
	return ctx.Html(view.Abouter(Abouts.Body))
}

func Termer(ctx *controller.Context) error {
	var Body AbouterDto

	if err := ctx.Bind(&Body); err != nil {
		return ctx.String(http.StatusBadRequest, "Parameters Binding Problem: " + err.Error())
	}

	var About model.Interface_about
	var ID uint = uint(1)

	match := storage.DB.Last(&About, ID)
	if match.Error != nil {
		fmt.Print("No Termer as ", ID)
		return ctx.String(http.StatusNotFound, "")
	}

	result := storage.DB.Model(About).Updates(&model.Interface_about{Terms: Body.Content})

	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Abouts model.Interface_about
	storage.DB.Last(&Abouts)
	return ctx.Html(view.Termer(Abouts.Terms))
}
