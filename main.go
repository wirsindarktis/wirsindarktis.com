package main

import (
	"os"

	"github.com/aerogo/aero"
	"github.com/wirsindarktis/wirsindarktis.com/pages"
)

func main() {
	app := aero.New()
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {
	// Certificate
	if os.Getenv("PRODUCTION") == "true" {
		app.Security.Load("security/server.crt", "security/server.key")
	}

	// Pages
	pages.Install(app)

	return app
}
