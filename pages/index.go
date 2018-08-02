package pages

import (
	"path"

	"github.com/aerogo/aero"
	"github.com/aerogo/layout"
	"github.com/wirsindarktis/wirsindarktis.com/components/css"
	"github.com/wirsindarktis/wirsindarktis.com/components/js"
	"github.com/wirsindarktis/wirsindarktis.com/layout"
	"github.com/wirsindarktis/wirsindarktis.com/pages/home"
)

// Install configures the application routes.
func Install(app *aero.Application) {
	l := layout.New(app)
	l.Render = fullpage.Render

	// Pages
	l.Page("/", home.Get)

	// Script bundle
	scriptBundle := js.Bundle()

	app.Get("/scripts", func(ctx *aero.Context) string {
		return ctx.JavaScript(scriptBundle)
	})

	// CSS bundle
	cssBundle := css.Bundle()

	app.Get("/styles", func(ctx *aero.Context) string {
		return ctx.CSS(cssBundle)
	})

	// Static files
	app.Get("/images/*file", func(ctx *aero.Context) string {
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return ctx.File(path.Join("images", ctx.Get("file")))
	})

	// Manifest
	app.Get("/manifest.json", func(ctx *aero.Context) string {
		return ctx.JSON(app.Config.Manifest)
	})
}
