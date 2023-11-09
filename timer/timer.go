package timer

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Timer struct {
	app.Compo

	name string
}

func (t *Timer) Render() app.UI {
	return app.Div().Body(
		app.H1().Body(
			app.Text("Hello, "),
			app.If(t.name != "",
				app.Text(t.name),
			).Else(
				app.Text("World!"),
			),
		),
		app.P().Body(
			app.Input().
				Type("text").
				Class("input").
				Value(t.name).
				Placeholder("What is your name?").
				AutoFocus(true).
				OnChange(t.ValueTo(&t.name)),
		),
	)
}
