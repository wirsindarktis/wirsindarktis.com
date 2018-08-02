package home

import (
	"github.com/aerogo/aero"
	"github.com/wirsindarktis/wirsindarktis.com/components"
)

// Get the homepage.
func Get(ctx *aero.Context) string {
	return ctx.HTML(components.Home())
}
