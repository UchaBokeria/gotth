package branches

import (
	"time"

	"main/server/common/storage"
	"main/server/model"
)

var Seed = []model.Branches{
	{
		Name:       	"ავტოშოპი",
		PhoneNumber: 	"+995 598 731 127",
		Region:         "წერეთელი",
		OpensAt:     	time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		ClosesAt:    	time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		CityID: 		1,
	},
	{
		Name:       	"ზეთების მაღაზია",
		PhoneNumber: 	"+995 570 456 221",
		Region:         "წერეთელი",
		OpensAt:     	 time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		ClosesAt:    	 time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		CityID: 		2,
	},
	{
		Name:       	"ავტოშოპი",
		PhoneNumber: 	"+995 594 832 126",
		Region:         "ჯიქია",
		OpensAt:     	 time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		ClosesAt:    	 time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		CityID: 		1,
	},
	{
		Name:       	"ავტოშოპი",
		PhoneNumber: 	"+995 592 224 552",
		Region:         "გლდანი",
		OpensAt:     	 time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		ClosesAt:    	 time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		CityID: 		1,
	},
	{
		Name:       	"ნაწილები და ზეთები",
		PhoneNumber: 	"+995 598 111 422",
		Region:         "ლილო",
		OpensAt:     	 time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		ClosesAt:    	 time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		CityID: 		3,
	},
	{
		Name:       	"ავტომანია",
		PhoneNumber: 	"+995 558 112 332",
		Region:         "ისანი",
		OpensAt:     	 time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		ClosesAt:    	 time.Date(0, 1, 1, 9, 45, 0, 0, time.UTC),
		CityID: 		1,
	},
}

func Populate() {
	for _, row := range Seed { storage.DB.Create(&row) }
}