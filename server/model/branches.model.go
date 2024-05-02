package model

import (
	"time"

	"gorm.io/gorm"
)

type Branches struct {
	gorm.Model
	Name       	string
	Slug       	string
	PhoneNumber string
	Region 		string
	Map         string
	Shifts      []Branch_shifts
	CityID 		int
	City 		Cities 			`gorm:"constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
}

type Branch_shifts struct {
	gorm.Model
	BranchesID    uint
	Name       	string
	Slug       	string
	OpensAt     time.Time 		`gorm:"type:time"`
	ClosesAt    time.Time 		`gorm:"type:time"`
}

type Cities struct {
	gorm.Model
	Name  		string
	Slug 		string
	State 		string
}