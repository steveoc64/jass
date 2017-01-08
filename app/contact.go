package main

import (
	"github.com/go-humble/router"
	// "honnef.co/go/js/dom"
)

func contact(context *router.Context) {
	// w := dom.GetWindow()
	// doc := w.Document()

	print("in contact function")
	ldTemplate("contact", ".jass-contact", &Session)
	fadeIn("jass-contact")
	noButtons()
}
