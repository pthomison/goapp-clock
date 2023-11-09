package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pthomison/goapp-timer/timer"
)

func main() {
	// Components routing:
	app.Route("/", &timer.Timer{})
	app.RunWhenOnBrowser()

	// HTTP routing:
	http.Handle("/", timer.App())

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
