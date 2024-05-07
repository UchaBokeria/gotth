package users

import (
	"main/server/common/storage"
	"main/server/model"
)

var Seed []model.Users = []model.Users{
	{ Email: "admin@yacco.ge", Fullname: "Administrator", Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" },
}

func Populate() {
	for _, row := range Seed { storage.DB.Create(&row) }
}