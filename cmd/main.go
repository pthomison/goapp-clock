package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pthomison/errcheck"
	"github.com/pthomison/goapp-clock/clock"
	"github.com/sirupsen/logrus"
)

func main() {
	build := flag.Bool("build", false, "Create static build")
	serve := flag.Bool("serve", false, "Run the backend server")

	out := flag.String("out", "docs", "Out directory for static build")
	path := flag.String("path", "", "Base path for static build")
	addr := flag.String("addr", "127.0.0.1:8000", "Address to listen on")
	flag.Parse()

	h := clock.App()

	BrowserRouting()

	if *build {
		Build(h, *path, *out)
	}

	if *serve {
		Serve(h, *addr)
	}
}

func BrowserRouting() {
	app.RouteFunc("/", func() app.Composer {
		return &clock.Root{}
	})
	app.RunWhenOnBrowser()
}

func Serve(h *app.Handler, addr string) {
	http.Handle("/", h)

	logrus.Info("Listening on ", addr)
	err := http.ListenAndServe(addr, nil)
	errcheck.Check(err)
}

func Build(h *app.Handler, path string, out string) {
	// Deploy under a path
	if path != "" {
		// test
		h.Resources = app.GitHubPages(path)
	}

	if err := app.GenerateStaticWebsite(out, h); err != nil {
		log.Fatalf("could not build: %v\n", err)
	}
}
