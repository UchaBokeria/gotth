package model

import (
	"time"

	"gorm.io/gorm"
)

type Branches struct {
	gorm.Model
	Name       	*string
	PhoneNumber *string
	OpensAt     time.Time
	ClosesAt    time.Time
	CityID 		int
	City 		Cities `gorm:"constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
}

type Cities struct {
	gorm.Model
	Name  		*string
	State 		string
}