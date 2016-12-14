package main

import (
	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

func shop(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	// Load up em templates
	sTemplate := MustGetTemplate("sales-bar")
	sTemplate.ExecuteEl(doc.QuerySelector(".jass-sales-bar"), &Session)

	sTemplate = MustGetTemplate("sale-items")
	sTemplate.ExecuteEl(doc.QuerySelector(".jass-sale-items"), &Session)

	// fade in multiple elements
	fadeIn("jass-sales-bar", "jass-sale-items")
	noButtons()
}
