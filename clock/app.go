package clock

import (
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func App() *app.Handler {
	return &app.Handler{
		Name:        "Timer",
		Description: "This is a timer",
		Styles: []string{
			"web/index.css",
		},
		Scripts: []string{
			"web/index.js",
		},
		AutoUpdateInterval: time.Second * 1,
		// Env: map[string]string{
		// 	"GOAPP_ROOT_PREFIX": "/docs",
		// },
		// RawHeaders: []string{
		// 	`<meta property="og:url" content="https://pojntfx.github.io/liwasc/">`,
		// 	`<meta property="og:title" content="liwasc">`,
		// 	`<meta property="og:description" content="List, wake and scan nodes in a network.">`,
		// 	`<meta property="og:image" content="https://pojntfx.github.io/liwasc/web/icon.png">`,
		// },
	}
}
