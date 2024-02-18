package main

import (
	"flag"
	"log"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pthomison/goapp-clock/clock"
)

func main() {
	build := flag.Bool("build", false, "Create static build")
	out := flag.String("out", "docs", "Out directory for static build")
	path := flag.String("path", "", "Base path for static build")
	flag.Parse()

	{ // Browser Execution
		app.RouteFunc("/", func() app.Composer {
			return &clock.Root{}
		})
		app.RunWhenOnBrowser()
	}

	if *build {
		h := clock.App()

		// Deploy under a path
		if *path != "" {
			h.Resources = app.GitHubPages(*path)
		}

		if err := app.GenerateStaticWebsite(*out, h); err != nil {
			log.Fatalf("could not build: %v\n", err)
		}
	}

}
