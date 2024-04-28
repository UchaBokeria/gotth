package model

import (
	"time"

	"gorm.io/gorm"
)

type Branches struct {
	gorm.Model
	Name       	string
	PhoneNumber string
	Region 		string
	OpensAt     time.Time 	`gorm:"type:time"`
	ClosesAt    time.Time 	`gorm:"type:time"`
	CityID 		int
	City 		Cities 		`gorm:"constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
}

type Cities struct {
	gorm.Model
	Name  		string
	Slug 		string
	State 		string
}