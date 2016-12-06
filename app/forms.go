package main

import (
	"honnef.co/go/js/dom"
)

func initForms() {
	w := dom.GetWindow()
	doc := w.Document()

	doc.QuerySelector("#hamburger").AddEventListener("click", false, func(evt dom.Event) {
		el := doc.QuerySelector("#slidemenu")
		el.Class().Add("cbp-spmenu-open")
	})

	doc.QuerySelector("#homepage").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		Session.Navigate("/")
	})

}
