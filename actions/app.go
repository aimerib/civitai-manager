package actions

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"civitai/models"
	"civitai/public"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/v3/pop/popmw"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/middleware/csrf"
	"github.com/gobuffalo/middleware/forcessl"
	"github.com/gobuffalo/middleware/i18n"
	"github.com/gobuffalo/middleware/paramlogger"
	"github.com/unrolled/secure"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")

var (
	app     *buffalo.App
	appOnce sync.Once
	T       *i18n.Translator
)

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	appOnce.Do(func() {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_civitai_session",
		})

		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//   c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))
		// Setup and use translations:
		// app.Use(translations())

		app.GET("/routes", HomeHandler)

		app.GET("/", ModelsIndexHandler)
		app.GET("/models/{id}", ModelsShowHandler)
		app.GET("/settings", SettingsHandler)
		app.POST("/settings/run-fetch-job", RunFetchJobHandler)
		app.Use(SetLayout)
		app.Use(CacheMiddleware)
		app.ServeFiles("/", http.FS(public.FS())) // serve files from the public directory
	})

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
// func translations() buffalo.MiddlewareFunc {
// 	var err error
// 	if T, err = i18n.New(locales.FS(), "en-US"); err != nil {
// 		app.Stop(err)
// 	}
// 	return T.Middleware()
// }

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}

func SetLayout(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		layout := "application.html"
		if c.Request().Header.Get("HX-Request") == "true" && strings.Trim(c.Request().RequestURI, "/") != "" {
			layout = "content.html"
		}
		r.HTMLLayout = layout
		return next(c)
	}
}

func CacheMiddleware(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if strings.HasPrefix(c.Request().URL.Path, "/images/") {
			fmt.Println("Cache images")
			c.Response().Header().Set("Cache-Control", "public, max-age=31536000")
		}
		return next(c)
	}
}
