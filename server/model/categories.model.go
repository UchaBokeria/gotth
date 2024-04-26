package model

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	Name 				*string
	Filters				[]Category_filters `gorm:"many2many:Category_filters_joins;constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
}

type Category_filters struct {
	gorm.Model
	Name 				string
	Slug 				string
	Options 			[]Category_filters_option `gorm:"many2many:Category_filters_options_join;constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
	Default_value 		string
}

type Category_filters_option struct {
    gorm.Model
    Name 				string
	Key                 string
    Value 				string
}