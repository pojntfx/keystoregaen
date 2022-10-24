package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Home struct {
	app.Compo
}

func (c *Home) Render() app.UI {
	return app.Div().
		Class("pf-c-page").
		Body(
			app.A().
				Class("pf-c-skip-to-content pf-c-button pf-m-primary").
				Href("#keystoregaen-main").
				Body(
					app.Text("Skip to content"),
				),
			&Navbar{},
			app.Main().
				Class("pf-c-page__main", "pf-c-page__main").
				ID("keystoregaen-main").
				TabIndex(-1).
				Body(
					app.Div().
						Class("pf-c-page__main-section", "pf-m-fill", "pf-l-flex", "pf-m-justify-content-center", "pf-m-align-items-center").
						Body(
							app.Div().
								Class("pf-l-flex pf-m-column pf-m-justify-content-center pf-m-align-items-center pf-m-space-items-xl").
								Body(
									app.H1().
										Class("pf-u-font-size-3xl pf-u-text-align-center pf-u-mt-md").
										Body(
											app.Text("Hello, world!"),
										),
								),
						),
				),
		)
}

func (c *Home) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}

func (c *Home) download(content []byte, name string, mimetype string) {
	buf := app.Window().JSValue().Get("Uint8Array").New(len(content))
	app.CopyBytesToJS(buf, content)

	blob := app.Window().JSValue().Get("Blob").New(app.Window().JSValue().Get("Array").New(buf), map[string]interface{}{
		"type": mimetype,
	})

	link := app.Window().Get("document").Call("createElement", "a")
	link.Set("href", app.Window().Get("URL").Call("createObjectURL", blob))
	link.Set("download", name)
	link.Call("click")
}
