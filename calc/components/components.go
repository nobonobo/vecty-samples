package components

import (
	"fmt"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"

	"github.com/nobonobo/vecty-samples/calc/dispatcher"
)

// LCD ...
type LCD struct {
	vecty.Core
	Label string `vecty:"prop"`
}

// Render ...
func (c *LCD) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(vecty.ClassMap{
			"siimple-grid-col":     true,
			"siimple-grid-col--12": true,
		}),
		elem.Input(
			vecty.Markup(
				vecty.ClassMap{
					"siimple-input":        true,
					"siimple-input--fluid": true,
				},
				vecty.Style("font-size", "2em"),
				vecty.Style("text-align", "right"),
				vecty.Style("height", "auto"),
				prop.Value(c.Label),
				vecty.Attribute("readonly", ""),
			),
		),
	)
}

// Button ...
type Button struct {
	vecty.Core
	Label  string            `vecty:"prop"`
	Size   int               `vecty:"prop"`
	Color  string            `vecty:"prop"`
	Action dispatcher.Action `vecty:"prop"`
}

// Render ...
func (c *Button) Render() vecty.ComponentOrHTML {
	if c.Size == 0 {
		c.Size = 3
	}
	if c.Color == "" {
		c.Color = "navy"
	}
	return elem.Button(
		vecty.Markup(
			vecty.ClassMap{
				"siimple-grid-col":                          true,
				fmt.Sprintf("siimple-grid-col--%d", c.Size): true,
				"siimple-btn":                               true,
				"siimple-btn--" + c.Color:                   true,
			},
			vecty.Attribute("align", "center"),
			vecty.Style("font-size", "2em"),
			vecty.Style("pading", "1em"),
			vecty.Style("height", "auto"),
			vecty.MarkupIf(c.Action != nil,
				event.Click(func(ev *vecty.Event) {
					dispatcher.Dispatch(c.Action)
				}),
			),
		),
		vecty.Text(c.Label),
	)
}
