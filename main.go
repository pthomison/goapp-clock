package main

import (
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pthomison/errcheck"
	"github.com/pthomison/goapp-clock/clock"
	"github.com/sirupsen/logrus"
)

const (
	addr = "127.0.0.1:8000"
)

func main() {
	// Components routing:
	app.RouteFunc("/", func() app.Composer {
		return &clock.Root{}
	})

	app.RouteFunc("/goapp-clock/", func() app.Composer {
		return &clock.Root{}
	})

	// app.Route("/goapp-clock", func() app.Composer {
	// 	return &clock.Root{}
	// })

	app.RunWhenOnBrowser()

	// HTTP routing:
	http.Handle("/", clock.App())

	logrus.Info("Listening on ", addr)
	err := http.ListenAndServe(addr, nil)
	errcheck.Check(err)
}
