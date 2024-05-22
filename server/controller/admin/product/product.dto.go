package product

type ProductDto struct {
	ID   				string 		`param:"id" form:"id"`
	Public 				string 		`form:"Public"`
	Name 				string 		`form:"Name"`
	Thumbnail 			[]byte 		`form:"Thumbnail"`
	DescriptionHtml		string		`form:"DescriptionHtml"`
	CategoryID			string		`form:"CategoryID"`
	TechnicalSheetUrl	string		`form:"TechnicalSheetUrl"`
}