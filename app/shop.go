package main

import (
	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

func shop(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	sTemplate := MustGetTemplate("sale-items")
	sTemplate.ExecuteEl(doc.QuerySelector(".jass-sale-items"), &Session.Items)
	fadeIn("jass-sale-items")
	noButtons()
}
