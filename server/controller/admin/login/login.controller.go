package login

import (
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"

	"golang.org/x/crypto/bcrypt"
)

func index(ctx *controller.Context) error {
	return ctx.Renders(view.Login())
}

type LoginCredsParams struct {
	Email string `param:"email"`
	Password string `param:"password"`
}

func login(ctx *controller.Context) error {
	var UserMatch model.Users
	var Parameters LoginCredsParams

	ctx.Bind(Parameters)
	storage.DB.Last(&UserMatch, &model.Users{ Email: Parameters.Email })


	Err := bcrypt.CompareHashAndPassword([]byte(UserMatch.Password), []byte(Parameters.Password))
    if Err != nil { return ctx.Renders(view.Login()) }

	return ctx.JSON(200, UserMatch)
}
