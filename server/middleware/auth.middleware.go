package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
)

type AuthHeaderCreds struct {
	Email string `header:"X-User"`
	Token string `header:"X-Token"`
}

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return controller.Register(func(ctx *controller.Context) error {
			var User model.Users
			var Parameters AuthHeaderCreds

			ctx.Set("ISADMIN", false)

			err := (&echo.DefaultBinder{}).BindHeaders(ctx, &Parameters);
			if err != nil || Parameters.Email == "" || Parameters.Token == ""{
				Parameters.Email = ctx.ReadCookie("user").Value
				Parameters.Token = ctx.ReadCookie("token").Value
			}

			if Parameters.Email == "" || Parameters.Token == "" {
				return ctx.Renders(http.StatusBadRequest, view.Login())
			}

			result := storage.DB.Where(&model.Users{Email: Parameters.Email, Token: Parameters.Token}).Last(&User)
			if result.Error != nil || result.RowsAffected < 1 { 
				return ctx.Renders(http.StatusUnauthorized, view.Login())
			}

			ctx.Set("ISADMIN", true)
			ctx.Set("USER", User)

			return next(ctx)
		})
	}
}

