package main

import (
	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
	"time"
)

func discover(context *router.Context) {
	print("in discover function")

	go func() {
		w := dom.GetWindow()
		doc := w.Document()

		// swap out the splash box for the model cycle
		doc.QuerySelector(".jass-model-cycle").Class().SetString("jass-model-cycle fade-in fast")
		if !doc.QuerySelector(".jass-splash-box").Class().Contains("hidden") {
			doc.QuerySelector(".jass-splash-box").Class().SetString("jass-splash-box fade-out fast")
			time.Sleep(200 * time.Millisecond)
			doc.QuerySelector(".jass-splash-box").Class().Add("hidden")
		}
		if !doc.QuerySelector(".jass-sale-items").Class().Contains("hidden") {
			doc.QuerySelector(".jass-sale-items").Class().SetString("jass-sale-items fade-out fast")
			time.Sleep(200 * time.Millisecond)
			doc.QuerySelector(".jass-sale-items").Class().Add("hidden")
		}
	}()
}
