package model

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Views       	int
	Title 			string
	Body  			string
	Public 			bool
	PublishedAt 	time.Time
	Url 			string
	ThumbnailID 	int
	Thumbnail 		Files			`gorm:"constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
	TypeID      	int
	Type        	News_types 		`gorm:"constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
}

type News_types struct {
	gorm.Model
	Name 			string
	Slug 			string
}