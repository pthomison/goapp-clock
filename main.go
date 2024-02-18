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
	// serve := flag.Bool("serve", false, "Build and serve the frontend")
	// laddr := flag.String("laddr", "127.0.0.1:8000", "Address to serve the frontend on")

	flag.Parse()

	// Components routing:
	app.RouteFunc("/", func() app.Composer {
		return &clock.Root{}
	})

	// app.RouteFunc("/goapp-clock/", func() app.Composer {
	// 	return &clock.Root{}
	// })

	// // app.Route("/goapp-clock", func() app.Composer {
	// // 	return &clock.Root{}
	// // })

	app.RunWhenOnBrowser()

	h := clock.App()

	if *build {
		// Deploy under a path
		if *path != "" {
			h.Resources = app.GitHubPages(*path)
		}

		// h.Resources = app.CustomProvider(*path, *path)

		if err := app.GenerateStaticWebsite(*out, h); err != nil {
			log.Fatalf("could not build: %v\n", err)
		}
	}

	// if *serve {
	// 	// HTTP routing:
	// 	http.Handle("/", h)
	// 	// http.Handle("/mygoapp", http.StripPrefix("/mygoapp", h))

	// 	logrus.Info("Listening on ", *laddr)
	// 	err := http.ListenAndServe(*laddr, nil)
	// 	errcheck.Check(err)
	// }

}
