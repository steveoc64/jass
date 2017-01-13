package main

import (
	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

func privacy(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	ldTemplate("privacy", ".jass-about", nil)
	// fadeIn("jass-about", "jass-options")
	fadeIn("jass-about")
	noButtons()
	w.ScrollTo(0, 0)

	doc.QuerySelector(".jass-about").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		if evt.Target().TagName() != "I" {
			Session.Navigate("/")
		}
	})
}
