package reasoner

type ReasonerDto struct {
	ID    	string 		`param:"id"`
	Name  	string 		`form:"Name"`
	Title 	string 		`form:"Title"`
	Desc   	string 		`form:"Desc"`
	Url   	string 		`form:"Url"`
	Icon   	[]byte 		`form:"Icon"`
}
