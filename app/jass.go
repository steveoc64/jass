package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
	"honnef.co/go/js/dom"
)

var jQuery = jquery.NewJQuery

func main() {
	w := dom.GetWindow()
	doc := w.Document()

	w.ScrollTo(0, 0)

	initRouter()
	initForms()
	initBurger()

	Session.LastWidth = dom.GetWindow().InnerWidth()
	Session.Orientation = "Landscape"
	if dom.GetWindow().InnerHeight() > dom.GetWindow().InnerWidth() {
		Session.Orientation = "Portrait"
	}
	if Session.Mobile() {
		Session.wasMobile = true
	}
	if Session.SubMobile() {
		Session.wasSubmobile = true
	}

	js.Global.Set("resize", func() {
		Session.Resize()
	})

	getItems()
	getBlogs()
	doSplashPage()
	showTopMenu()

	doc.QuerySelector("#option-shop").AddEventListener("click", false, func(evt dom.Event) {
		Session.Navigate("/shop")
	})

	doc.QuerySelector("#option-discover").AddEventListener("click", false, func(evt dom.Event) {
		Session.Navigate("/discover")
	})

	// doc.QuerySelector("#option-merchandise").AddEventListener("click", false, func(evt dom.Event) {
	// 	Session.Navigate("/merchandise")
	// })

	// doc.QuerySelector("#option-stories").AddEventListener("click", false, func(evt dom.Event) {
	// 	Session.Navigate("/stories")
	// })

	// map options to routes
	doc.QuerySelector("[name=opt-b]").AddEventListener("click", false, func(evt dom.Event) {
		Session.Navigate("/blog")
	})
	doc.QuerySelector("[name=opt-m]").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		js.Global.Get("location").Set("href", "https://shop.polymer-project.org/list/ladies_outerwear")
	})

	print("Your current jQuery version is: " + jQuery().Jquery)
}

func getDivOffset(el dom.Element) int {
	retval := float64(0.0)
	pel := el.(dom.HTMLElement).OffsetParent()
	if pel != nil {
		for {
			retval += el.(dom.HTMLElement).OffsetTop()
			el = el.(dom.HTMLElement).OffsetParent()
			if el == nil {
				return int(retval)
			}
		}
	}
	return int(retval)
}
