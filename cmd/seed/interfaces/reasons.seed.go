package interfaces

import (
	"main/server/common/storage"
	"main/server/model"
)

var Reasons = []model.Interface_reasons{
	{
		InterfaceID: 	1,
		Name:   		"Garantie",
		Slug:   		"garantie",
		Title:  		"გარანტია",
		Desc:   		"იხილეთ Yacco - ს გარანტია ზეთებზე და დარწმუნდი ხარისხში. დამატებითი შეკითხვების ინფორმაციისთვის დაგვიკავშირდით",
		Url:    		"https://yaccogaranties.com/accueil_yacco.php",
		IconID: 		GetPointedInt(7),
	},
	{
		InterfaceID: 	1,
		Name:   		"Branches",
		Slug:   		"branches",
		Title:  		"სად ვიშოვოთ",
		Desc:   		"იხილეთ სად შეგიძლიათ იშოვოთ Yacco პროდუქცია თქვენს ახლოს. დამატებითი შეკითხვების ინფორმაციისთვის დაგვიკავშირდით",
		Url:    		"/branches/",
		IconID: 		GetPointedInt(8),
	},
	{
		InterfaceID: 	1,
		Name:   		"Garantie",
		Slug:   		"garantie",
		Title:  		"გაიგე მეტი",
		Desc:   		"გაიგე მეტი კომპანიის ისტორიაზე და ჩვენს მიღწევებზე. დამატებითი შეკითხვების ინფორმაციისთვის დაგვიკავშირდით",
		Url:    		"/about/",
		IconID: 		GetPointedInt(9),
	},
	{
		InterfaceID: 	1,
		Name:   		"Faq",
		Slug:   		"fa",
		Title:  		"ხშირად დასმული შეკითხვები",
		Desc:   		"იხილეთ მეტი ინფორმაცია და ხშირად დასმული შეკითხვები. დამატებითი შეკითხვების ინფორმაციისთვის დაგვიკავშირდით",
		Url:    		"/faq/",
		IconID: 		GetPointedInt(10),
	},
	{
		InterfaceID: 	1,
		Name:   		"News",
		Slug:   		"news",
		Title:  		"სიახლეები",
		Desc:   		"იხილეთ სიახლეები ვიდეო ფოტო და ბლოგ სახით. დამატებითი შეკითხვების ინფორმაციისთვის დაგვიკავშირდით",
		Url:    		"/news/",
		IconID: 		GetPointedInt(11),
	},
	{
		InterfaceID: 	1,
		Name:   		"terms",
		Slug:   		"terms",
		Title:  		"წესები და პირობები",
		Desc:   		"იხილეთ წესები და პირობები, დამატებითი შეკითხვების ინფორმაციისთვის დაგვიკავშირდით",
		Url:    		"/terms",
		IconID: 		GetPointedInt(12),
	},
}

func Reason() {
	for _, row := range Reasons { storage.DB.Create(&row) }
}