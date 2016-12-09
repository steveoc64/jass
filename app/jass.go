package main

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func main() {
	dom.GetWindow().ScrollTo(0, 0)

	initRouter()
	initForms()

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

	fadeIn()

	// cyclePhotos()
}
