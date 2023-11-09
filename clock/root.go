package clock

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Root struct {
	app.Compo

	clock *Clock
}

func (r *Root) OnMount(ctx app.Context) {
	r.clock = &Clock{}
}

func (r *Root) Render() app.UI {
	root := app.Div().
		Class("has-background-dark").
		Class("columns").
		Body(
			app.Div().
				Class("column").
				Class("is-two-thirds").
				Class("is-offset-2").
				Body(
					r.clock,
				),
		)

	return root
}
