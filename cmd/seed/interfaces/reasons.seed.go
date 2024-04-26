package interfaces

import (
	"main/server/common/storage"
	"main/server/model"
)

func Reasons() {
	for _, row := range []model.Interface_reasons{
		{
			InterfaceID: 	1,
			Name:   		"Garantie",
			Slug:   		"garantie",
			Title:  		"გარანტია",
			Desc:   		"იხილეთ ოფიციალური Yacco - ს გარანტია ზეთებზე.",
			Url:    		"https://yaccogaranties.com/accueil_yacco.php",
			IconID: 		7,
		},
		{
			InterfaceID: 	1,
			Name:   		"Branches",
			Slug:   		"branches",
			Title:  		"სად ვიშოვოთ",
			Desc:   		"იხილეთ სად შეგიძლიათ იშოვოთ Yacco პროდუქცია თქვენს ახლოს",
			Url:    		"/branches/",
			IconID: 		8,
		},
		{
			InterfaceID: 	1,
			Name:   		"Garantie",
			Slug:   		"garantie",
			Title:  		"მაღალი ხარისხი",
			Desc:   		"ეს თანამშრომლობა გარდამტეხი იყო როგორც კომპანიისთვის, ასევე ქართული ბაზრისათვისაც. თეგეტა იყო ერთ-ერთი პირველი",
			Url:    		"/about",
			IconID: 		9,
		},
		{
			InterfaceID: 	1,
			Name:   		"Garantie",
			Slug:   		"garantie",
			Title:  		"გარანტია",
			Desc:   		"იხილეთ ოფიციალური Yacco - ს გარანტია ზეთებზე.",
			Url:    		"https://yaccogaranties.com/accueil_yacco.php",
			IconID: 		10,
		},
		{
			InterfaceID: 	1,
			Name:   		"Branches",
			Slug:   		"branches",
			Title:  		"სად ვიშოვოთ",
			Desc:   		"იხილეთ სად შეგიძლიათ იშოვოთ Yacco პროდუქცია თქვენს ახლოს",
			Url:    		"/branches/",
			IconID: 		11,
		},
		{
			InterfaceID: 	1,
			Name:   		"Garantie",
			Slug:   		"garantie",
			Title:  		"მაღალი ხარისხი",
			Desc:   		"ეს თანამშრომლობა გარდამტეხი იყო როგორც კომპანიისთვის, ასევე ქართული ბაზრისათვისაც. თეგეტა იყო ერთ-ერთი პირველი",
			Url:    		"/about",
			IconID: 		12,
		},
	} { storage.DB.Create(&row) }
}