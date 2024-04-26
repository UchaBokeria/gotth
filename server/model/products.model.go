package model

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name 				string
	Slug 				string
	Description 		string
	TechnicalSheetID	int
	TechnicalSheet 		Files
	ThumbnailID 		int
	Thumbnail 			Files
	Packing         	Product_packaging
	Approvals			Product_approvals
	Specifications  	Product_specifications
}

type Product_packaging struct {
	gorm.Model
	ProductsID      	uint
	Name				string
	Slug 				string
	Reference			string
	Conditioning		string
	Cardboard 			string
}

type Product_approvals struct {
	gorm.Model
	ProductsID      	uint
	Name				string
	Slug 				string
}

type Product_specifications struct {
	gorm.Model
	ProductsID      	uint
	Name				string
	Slug 				string
}

