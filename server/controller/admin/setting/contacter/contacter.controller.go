package contacter

import (
	"fmt"
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
	"net/http"
)

func Contact(ctx *controller.Context) error {
	var Body ContactDto
	var Contact model.Interface_contact

	storage.DB.Last(&Contact)

	if err := ctx.Bind(&Body); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	} else if Body.Email == "" || Body.Phone == "" {
		return fmt.Errorf("parameters are not provided :: Email: `" + Body.Email +"`, Fullname: `"+ Body.Phone + "`")
	}

	Contact.Email = Body.Email
	Contact.Phone = Body.Phone
	Contact.Location = Body.Location
	Contact.LocationLink = Body.LocationLink
	Contact.LocationIframe = Body.LocationIframe
	result := storage.DB.Save(&Contact)

	if result.Error != nil {
		return ctx.String(http.StatusBadRequest, "N")
	}

	return ctx.Html(view.Contacter(Contact))
}

func Socials(ctx *controller.Context) error {
	var Body SocialDto

	if err := ctx.Bind(&Body); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}

	var SocialMediaFacebook model.Social_media
	resultFacebook := storage.DB.Model(SocialMediaFacebook).
						 Where(&model.Social_media{Name: "Facebook"}).
						 Updates(&model.Social_media{Url: Body.Facebook})
						 
	if resultFacebook.Error != nil { return ctx.String(http.StatusBadRequest, "Facebook") }

	var SocialMediaInstagram model.Social_media
	resultInstagram := storage.DB.Model(SocialMediaInstagram).
						 Where(&model.Social_media{Name: "Instagram"}).
						 Updates(&model.Social_media{Url: Body.Instagram})
						 
	if resultInstagram.Error != nil { return ctx.String(http.StatusBadRequest, "Instagram") }

	var SocialMediaTwitter model.Social_media
	resultTwitter := storage.DB.Model(SocialMediaTwitter).
						 Where(&model.Social_media{Name: "Twitter"}).
						 Updates(&model.Social_media{Url: Body.Twitter})
						 
	if resultTwitter.Error != nil { return ctx.String(http.StatusBadRequest, "Twitter") }

	var SocialMediaYouTube model.Social_media
	resultYouTube := storage.DB.Model(SocialMediaYouTube).
						 Where(&model.Social_media{Name: "YouTube"}).
						 Updates(&model.Social_media{Url: Body.YouTube})

	if resultYouTube.Error != nil { return ctx.String(http.StatusBadRequest, "YouTube") }

	var SocialMedia []model.Social_media
	storage.DB.Find(&SocialMedia)

	return ctx.Html(view.SocialMediaer(SocialMedia))
}
