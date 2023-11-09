package timer

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func App() *app.Handler {
	return &app.Handler{
		Name:        "Timer",
		Description: "This is a timer",
		Styles: []string{
			"web/index.css",
		},
		Scripts: []string{
			"web/index.js            ",
		},
	}
}
