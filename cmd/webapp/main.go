package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pthomison/goapp-clock/clock"
)

func main() {
	app.RouteFunc("/", func() app.Composer {
		return &clock.Root{}
	})
	app.RunWhenOnBrowser()
}
