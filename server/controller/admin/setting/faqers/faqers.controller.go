package faqers

import (
	"fmt"
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
	"net/http"
	"strconv"
)

func Faqers(ctx *controller.Context) error {
	var Body FaqersDto

	if err := ctx.Bind(&Body); err != nil {
		return ctx.String(http.StatusBadRequest, "Parameters Binding Problem: " + err.Error())
	}

	var faq model.Faq
	ID, _ := strconv.Atoi(Body.ID)

	match := storage.DB.Last(&faq, uint(ID))
	if match.Error != nil {
		fmt.Print("No Faq as ", uint(ID))
		return ctx.String(http.StatusNotFound, "")
	}

	result := storage.DB.Model(faq).Updates(&model.Faq{
		Name: Body.Name,
		Answer: Body.Answer,
		Question: Body.Question,
	})

	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Faqs []model.Faq
	storage.DB.Order("created_at desc").Find(&Faqs)
	return ctx.Html(view.Faqers(Faqs))
}

func FaqersNew(ctx *controller.Context) error {
	var Body FaqersDto

	if err := ctx.Bind(&Body); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}

	result := storage.DB.Create(&model.Faq{
		Name: Body.Name,
		Answer: Body.Answer,
		Question: Body.Question,
	})
						 
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Faqs []model.Faq
	storage.DB.Order("created_at desc").Find(&Faqs)
	return ctx.Html(view.Faqers(Faqs))
}

func FaqersRemove(ctx *controller.Context) error {
	var Faq model.Faq
	var ID struct{ ID string `param:"id"`}

	if err := ctx.Bind(&ID); err != nil {
		return ctx.String(http.StatusBadRequest, "Parameters Binding Problem: " + err.Error())
	}

	storage.DB.First(&Faq, ID)
	result := storage.DB.Delete(&Faq)
	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, result.Error.Error())
	}

	var Faqs []model.Faq
	storage.DB.Order("created_at desc").Find(&Faqs)
	return ctx.Html(view.Faqers(Faqs))
}
