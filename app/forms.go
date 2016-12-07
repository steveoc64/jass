package main

import "honnef.co/go/js/dom"

func initForms() {
	w := dom.GetWindow()
	doc := w.Document()

	doc.QuerySelector("#hamburger").AddEventListener("click", false, func(evt dom.Event) {
		doc.QuerySelector("#hamburger").Class().Toggle("active")
		doc.QuerySelector("#slidemenu").Class().Toggle("cbp-spmenu-open")
		print("clicked on burger")
	})

	// doc.QuerySelector("#homepage").AddEventListener("click", false, func(evt dom.Event) {
	// 	evt.PreventDefault()
	// 	Session.Navigate("/")
	// })

}
