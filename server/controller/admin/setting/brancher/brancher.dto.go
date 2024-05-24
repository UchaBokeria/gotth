package brancher

type BrancherDto struct {
	ID    	     string 		`param:"id"`
	Map 		 string 		`form:"Map"`
	Name 		 string 		`form:"Name"`
	DistrictID	 string 		`form:"DistrictID"`
	PhoneNumber  string 		`form:"PhoneNumber"`
}
