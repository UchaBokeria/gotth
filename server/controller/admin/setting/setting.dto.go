package setting


type ContactDto struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
	Location string `json:"location"`
	LocationLink string `json:"locationLink"`
	LocationIframe string `json:"locationIframe"`
}

type SocialDto struct {
	Facebook  string `json:"Facebook"`
	Instagram string `json:"Instagram"`
	Twitter   string `json:"Twitter"`
	YouTube   string `json:"YouTube"`
}

type SlideshowerDto struct {
	ID    string `param:"id"`
	Name  string `form:"Name"`
	Slogan string `form:"Slogan"`
	Desc   string `form:"Desc"`
	Url   string `form:"Url"`
	Pic   []byte `form:"Pic"`
}

type ReasonerDto struct {
	ID    	string 		`param:"id"`
	Name  	string 		`form:"Name"`
	Title 	string 		`form:"Title"`
	Desc   	string 		`form:"Desc"`
	Url   	string 		`form:"Url"`
	Icon   	[]byte 		`form:"Icon"`
}

type FaqersDto struct {
	ID    	    string 		`param:"id"`
	Name 		string 		`json:"name"`
	Answer 		string 		`json:"answer"`
	Question  	string 		`json:"question"`
}

type NewserDto struct {
	ID         		string       `param:"id"`
	Title			string       `form:"Title"`
	Body 			string       `form:"Body"`
	Public 			string       `form:"Public"`
	Url 			string       `form:"Url"`
	Thumbnail		[]byte	     `form:"Thumbnail"`
	TypeID 			string       `form:"TypeID"`
}
