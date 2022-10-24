package components

import (
	"bytes"
	"log"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygaen/pkg/components"
	"github.com/pojntfx/keystoregaen/pkg/utils"
)

const (
	auditStorageKey = "keystoregaenAudit"
)

type Home struct {
	app.Compo

	err     error
	loading bool

	loadingReady chan struct{}

	removeEventListeners []func()

	showAuditModal bool
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

									c.setLoading(true)

									go func() {
										<-c.loadingReady

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

										c.setLoading(false)
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
			app.If(
				c.showAuditModal,
				&components.ConfirmationModal{
					ID:    "audit-modal",
					Icon:  "fas fa-exclamation-triangle",
					Title: "keystoregaen has not yet been audited!",
					Class: "pf-m-warning",
					Body:  "While we try to make keystoregaen as secure as possible, it has not yet undergone a formal security audit by a third party. Please keep this in mind if you use it for security-critical applications.",

					ActionLabel: "Yes, I understand",
					ActionClass: "pf-m-warning",

					CancelLink:  "https://en.wikipedia.org/wiki/Information_security_audit",
					CancelLabel: "What is an audit?",

					OnClose: func() {
						c.showAuditModal = false
						c.writeToLocalStorage()

						c.Update()
					},
					OnAction: func() {
						c.showAuditModal = false
						c.writeToLocalStorage()

						c.Update()
					},
				},
			),
			app.If(
				c.loading,
				&LoadingModal{
					Title:       "Generating Keystore",
					Description: "Calculating primes ...",

					OnReady: func() {
						c.loadingReady <- struct{}{}
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

func (c *Home) setLoading(loading bool) {
	c.loading = loading

	c.Update()
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

func (c *Home) readFromLocalStorage() {
	if showAuditModal := app.Window().Get("localStorage").Call("getItem", auditStorageKey); showAuditModal.IsNull() || showAuditModal.IsUndefined() || showAuditModal.String() == "true" {
		c.showAuditModal = true
	}
}

func (c *Home) writeToLocalStorage() {
	app.Window().Get("localStorage").Call("setItem", auditStorageKey, c.showAuditModal)
}

func (c *Home) OnMount(ctx app.Context) {
	c.loadingReady = make(chan struct{})

	c.readFromLocalStorage()

	c.removeEventListeners = []func(){
		app.Window().AddEventListener("storage", func(ctx app.Context, e app.Event) { // This event only fires in other tabs; it does not lead to local race conditions with c.writeKeysToLocalStorage
			c.readFromLocalStorage()

			c.Update()
		}),
	}
}

func (c *Home) OnDismount() {
	if c.removeEventListeners != nil {
		for _, clearListener := range c.removeEventListeners {
			clearListener()
		}
	}
}
