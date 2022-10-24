package components

import (
	"log"
	"strconv"

	app "github.com/maxence-charriere/go-app/v9/pkg/app"
)

type KeystoreGenerationForm struct {
	app.Compo

	storepass,
	keypass,
	alias,
	cname,
	validity,
	bits string

	OnSubmit func(
		storepass,
		keypass,
		alias,
		cname string,
		validity,
		bits uint32,
	)
}

func (c *KeystoreGenerationForm) Render() app.UI {
	return app.Div().
		Class("pf-c-card keystore-generation-form").
		Body(
			app.Div().
				Class("pf-c-card__body").
				Body(
					app.Form().
						Class("pf-c-form").
						OnSubmit(func(ctx app.Context, e app.Event) {
							e.PreventDefault()

							validity, err := strconv.ParseUint(c.validity, 10, 32)
							if err != nil {
								log.Println("Could not parse validity:", err)

								return
							}

							bits, err := strconv.ParseUint(c.bits, 10, 32)
							if err != nil {
								log.Println("Could not parse bits:", err)

								return
							}

							c.OnSubmit(
								c.storepass,
								c.keypass,
								c.alias,
								c.cname,
								uint32(validity),
								uint32(bits),
							)

							c.cancel()
						}).
						Body(
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("storepass-input").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Body(
															app.Text("Keystore password"),
														),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Body(
															app.Text("*"),
														),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Required(true).
												Type("password").
												ID("storepass-input").
												Name("storepass-input").
												OnInput(func(ctx app.Context, e app.Event) {
													c.storepass = ctx.JSSrc().Get("value").String()
												}).
												Value(c.storepass),
										),
								),
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("keypass-input").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Body(
															app.Text("Certificate password"),
														),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Body(
															app.Text("*"),
														),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Required(true).
												Type("password").
												ID("keypass-input").
												Name("keypass-input").
												OnInput(func(ctx app.Context, e app.Event) {
													c.keypass = ctx.JSSrc().Get("value").String()
												}).
												Value(c.keypass),
										),
								),
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("alias-input").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Body(
															app.Text("Certificate alias"),
														),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Body(
															app.Text("*"),
														),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Required(true).
												Type("text").
												ID("alias-input").
												Name("alias-input").
												OnInput(func(ctx app.Context, e app.Event) {
													c.alias = ctx.JSSrc().Get("value").String()
												}).
												Value(c.alias),
										),
								),
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("cname-input").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Body(
															app.Text("Certificate CNAME"),
														),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Body(
															app.Text("*"),
														),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Required(true).
												Type("text").
												ID("cname-input").
												Name("cname-input").
												OnInput(func(ctx app.Context, e app.Event) {
													c.cname = ctx.JSSrc().Get("value").String()
												}).
												Value(c.cname),
										),
								),
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("validity-input").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Body(
															app.Text("Certificate validity"),
														),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Body(
															app.Text("*"),
														),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Required(true).
												Type("number").
												ID("validity-input").
												Name("validity-input").
												Aria("describedby", "validity-input-helper").
												OnInput(func(ctx app.Context, e app.Event) {
													c.validity = ctx.JSSrc().Get("value").String()
												}).
												Value(c.validity),
											app.P().
												Class("pf-c-form__helper-text").
												ID("validity-input-helper").
												Aria("live", "polite").
												Body(
													app.Text("Days from now until certificate is no longer valid."),
												),
										),
								),
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("rsa-bits-input").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Body(
															app.Text("RSA bits"),
														),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Body(
															app.Text("*"),
														),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Required(true).
												Type("number").
												ID("rsa-bits-input").
												Name("rsa-bits-input").
												Aria("describedby", "rsa-bits-input-helper").
												OnInput(func(ctx app.Context, e app.Event) {
													c.bits = ctx.JSSrc().Get("value").String()
												}).
												Value(c.bits),
											app.P().
												Class("pf-c-form__helper-text").
												ID("rsa-bits-input-helper").
												Aria("live", "polite").
												Body(
													app.Text("Bits to generate the RSA private key with."),
												),
										),
								),
							app.Div().
								Class("pf-c-form__group pf-m-action pf-u-mt-0").
								Body(
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Div().
												Class("pf-c-form__actions").
												Body(
													app.Button().
														Class("pf-c-button pf-m-primary").
														Type("submit").
														Body(
															app.Text("Generate keystore"),
														),
												),
										),
								),
						),
				),
		)
}

func (c *KeystoreGenerationForm) OnMount(ctx app.Context) {
	c.validity = "365"
	c.bits = "2048"
}

func (c *KeystoreGenerationForm) cancel() {
	c.storepass = ""
	c.keypass = ""
	c.alias = ""
	c.cname = ""
	c.validity = "365"
	c.bits = "2048"
}
