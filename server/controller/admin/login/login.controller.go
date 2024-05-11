package login

import (
	"encoding/json"
	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func index(ctx *controller.Context) error {
	return ctx.Renders(http.StatusOK, view.Login())
}

type CredsParams struct {
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func changePassword(ctx *controller.Context) error {
	var Parameters CredsParams
	if err := (&echo.DefaultBinder{}).BindBody(ctx, &Parameters);

	err != nil || Parameters.Email == "" || Parameters.Password == "" {
		return ctx.String(http.StatusBadRequest, "atrakeb")
	}

	Hash, err := bcrypt.GenerateFromPassword([]byte(Parameters.Password), 12)
	if err != nil { return ctx.String(http.StatusBadRequest, "bycrypt!!") }

	Result := storage.DB.Table("users").Where(&model.Users{ Email: Parameters.Email }).Update("password", string(Hash))
	if Result.Error != nil || Result.RowsAffected < 1 { return ctx.String(http.StatusBadRequest, "No rows affected") }
	return ctx.String(http.StatusOK, "success")
}

func login(ctx *controller.Context) error {
	var UserMatch model.Users
	var Parameters CredsParams

	if err := (&echo.DefaultBinder{}).BindBody(ctx, &Parameters);
	err != nil || Parameters.Email == "" || Parameters.Password == "" {
		return ctx.Renders(http.StatusBadRequest, view.Login())
	}
	
	Result := storage.DB.Last(&UserMatch, &model.Users{ Email: Parameters.Email })
	if Result.Error != nil || Result.RowsAffected < 1 {
		return ctx.Renders(http.StatusNotFound, view.Login())
	}

	err := bcrypt.CompareHashAndPassword([]byte(UserMatch.Password), []byte(Parameters.Password))
    if err != nil { return ctx.Renders(http.StatusUnauthorized, view.Login()) }

	Expires := time.Now().Add(24 * time.Hour)
	AuthHtmxMeta, _ := json.Marshal(struct {
		Expires   time.Time    `json:"Expires"`
        Email     string    `json:"Email"`
        Token     string    `json:"Token"`
	}{
		Expires: Expires, Email: UserMatch.Email, Token: UserMatch.Token,
	})      

	ctx.WriteCookie(controller.Cookie{ Key: "user", Value: UserMatch.Email, Expires: Expires })
	ctx.WriteCookie(controller.Cookie{ Key: "token", Value: UserMatch.Token, Expires: Expires })
	ctx.Response().Header().Add("HX-AUTH-META", string(AuthHtmxMeta))

	return ctx.Html(view.Admin(templ.NopComponent))
}
