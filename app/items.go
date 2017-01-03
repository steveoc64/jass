package main

import (
	"./shared"
	"honnef.co/go/js/dom"
)

func getItems() {
	Session.Items = []shared.Item{}
	GetJSON("/api/items", &Session.Items, func() {
		print("/api/items complete", Session.Items)
		w := dom.GetWindow()
		doc := w.Document()

		doc.QuerySelector(".jass-sales-bar").AddEventListener("click", false, func(evt dom.Event) {
			print("clicksed on sales bar")
			Session.Navigate("/cart")
		})
	})
}
