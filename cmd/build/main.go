package main

import (
	"flag"
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pthomison/goapp-clock/clock"
)

func main() {
	out := flag.String("out", "docs", "Out directory for static build")
	path := flag.String("path", "", "Base path for static build")
	flag.Parse()

	app.RouteFunc("/", func() app.Composer {
		return &clock.Root{}
	})
	h := clock.App()

	// Deploy under a path
	if *path != "" {
		// test
		h.Resources = app.GitHubPages(*path)
	}

	if err := app.GenerateStaticWebsite(*out, h); err != nil {
		log.Fatalf("could not build: %v\n", err)
	}
}
