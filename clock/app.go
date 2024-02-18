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
		Icon: app.Icon{
			Default:    "/web/static/android-chrome-192x192.png",
			Large:      "/web/static/android-chrome-512x512.png",
			SVG:        "/web/static/favicon.ico",
			AppleTouch: "/web/static/apple-touch-icon.png",
		},
	}
}
