package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keystoregaen/pkg/components"
)

func main() {
	// Client-side code
	{
		app.Route("/", &components.Home{})
		app.RunWhenOnBrowser()
	}

	// Server-/build-side code

	// Parse the flags
	serve := flag.Bool("serve", false, "Serve the app instead of building it")
	laddr := flag.String("laddr", "0.0.0.0:21255", "Address to listen on when serving the app")
	dist := flag.String("dist", "build", "Directory to build the app to")
	prefix := flag.String("prefix", "/keystoregaen", "Prefix to build the app for")

	flag.Parse()

	// Define the handler
	h := &app.Handler{
		Title:           "keystoregaen",
		Name:            "keystoregaen",
		ShortName:       "keystoregaen",
		Description:     "Generate Java keystores in your browser.",
		LoadingLabel:    "Generate Java keystores in your browser.",
		Author:          "Felicitas Pojtinger",
		ThemeColor:      "#151515",
		BackgroundColor: "#151515",
		Icon: app.Icon{
			Default: "/web/default.png",
			Large:   "/web/large.png",
		},
		Keywords: []string{
			"java",
			"jkcs",
			"keystore",
			"android",
			"android-certificate",
		},
		RawHeaders: []string{
			`<meta property="og:url" content="https://pojntfx.github.io/keystoregaen/">`,
			`<meta property="og:title" content="keystoregaen">`,
			`<meta property="og:description" content="Generate Java keystores in your browser.">`,
			`<meta property="og:image" content="https://pojntfx.github.io/keystoregaen/web/default.png">`,
		},
		Styles: []string{
			"/web/main.css",
		},
	}

	// Serve if specified
	if *serve {
		log.Println("Listening on", *laddr)

		if err := http.ListenAndServe(*laddr, h); err != nil {
			log.Fatal("could not serve:", err)
		}

		return
	}

	// Build if not specified
	h.Resources = app.GitHubPages(*prefix)

	if err := app.GenerateStaticWebsite(*dist, h); err != nil {
		log.Fatal("could not build static website:", err)
	}
}
