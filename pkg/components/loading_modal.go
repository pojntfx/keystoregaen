package components

import (
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"

	"github.com/pojntfx/keygaen/pkg/components"
)

type LoadingModal struct {
	app.Compo

	Title       string
	Description string

	OnReady func()
}

func (c *LoadingModal) Render() app.UI {
	return &components.Modal{
		ID:    "loading-modal",
		Title: c.Title,
		Body: []app.UI{
			app.Div().
				Class("pf-m-fill pf-l-flex pf-m-column pf-m-justify-content-center pf-m-align-items-center pf-u-mt-md").
				Body(
					app.Raw(`<svg class="pf-c-spinner" role="progressbar" viewBox="0 0 100 100" aria-label="Loading...">
  <circle class="pf-c-spinner__path" cx="50" cy="50" r="45" fill="none" />
</svg>`),
					app.P().
						Class("pf-u-mt-sm").
						Body(
							app.Text(c.Description),
						),
				),
		},
	}
}

func (c *LoadingModal) OnMount(ctx app.Context) {
	<-time.After(time.Millisecond * 50)

	c.OnReady()
}
