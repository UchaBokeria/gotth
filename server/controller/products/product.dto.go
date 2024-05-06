package products

type FiltersQuery struct {
	Category 	 	string 		`query:"category"`
}

type ProductDetail struct {
	ID              string      `param:"id"`
}