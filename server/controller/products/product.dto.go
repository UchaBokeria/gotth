package products

type FiltersQuery struct {
	Category 	 	string 		`query:"category"`
	Searcher        string      `query:"searcher"`
}

type ProductDetail struct {
	ID              string      `param:"id"`
}