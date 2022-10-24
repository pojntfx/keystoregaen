package components

import (
	"bytes"
	"log"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygaen/pkg/components"
	"github.com/pojntfx/keystoregaen/pkg/utils"
)

type Home struct {
	app.Compo

	err error
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
							&KeystoreGenerationForm{
								OnSubmit: func(storepass, keypass, alias, cname string, validity, bits uint32) {
									out := bytes.NewBuffer([]byte{})

									go func() {
										log.Println("Generating keystore ...")

										if err := utils.GenerateKeystore(
											storepass,
											keypass,
											alias,
											cname,
											time.Hour*24*time.Duration(validity),
											bits,
											out,
										); err != nil {
											c.panic(err)

											return
										}

										c.download(out.Bytes(), "keystoregaen.jks", "application/octet-stream")
									}()
								},
							},
						),
				),
			app.If(
				c.err != nil,
				&components.ErrorModal{
					ID:          "error-modal",
					Icon:        "fas fa-times",
					Title:       "An Error Occurred",
					Class:       "pf-m-danger",
					Body:        "The following details may be of help:",
					Error:       c.err,
					ActionLabel: "Close",

					OnClose: func() {
						c.recover()
					},
					OnAction: func() {
						c.recover()
					},
				},
			),
		)
}

func (c *Home) panic(err error) {
	log.Println(err)

	c.err = err

	c.Update()
}

func (c *Home) recover() {
	c.err = nil
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
