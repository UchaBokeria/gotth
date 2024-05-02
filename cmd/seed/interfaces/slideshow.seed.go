package interfaces

import (
	"main/server/common/storage"
	"main/server/model"
)

var Slideshows = []model.Interface_slideShow{
	{
		InterfaceID: 	1,
		Name:   		"Slideshow 2",
		Slug:   		"slideshow-2",
		Slogan: 		"რალის ჩემპიონატი 2022 წელი",
		Desc:   		"YACCO ბრენდმა უმაღლეი შედეგი დადო ! 1930-იანი წლების დასაწყისში Yacco-მ განიცადა მთელი რიგი იმედგაცრუებები თავის მთავარ მომხმარებლებთან, სერიოზული ფინანსური სირთულეების ფონზე: შეწყდა ხელშეკრულებები ამილკართან, დონეტთან და ვოისინთან.",
		Url:    		"/categories/",
		TypeID: 		1,
		PicID:  		2,
		Index:  		2,
	},
	{
		InterfaceID: 	1,
		Name:   		"Slideshow 1",
		Slug:   		"slideshow-1",
		Slogan: 		"შეიძინე YACCO ბრედნდის",
		Desc:   		"უმაღლესი ხარისხის ზეთი საქართველოში, კომპანია მთელი ძალისხმევის კონცენტრირებას ახდენს საპოხი მასალების მიწოდებაზე, ხოლო დამხმარე საქმიანობა ძალიან სწრაფად მცირდება",
		Url:    		"/products?categories=1",
		TypeID: 		1,
		PicID:  		1,
		Index:  		1,
	},
	{
		InterfaceID: 	1,
		Name:   		"Slideshow 1",
		Slug:   		"slideshow-1",
		Slogan: 		"1919 ფონდი \"მსოფლიო რეკორდების ნავთობი\"",
		Desc:   		"ლეგენდარული ზეთის ბრენდის სახელის პირველი ნაბიჯები, Hispano-ს დირექტორებს სურთ შექმნან შვილობილი კომპანია, რომელიც განკუთვნილია ჩარხების წარმოებისთვის და ხელსაწყოების და საპოხი ზეთების განაწილებისთვის",
		Url:    		"/about",
		TypeID: 		1,
		PicID:  		3,
		Index:  		3,
	},
}

func Slideshow() {
	for _, row := range Slideshows { storage.DB.Create(&row) }
}