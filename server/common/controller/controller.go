// Package controller provides utilities for managing route handlers in an Echo framework application.
// It enhances Echo's capabilities by adding additional features to route handlers and context management.
// This package facilitates the creation of robust web applications by extending Echo's functionality.
//
// Features:
//   - Enhanced Context Management: The Context type wraps Echo's standard context (echo.Context)
//     and adds additional functionalities for handling web requests.
//   - Middleware Initialization: The Initialize function creates middleware for initializing a custom Context instance,
//     allowing extended features to be added to route handlers.
//   - Route Handler Registration: The Register function registers route handlers with additional features
//     and replaces Echo's context with the custom controller.Context type.
//   - HTML Rendering: The Html method renders templ components and returns them as HTML,
//     supporting both full page rendering and fragment rendering for htmx requests.
//   - htmx Request Detection: The IsHtmx method checks if a request is made via htmx by examining the "Hx-Request" header.
package controller

import (
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"main/build/view"
	"main/server/common/globals"
	"main/server/model"
)

// Context wraps the Echo's standard context (echo.Context) to provide additional functionalities.
//
// Functions added by it:
//
// - Initialize() echo.MiddlewareFunc
//
// - Register() echo.HandlerFunc
//
// - Html() error
//
// - IsHtmx() bool
type Context struct { echo.Context }

// Initialize creates a middleware function for initializing a controller Context instance.
// It wraps the standard Echo context with the controller Context type, allowing additional functionalities to be added.
//
// Example usage:
//   app := echo.New()
//   app.Use(controller.Initialize())
//
// Returns:
//   - A middleware function that initializes a controller Context instance.
//
// Notes:
//   - This function is intended to be used as middleware in an Echo application.
//   - It wraps the standard Echo context with the controller Context type, allowing additional functionalities to be added.
//   - The controller Context instance can be accessed within route handlers to utilize the extended features provided.
func Initialize() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx :=  &Context{Context: c}
			return next(ctx)
		}
	}
}

// Registers echo's route handlers with additional features to an Echo's Context instance.
// It takes a route handler function as an argument, enhances it with additional functionality,
// and replaces the Echo context with the custom controller.Context type to provide extended features.
//
// Example usage:
//   app := echo.New()
//   app.GET("/path", controller.Register(func(ctx *controller.Context) error {
//		if ctx.IsHtmx() {
//			return ctx.Html(view.Fragment())
//      }
//
//      return ctx.Html(view.FullPage())
//   }))
// Parameters:
//   - app: An instance of Echo framework where the route handler will be registered.
//   - func(ctx *controller.Context): The route handler function to be registered. It should have the signature func(Context) error.
//
// Notes:
//   - The route handler function should take controller.Context as an argument to utilize the additional features provided by this package.
//   - The controller.Context type extends the standard echo.Context with extra methods and features, like Html() IsHtmx() and others.
func Register(handlerFunc func(*Context) error) echo.HandlerFunc {
    return func(c echo.Context) error { return handlerFunc(c.(*Context)) }
}

// Html renders the given templ component and returns it as HTML.
// If the request is made via htmx, it returns the component as a fragment.
// Otherwise, it embeds the component within the layout of the base HTML.
//
// Example usage:
//   app := echo.New()
//   app.GET("/fullpage", controller.Register(func(ctx *controller.Context) error {
//      return ctx.Html(view.FullPage())
//   }))
//
// Parameters:
//   - component: The templ component to be rendered.
//
// Returns:
//   - An error, if any, encountered during rendering.
//
// Notes:
//   - The method requires access to the request context and response writer via the Context instance.
//   - If the request is made via htmx, the component is rendered as a fragment without any layout.
//   - If the request is not made via htmx, the component is rendered within the layout of the base HTML.
//   - Make sure to handle any errors returned by this method appropriately.
func (ctx *Context) Html(component templ.Component) error {
	if ctx.IsHtmx() { return component.Render(ctx.Request().Context(), ctx.Response()) }

	Interface := ctx.Get("Interface").(model.Interface)
	Base := view.Pages(Interface, component)
	return Base.Render(ctx.Request().Context(), ctx.Response())
}

// IsHtmx checks if the request is made via htmx (Hypertext Markup eXtension).
// It examines the request headers and returns true if the "Hx-Request" header is set to "true".
//
// Example usage:
//   isHtmxRequest := ctx.IsHtmx()
//
// Returns:
//   - true if the request is made via htmx, false otherwise.
//
// Notes:
//   - The method requires access to the request context via the Context instance.
//   - It examines the "Hx-Request" header to determine if the request is an htmx request.
//   - This method can be used to conditionally render content or handle logic based on the type of request.
func (ctx *Context) IsHtmx() bool {
	return ctx.Request().Header.Get("Hx-Request") == "true"
}

type QueryPageParameter struct {
	Page		string		`query:"page"`
	PageSize	string		`query:"pageSize"`
}

func (ctx *Context) Page() int {
	var Query QueryPageParameter

	if ctx.QueryParam("page") == "" { 
		Query.Page = "1" 
	} else {
		ctx.Bind(&Query)
	}

	page, _ := strconv.Atoi(Query.Page)
	if page <= 0 { page = 1 }
	return page
}

func (ctx *Context) PageSize() int {
	var Query QueryPageParameter
	if ctx.QueryParam("pageSize") == "" { Query.PageSize = "-1" }
	
	ctx.Bind(&Query)
	pageSize, _ := strconv.Atoi(Query.PageSize)

	if pageSize <= 0 || pageSize > globals.Env.PageMaxSize {
		pageSize = globals.Env.PageMaxSize
	}
	return pageSize
}