package main

import (
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gopherjs/vecty"

	"github.com/nobonobo/vecty-samples/calc/actions"
	"github.com/nobonobo/vecty-samples/calc/dispatcher"
	"github.com/nobonobo/vecty-samples/calc/router"
	"github.com/nobonobo/vecty-samples/calc/store"
	"github.com/nobonobo/vecty-samples/calc/views"
)

func handler(act dispatcher.Action) {
	log.Printf("%#v", act)
	switch v := act.(type) {
	default:
		log.Println("unknown action:", v)
		return
	case actions.Insert:
		if v.Char == '.' {
			if strings.Contains(store.State.NowInput, ".") {
				return
			}
		} else {
			if store.State.IsZero() {
				store.State.NowInput = store.State.NowInput[:len(store.State.NowInput)-1]
			}
		}
		store.State.NowInput += string(v.Char)
		store.State.Modified = true
	case actions.Invert:
		if strings.HasPrefix(store.State.NowInput, "-") {
			store.State.NowInput = store.State.NowInput[1:]
		} else {
			store.State.NowInput = "-" + store.State.NowInput
		}
	case actions.Clear:
		store.State.NowInput = "0"
		if !store.State.Modified {
			store.State.Operator = rune(0)
		}
		if v.All {
			store.State.LastInput = ""
		}
	case actions.Percent:
		val, err := store.State.Now()
		if err != nil {
			log.Println(err)
			return
		}
		val /= 100
		store.State.NowInput = strconv.FormatFloat(val, 'G', -1, 64)
	case actions.Operator:
		switch v.Char {
		default:
			log.Println("unsuported operator:", v.Char)
			return
		case '+':
		case '-':
		case '*':
		case '/':
		}
		store.State.Operator = v.Char
		store.State.LastInput = store.State.NowInput
		store.State.NowInput = "0"
		return
	case actions.Equal:
		if store.State.Operator != rune(0) {
			var result float64
			now, err := store.State.Now()
			if err != nil {
				log.Println(err)
				return
			}
			last, err := store.State.Last()
			if err != nil {
				log.Println(err)
				return
			}
			switch store.State.Operator {
			default:
				log.Println("unsuported operator:", store.State.Operator)
				return
			case '+':
				result = last + now
			case '-':
				result = last - now
			case '*':
				result = last * now
			case '/':
				if now != 0 {
					result = math.Round(last / now)
				} else {
					store.State.NowInput = "ERR"
					break
				}
			}
			if store.State.Modified {
				store.State.Modified = false
				store.State.LastInput = store.State.NowInput
			}
			store.State.NowInput = strconv.FormatFloat(result, 'G', -1, 64)
			time.AfterFunc(30*time.Millisecond, func() {
				store.State.LastInput = store.State.NowInput
				store.State.NowInput = "0"
			})
		}
	}
	router.Rerender()
}

func main() {
	vecty.AddStylesheet("https://cdn.jsdelivr.net/npm/siimple@3.0.0/dist/siimple.min.css")
	router.HandleFunc("/", func(ctx *router.Context) {
		router.RenderBody(&views.TopView{State: store.State})
	})
	dispatcher.Register(handler)
	router.Start()
}
