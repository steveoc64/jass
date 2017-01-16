package main

import (
	"./shared"
	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

func drawSalesBar() {
	w := dom.GetWindow()
	doc := w.Document()

	ldTemplate("sales-bar", ".jass-sales-bar", &Session)
	doc.QuerySelector(".jass-sales-bar").AddEventListener("click", false, func(evt dom.Event) {
		print("clicksed on sales bar")
		Session.Navigate("/cart")
	})
}

func shop(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	Session.Products = []shared.Product{}
	GetJSON("/api/products", &Session.Products, func() {
		print("/api/products complete", Session.Products)
		drawSalesBar()

		// Load up em templates
		sTemplate := MustGetTemplate("sale-items")
		sTemplate.ExecuteEl(doc.QuerySelector(".jass-sale-items"), &Session)

		// fade in multiple elements
		fadeIn("jass-sales-bar", "jass-sale-items")
		noButtons()
	})
}

func cart(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	print("shopping cart")

	// Load up em templates
	sTemplate := MustGetTemplate("cart")
	sTemplate.ExecuteEl(doc.QuerySelector(".jass-cart"), &Session)
	fadeIn("jass-cart")
	noButtons()
}
