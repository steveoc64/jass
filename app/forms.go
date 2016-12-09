package main

import "honnef.co/go/js/dom"

func initForms() {
	w := dom.GetWindow()
	doc := w.Document()

	doc.QuerySelector(".hamburger").AddEventListener("click", false, func(evt dom.Event) {
		print("clicked on burger")
		doc.QuerySelector(".hamburger").Class().Toggle("is-active")
		// doc.QuerySelector("#slidemenu").Class().Toggle("cbp-spmenu-open")
	})

	// doc.QuerySelector("#homepage").AddEventListener("click", false, func(evt dom.Event) {
	// 	evt.PreventDefault()
	// 	Session.Navigate("/")
	// })

}
