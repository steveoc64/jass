package main

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

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
}
