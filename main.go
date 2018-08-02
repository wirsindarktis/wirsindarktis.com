package main

import (
	"github.com/aerogo/aero"
	"github.com/wirsindarktis/wirsindarktis.com/pages"
)

func main() {
	app := aero.New()
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {
	pages.Install(app)
	return app
}
