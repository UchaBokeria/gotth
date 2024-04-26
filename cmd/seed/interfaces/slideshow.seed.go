package interfaces

import (
	"main/server/common/storage"
	"main/server/model"
)

func Slideshow() {
	for _, row := range []model.Interface_slideShow{
		{
			InterfaceID: 	1,
			Name:   		"Slideshow 1",
			Slug:   		"slideshow-1",
			Slogan: 		"შეიძინე YACCO ბრედნდის",
			Desc:   		"უმაღლესი ხარისხის ზეთი საქართველოში",
			Url:    		"/products?categories=1",
			TypeID: 		1,
			PicID:  		2,
			Index:  		"2",
		},
		{
			InterfaceID: 	1,
			Name:   		"Slideshow 1",
			Slug:   		"slideshow-1",
			Slogan: 		"Slogan for Slideshow 1",
			Desc:   		"Description for Slideshow 1",
			Url:    		"/categories/",
			TypeID: 		1,
			PicID:  		1,
			Index:  		"1",
		},
		{
			InterfaceID: 	1,
			Name:   		"Slideshow 1",
			Slug:   		"slideshow-1",
			Slogan: 		"Slogan for Slideshow 1",
			Desc:   		"Description for Slideshow 1",
			Url:    		"/about",
			TypeID: 		1,
			PicID:  		3,
			Index:  		"3",
		},
	} { storage.DB.Create(&row) }
}