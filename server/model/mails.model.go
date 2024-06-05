package model

import "gorm.io/gorm"

type Mails struct {
	gorm.Model
	Subject     string
	Body     	string
	From   		string
	To 			string
	Cc 			*string
	Bcc 		*string
}

type Subscribes struct {
	gorm.Model
	Email string
}
