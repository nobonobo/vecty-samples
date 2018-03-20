package main

import (
	"bytes"

	"github.com/gopherjs/vecty/event"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"

	"github.com/nobonobo/vecty-samples/html2go/trans"
)

// TopView ...
type TopView struct {
	vecty.Core
	Input    string
	WithBody bool
	Output   string
}

// Render ...
func (c *TopView) Render() vecty.ComponentOrHTML {
	vecty.SetTitle("Html2Go:Top")
	return elem.Body(
		elem.Navigation(
			vecty.Markup(
				vecty.ClassMap{
					"navbar":           true,
					"navbar-expand-lg": true,
					"bg-light":         true,
					"navbar-light":     true,
					"fixed-top":        true,
				},
			),
			elem.Anchor(
				vecty.Markup(
					prop.Href("/"),
					vecty.Class("navbar-brand"),
				),
				vecty.Text("Html2Go(for Vecty)"),
			),
			elem.Button(
				vecty.Markup(
					prop.Type("button"),
					vecty.Class("navbar-toggler"),
					vecty.Attribute("data-toggle", "collapse"),
					vecty.Attribute("data-target", "#menu"),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("navbar-toggler-icon"),
					),
				),
			),
			elem.Div(
				vecty.Markup(
					prop.ID("menu"),
					vecty.ClassMap{
						"navbar-collapse": true,
						"collapse":        true,
					},
				),
				elem.UnorderedList(
					vecty.Markup(
						vecty.Class("navbar-nav"),
					),
					elem.ListItem(
						vecty.Markup(
							vecty.ClassMap{
								"nav-item": true,
								"active":   true,
							},
						),
						elem.Anchor(
							vecty.Markup(
								prop.Href("#/"),
								vecty.Class("nav-link"),
							),
							vecty.Text("Top"),
						),
					),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("container"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("row"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("col", "col-12", "col-md-6"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("card"),
						),
						elem.Div(
							vecty.Markup(
								vecty.Class("card-body"),
							),
							elem.Heading5(
								vecty.Markup(
									vecty.Class("card-title"),
								),
								vecty.Text("Html"),
							),
							elem.Paragraph(
								elem.TextArea(
									vecty.Markup(
										prop.ID("html"),
										vecty.Class("form-control"),
										event.Input(func(ev *vecty.Event) {
											c.Input = ev.Target.Get("value").String()
											c.Update()
											vecty.Rerender(c)
										}),
									),
									vecty.Text(c.Input),
								),
								elem.Div(
									vecty.Markup(
										vecty.Class("form-check"),
									),
									elem.Input(
										vecty.Markup(
											prop.ID("withBody"),
											prop.Type(prop.TypeCheckbox),
											vecty.Class("form-check-input"),
											prop.Checked(c.WithBody),
											event.Click(func(ev *vecty.Event) {
												c.WithBody = !c.WithBody
												c.Update()
												vecty.Rerender(c)
											}),
										),
									),
									elem.Label(
										vecty.Markup(
											vecty.Class("form-check-label"),
										),
										vecty.Text("With Body"),
									),
								),
							),
						),
					),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("col", "col-12", "col-md-6"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("card"),
						),
						elem.Div(
							vecty.Markup(
								vecty.Class("card-body"),
							),
							elem.Heading5(
								vecty.Markup(
									vecty.Class("card-title"),
								),
								vecty.Text("Go(for Vecty)"),
							),
							elem.Paragraph(
								elem.Preformatted(
									elem.Code(vecty.Text(c.Output)),
								),
							),
						),
					),
				),
			),
		),
	)
}

// Update ...
func (c *TopView) Update() {
	in := bytes.NewBufferString(c.Input)
	out := bytes.NewBuffer(nil)
	if err := trans.Trans(out, in, c.WithBody); err != nil {
		c.Output = err.Error()
		vecty.Rerender(c)
		return
	}
	c.Output = out.String()
}

const sample = `<div class="container">
      <div class="py-5 text-center">
        <img class="d-block mx-auto mb-4" src="https://getbootstrap.com/assets/brand/bootstrap-solid.svg" alt="" width="72" height="72">
        <h2>Checkout form</h2>
        <p class="lead">Below is an example form built entirely with Bootstrap's form controls. Each required form group has a validation state that can be triggered by attempting to submit the form without completing it.</p>
      </div>
`

func main() {
	view := &TopView{Input: sample}
	view.Update()
	vecty.RenderBody(view)
}
