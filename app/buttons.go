package main

import (
	"honnef.co/go/js/dom"
)

var allBtns = []string{
	// "stories",
	"shop",
	"discover",
	// "merchandise",
}

var btnSpeeds = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
}

func noButtons() {
	w := dom.GetWindow()
	doc := w.Document()

	doc.QuerySelector(".jass-option-bar").Class().SetString("jass-option-bar hidden")
	closeBurger()
}

func showButtons(btns ...string) {
	w := dom.GetWindow()
	doc := w.Document()

	for _, v := range allBtns {
		doc.QuerySelector("#option-" + v).Class().SetString("hidden")
	}

	if len(btns) > 0 {
		doc.QuerySelector(".jass-option-bar").Class().SetString("jass-option-bar fade-in fast")

		for k, v := range btns {
			doc.QuerySelector("#option-" + v).Class().SetString("button button-outline jass-option-button fade-in " + btnSpeeds[k])
		}
	} else {
		doc.QuerySelector(".jass-option-bar").Class().SetString("jass-option-bar hidden")
	}
	closeBurger()

}
