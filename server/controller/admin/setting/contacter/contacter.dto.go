package contacter

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
