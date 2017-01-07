package main

import (
	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

func about(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	ldTemplate("about", ".jass-about", nil)
	// fadeIn("jass-about", "jass-options")
	fadeIn("jass-about")
	noButtons()
	w.ScrollTo(0, 0)
	doc.QuerySelector(".jass-options").Class().Add("jass-green-top")

	doc.QuerySelector(".jass-about").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		doc.QuerySelector(".jass-options").Class().Remove("jass-green-top")
		if evt.Target().TagName() != "I" {
			Session.Navigate("/")
		}
	})

	addSocialButtons(Session.GetURL()+"/about", "About Chamelee")
}
