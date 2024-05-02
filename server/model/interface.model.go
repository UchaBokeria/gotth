package model

import "gorm.io/gorm"

type Interface struct {
	gorm.Model
	Ver  		int
	Name 		string
	Slug 		string
	News        []News 					`gorm:"many2many:Interface_news_joins;constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
	SlideShow	[]Interface_slideShow
	Reasons		[]Interface_reasons
	Contact		Interface_contact
	About		Interface_about
	SocialMedia []Social_media
}

type Interface_slideShow struct {
	gorm.Model
	InterfaceID uint
	Name 		string
	Slug 		string
	Slogan      string
	Desc        string
	Url         string
	Index       int
	TypeID      int
	Type		File_types
	PicID 		int
	Pic 		Files
}

type Interface_reasons struct {
	gorm.Model
	InterfaceID uint
	Name 		string
	Slug 		string
	Title		string
	Desc		string
	Url			string
	IconID 		int
	Icon 		Files
}

type Interface_contact struct {
	gorm.Model
	InterfaceID uint
	Ver         int
	Name 		string
	Slug 		string
	Phone       string
	Email       string
	Location    string
	ShortDesc   string
}

type Social_media struct {
	gorm.Model
	InterfaceID uint
	Name 		string
	Slug 		string
	Url         string
	IconID 		int
	Icon 		Files
}

type Interface_about struct {
	gorm.Model
	InterfaceID uint
	Ver    		int
	Body   		string
	Terms  		string
}
