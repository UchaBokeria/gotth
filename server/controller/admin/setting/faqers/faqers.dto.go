package faqers

type FaqersDto struct {
	ID    	    string 		`param:"id"`
	Name 		string 		`json:"name"`
	Answer 		string 		`json:"answer"`
	Question  	string 		`json:"question"`
}
