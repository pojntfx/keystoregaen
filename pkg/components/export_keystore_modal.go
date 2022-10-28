package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygaen/pkg/components"
)

const (
	exportKeystoreForm = "export-keystore-form"
)

type ExportKeystoreModal struct {
	app.Compo

	OnDownloadKey func(base64encode bool)
	OnViewKey     func()

	OnOK func()

	base64encode bool
}

func (c *ExportKeystoreModal) Render() app.UI {
	return &components.Modal{
		ID:    "export-key-modal",
		Title: "Export Keystore",
		Body: []app.UI{
			app.Div().
				Class("pf-c-card pf-m-compact pf-m-flat").
				Body(
					app.Div().
						Class("pf-c-card__body").
						Body(
							app.Form().
								Class("pf-c-form").
								ID(exportKeystoreForm).
								OnSubmit(func(ctx app.Context, e app.Event) {
									e.PreventDefault()
								}).
								Body(
									app.Div().
										Aria("role", "group").
										Class("pf-c-form__group").
										Body(
											app.Div().
												Class("pf-c-form__group-control").
												Body(
													app.Div().
														Class("pf-c-check").
														Body(
															&components.Controlled{
																Component: app.Input().
																	Class("pf-c-check__input").
																	Type("checkbox").
																	ID("base64-checkbox").
																	OnInput(func(ctx app.Context, e app.Event) {
																		c.base64encode = !c.base64encode
																	}),
																Properties: map[string]interface{}{
																	"checked": c.base64encode,
																},
															},
															app.Label().
																Class("pf-c-check__label").
																For("base64-checkbox").
																Body(
																	app.I().
																		Class("fas fa-shield-alt pf-u-mr-sm"),
																	app.Text("Base64 encode"),
																),
															app.Span().
																Class("pf-c-check__description").
																Text("To increase portability, base64-encode the keystore."),
														),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-card__footer").
						Body(
							app.Button().
								Class(func() string {
									classes := "pf-c-button pf-m-control pf-u-mr-sm pf-u-display-block pf-u-display-inline-block-on-md pf-u-w-100 pf-u-w-initial-on-md"
									if c.base64encode {
										classes += " pf-u-mb-md pf-u-mb-0-on-md"
									}

									return classes
								}()).
								Type("submit").
								Form(exportKeystoreForm).
								OnClick(func(ctx app.Context, e app.Event) {
									c.OnDownloadKey(c.base64encode)
								}).
								Body(
									app.Span().
										Class("pf-c-button__icon pf-m-start").
										Body(
											app.I().
												Class("fas fa-download").
												Aria("hidden", true),
										),
									app.Text("Download keystore"),
								),
							app.If(
								c.base64encode,
								app.Button().
									Class("pf-c-button pf-m-control pf-u-mr-sm pf-u-display-block pf-u-display-inline-block-on-md pf-u-w-100 pf-u-w-initial-on-md").
									Type("submit").
									Form(exportKeystoreForm).
									OnClick(func(ctx app.Context, e app.Event) {
										c.OnViewKey()
									}).
									Body(
										app.Span().
											Class("pf-c-button__icon pf-m-start").
											Body(
												app.I().
													Class("fas fa-eye").
													Aria("hidden", true),
											),
										app.Text("View keystore"),
									),
							),
						),
				),
		},
		Footer: []app.UI{
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("button").
				Text("OK").
				OnClick(func(ctx app.Context, e app.Event) {
					c.clear()
					c.OnOK()
				}),
		},
		OnClose: func() {
			c.clear()
			c.OnOK()
		},
	}
}

func (c *ExportKeystoreModal) clear() {
	c.base64encode = false
}
