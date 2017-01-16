package main

import (
	"strconv"

	"./shared"
	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

func drawSalesBar() {
	w := dom.GetWindow()
	doc := w.Document()

	ldTemplate("sales-bar", ".jass-sales-bar", &Session)

	doc.QuerySelector(".jass-sales-bar").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		print("clicksed on sales bar")
		Session.Navigate("/cart")
		hideFooter()
	})
}

func drawCatBar() {
	w := dom.GetWindow()
	doc := w.Document()

	ldTemplate("sale-cats", ".jass-sale-cats", &Session)

	doc.QuerySelector(".jass-sale-cats").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		t := evt.Target()
		if t.TagName() == "SPAN" {
			cat_id, _ := strconv.Atoi(t.GetAttribute("data-id"))
			setCategory(t.InnerHTML(), cat_id)
		}
	})

	for _, el := range doc.QuerySelectorAll(".sale-category") {
		if el.InnerHTML() == Session.SelectCat {
			el.Class().Add("highlight")
		}
	}
}

func setCategory(cat string, cat_id int) {
	w := dom.GetWindow()
	doc := w.Document()

	Session.SelectCat = cat
	Session.SelectCatID = cat_id
	for _, el := range doc.QuerySelectorAll(".sale-category") {
		c := el.Class()
		if el.InnerHTML() == cat {
			c.Add("highlight")
		} else {
			c.Remove("highlight")
		}
	}

	// print("reloading template", Session.SelectCatID)
	ldTemplate("sale-items", ".jass-sale-items", &Session)
}

func hideFooter() {
	w := dom.GetWindow()
	doc := w.Document()
	doc.QuerySelector(".jass-footer").Class().Add("hidden")
}

func shop(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	// Session.SelectCat = ""
	// Session.SelectCatID = 0

	Session.Category = []shared.Category{}
	GetJSON("/api/category", &Session.Category, func() {
		print("got cats", Session.Category)
		Session.Products = []shared.Product{}
		GetJSON("/api/products", &Session.Products, func() {
			print("/api/products complete", Session.Products)
			drawSalesBar()
			drawCatBar()
			hideFooter()

			ldTemplate("sale-items", ".jass-sale-items", &Session)

			// fade in multiple elements
			fadeIn("jass-sales-bar", "jass-sale-cats", "jass-sale-items")
			noButtons()

			doc.QuerySelector(".jass-logo-small").AddEventListener("click", false, func(evt dom.Event) {
				Session.Navigate("/")
			})
		})
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
