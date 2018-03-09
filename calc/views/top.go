package views

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"

	"github.com/nobonobo/vecty-samples/calc/actions"
	"github.com/nobonobo/vecty-samples/calc/components"
	"github.com/nobonobo/vecty-samples/calc/store"
)

// TopView ...
type TopView struct {
	vecty.Core
	State *store.StateDef
}

// Render ...
func (c *TopView) Render() vecty.ComponentOrHTML {
	clearLabel := "Ｃ"
	clearAll := false
	if c.State.IsZero() {
		clearLabel = "AC"
		clearAll = true
	}
	return elem.Body(
		elem.Div(
			vecty.Markup(vecty.ClassMap{
				"siimple-navbar":        true,
				"siimple-navbar--small": true,
				"siimple-navbar--navy":  true,
			}),
			elem.Div(
				vecty.Markup(vecty.ClassMap{
					"siimple-layout--left": true,
				}),
				elem.Div(
					vecty.Markup(vecty.ClassMap{
						"siimple-navbar-title": true,
					}),
					vecty.Text("Calc"),
				),
			),
		),
		elem.Div(
			vecty.Markup(vecty.ClassMap{
				"siimple-content":        true,
				"siimple-content--small": true,
			}),
			elem.Div(
				vecty.Markup(vecty.ClassMap{
					"siimple-grid": true,
				}),
				elem.Div(
					vecty.Markup(vecty.ClassMap{
						"siimple-grid-row": true,
					}),
					&components.LCD{Label: c.State.NowInput},
				),
				elem.Div(
					vecty.Markup(
						vecty.ClassMap{
							"siimple-grid-row": true,
						},
						vecty.Style("padding-top", "1em"),
					),
					&components.Button{Label: clearLabel, Color: "gray", Action: actions.Clear{All: clearAll}},
					&components.Button{Label: "+/-", Color: "gray", Action: actions.Invert{}},
					&components.Button{Label: "％", Color: "gray", Action: actions.Percent{}},
					&components.Button{Label: "÷", Color: "orange", Action: actions.Operator{Char: '/'}},
				), elem.Div(
					vecty.Markup(
						vecty.ClassMap{
							"siimple-grid-row": true,
						},
						vecty.Style("padding-top", "1em"),
					),
					&components.Button{Label: "７", Action: actions.Insert{Char: '7'}},
					&components.Button{Label: "８", Action: actions.Insert{Char: '8'}},
					&components.Button{Label: "９", Action: actions.Insert{Char: '9'}},
					&components.Button{Label: "×", Color: "orange", Action: actions.Operator{Char: '*'}},
				),
				elem.Div(
					vecty.Markup(
						vecty.ClassMap{
							"siimple-grid-row": true,
						},
						vecty.Style("padding-top", "1em"),
					),
					&components.Button{Label: "４", Action: actions.Insert{Char: '4'}},
					&components.Button{Label: "５", Action: actions.Insert{Char: '5'}},
					&components.Button{Label: "６", Action: actions.Insert{Char: '6'}},
					&components.Button{Label: "ー", Color: "orange", Action: actions.Operator{Char: '-'}},
				),
				elem.Div(
					vecty.Markup(
						vecty.ClassMap{
							"siimple-grid-row": true,
						},
						vecty.Style("padding-top", "1em"),
					),
					&components.Button{Label: "１", Action: actions.Insert{Char: '1'}},
					&components.Button{Label: "２", Action: actions.Insert{Char: '2'}},
					&components.Button{Label: "３", Action: actions.Insert{Char: '3'}},
					&components.Button{Label: "＋", Color: "orange", Action: actions.Operator{Char: '+'}},
				),
				elem.Div(
					vecty.Markup(
						vecty.ClassMap{
							"siimple-grid-row": true,
						},
						vecty.Style("padding-top", "1em"),
					),
					&components.Button{Label: "０", Size: 6, Action: actions.Insert{Char: '0'}},
					&components.Button{Label: "．", Action: actions.Insert{Char: '.'}},
					&components.Button{Label: "＝", Color: "orange", Action: actions.Equal{}},
				),
			),
		),
	)
}
