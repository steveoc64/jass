package main

import (
	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
	"time"
)

func shop(context *router.Context) {
	print("in shop function")

	go func() {
		w := dom.GetWindow()
		doc := w.Document()

		// swap out the splash box for the model cycle

		sTemplate := MustGetTemplate("sale-items")
		print("sTemplate", sTemplate)
		sTemplate.ExecuteEl(doc.QuerySelector(".jass-sale-items"), nil)

		doc.QuerySelector(".jass-sale-items").Class().SetString("jass-sale-items fade-in fast")
		if !doc.QuerySelector(".jass-splash-box").Class().Contains("hidden") {
			doc.QuerySelector(".jass-splash-box").Class().SetString("jass-splash-box fade-out one")
			time.Sleep(200 * time.Millisecond)
			doc.QuerySelector(".jass-splash-box").Class().Add("hidden")
		}
		if !doc.QuerySelector(".jass-model-cycle").Class().Contains("hidden") {
			doc.QuerySelector(".jass-model-cycle").Class().SetString("jass-model-cycle fade-out one")
			time.Sleep(200 * time.Millisecond)
			doc.QuerySelector(".jass-model-cycle").Class().Add("hidden")
		}

		// load a template for the sales items
	}()
}
