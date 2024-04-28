package interfaces

import (
	"main/server/common/storage"
	"main/server/model"
)

var SocialMedias = []model.Social_media{
	{
		InterfaceID:    1,
		Name:   		"Facebook",
		Slug:   		"facebook",
		Url:    		"https://www.facebook.com/yacco",
		IconID: 		4,
	},
	{
		InterfaceID:    1,
		Name:   		"Instagram",
		Slug:   		"instagram",
		Url:    		"https://www.instagram.com/yacco",
		IconID: 		5,
	},
	{
		InterfaceID:    1,
		Name:   		"Twitter",
		Slug:   		"twitter",
		Url:    		"https://www.twitter.com/yacco",
		IconID: 		6,
	},
	{
		InterfaceID:    1,
		Name:   		"YouTube",
		Slug:   		"youtube",
		Url:    		"https://www.youtube.com/example",
		IconID: 		2,
	},
}

func SocialMedia() {
	for _, row := range SocialMedias { storage.DB.Create(&row) }
}