package setting

import (
	"fmt"
	"main/build/view"
	"main/server/common/controller"
	uploader "main/server/common/helpers"
	"main/server/common/storage"
	"main/server/model"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func index(ctx *controller.Context) error {
	var Interface model.Interface
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
	
	return ctx.Html(view.Setting(Interface, Faq, Newz))
}

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

/* Slideshower */

func Slideshower(ctx *controller.Context) error {
	var Body SlideshowerDto

	if err := ctx.Bind(&Body); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}

	ID, _ := strconv.Atoi(Body.ID)
	var slide model.Interface_slideShow

	match := storage.DB.Last(&slide, uint(ID))
	if match.Error != nil {
		fmt.Print("No slide as ", uint(ID))
		return ctx.String(http.StatusNotFound, "")
	}

	Parameters := model.Interface_slideShow{
		Name: Body.Name,
		Slogan: Body.Slogan,
		Desc: Body.Desc,
		Url: Body.Url,
	}

	file, err := ctx.FormFile("Pic")
	if file != nil && err == nil {	
		Upload := uploader.File(file)
		if !Upload.Success {
			fmt.Print(Upload.Message) 
			return ctx.String(http.StatusBadRequest, Upload.Message)
		}
		Parameters.PicID = &Upload.ID
	}

	result := storage.DB.Model(slide).Updates(&Parameters)
	if result.Error != nil { return ctx.String(http.StatusBadRequest, "") }

	var Slides []model.Interface_slideShow
	storage.DB.Order("interface_slide_shows.index ASC").Preload("Pic").Find(&Slides)

	return ctx.Html(view.Slideshower(Slides))
}

func SlideshowerNew(ctx *controller.Context) error {
	var Body SlideshowerDto

	if err := ctx.Bind(&Body); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}

	var lastSlide model.Interface_slideShow
	storage.DB.Last(&lastSlide)

	Parameters := model.Interface_slideShow{
		InterfaceID: 1,
		Name: Body.Name,
		Slogan: Body.Slogan,
		Desc: Body.Desc,
		Url: Body.Url,
		Index: (lastSlide.Index + 1),
		TypeID: 1,
	}

	file, err := ctx.FormFile("Pic")
	if file != nil && err == nil {	
		Upload := uploader.File(file)
		if !Upload.Success {
			fmt.Print("File did not upload -> " + Upload.Message) 
			// return ctx.String(http.StatusBadRequest, Upload.Message)
		} else {
			Parameters.PicID = &Upload.ID
		}
	}


	result := storage.DB.Create(&Parameters)
	if result.Error != nil { return ctx.String(http.StatusBadRequest, result.Error.Error()) }

	var Slides []model.Interface_slideShow
	storage.DB.Order("interface_slide_shows.index ASC").Preload("Pic").Find(&Slides)

	return ctx.Html(view.Slideshower(Slides))
}

func SlideRemove(ctx *controller.Context) error {
	var Slide model.Interface_slideShow
	var ID struct{ ID string `param:"id"`}

	if err := ctx.Bind(&ID); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}

	storage.DB.First(&Slide, ID)
	storage.DB.Delete(&Slide)

	var Slides []model.Interface_slideShow
	storage.DB.Order("interface_slide_shows.index ASC").Preload("Pic").Find(&Slides)
	return ctx.Html(view.Slideshower(Slides))
}

/* Reasoners */
func Reasoner(ctx *controller.Context) error {
	var Body ReasonerDto

	if err := ctx.Bind(&Body); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}

	ID, _ := strconv.Atoi(Body.ID)
	var slide model.Interface_reasons

	match := storage.DB.Last(&slide, uint(ID))
	if match.Error != nil {
		fmt.Print("No Reason as ", uint(ID))
		return ctx.String(http.StatusNotFound, "")
	}

	Parameters := model.Interface_reasons{
		Name: Body.Name,
		Title: Body.Title,
		Desc: Body.Desc,
		Url: Body.Url,
	}

	file, err := ctx.FormFile("Icon")
	if file != nil && err == nil {
		Upload := uploader.File(file)
		if !Upload.Success {
			fmt.Print(Upload.Message) 
			return ctx.String(http.StatusBadRequest, Upload.Message)
		}
		
		Parameters.IconID = &Upload.ID
	}

	result := storage.DB.Model(slide).Updates(&Parameters)
	if result.Error != nil { return ctx.String(http.StatusBadRequest, "") }

	var Reasons []model.Interface_reasons
	storage.DB.Order("interface_reasons.created_at desc").Preload("Icon").Find(&Reasons)

	return ctx.Html(view.Reasoners(Reasons))
}

func ReasonerNew(ctx *controller.Context) error {
	var Body ReasonerDto

	if err := ctx.Bind(&Body); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}

	var Parameters = model.Interface_reasons{
			InterfaceID: 1,
			Name: Body.Name,
			// Slug: Body.Name,
			Title: Body.Title,
			Desc: Body.Desc,
			Url: Body.Url,
		}

	file, err := ctx.FormFile("Icon")
	if file != nil && err == nil {	
		Upload := uploader.File(file)
		if !Upload.Success {
			fmt.Print("File did not upload -> " + Upload.Message) 
			// return ctx.String(http.StatusBadRequest, Upload.Message)
		} else {
			Parameters.IconID = &Upload.ID
		}
	}

	result := storage.DB.Create(&Parameters)
						 
	if result.Error != nil { return ctx.String(http.StatusBadRequest, result.Error.Error()) }

	var Reasons []model.Interface_reasons
	storage.DB.Order("interface_reasons.created_at desc").Preload("Icon").Find(&Reasons)
	return ctx.Html(view.Reasoners(Reasons))
}

func ReasonerRemove(ctx *controller.Context) error {
	var Reason model.Interface_reasons
	var ID struct{ ID string `param:"id"`}

	if err := ctx.Bind(&ID); err != nil {
		fmt.Print("Parameters Binding Problem: ", err)
		return err
	}

	storage.DB.First(&Reason, ID)
	storage.DB.Delete(&Reason)

	var Reasons []model.Interface_reasons
	storage.DB.Order("interface_reasons.created_at desc").Preload("Icon").Find(&Reasons)
	return ctx.Html(view.Reasoners(Reasons))
}

/* Faqers */

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

/* Faqers */

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

/* About */

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
