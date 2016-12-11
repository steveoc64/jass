package main

import (
	"time"

	"honnef.co/go/js/dom"
)

var elements = []string{
	"jass-logo-box",
	"jass-splash-box",
	"jass-sale-items",
	"jass-model-cycle",
}

func fadeIn(element string) {
	// print("fade in element", element)
	w := dom.GetWindow()
	doc := w.Document()

	// fade in the initial element
	doc.QuerySelector("." + element).Class().SetString(element + " fade-in fast")

	// fade out everything else
	gotSome := false
	for _, v := range elements {
		if v != element {
			c := doc.QuerySelector("." + v).Class()
			if !c.Contains("hidden") {
				gotSome = true
				// print("fade out element ", k, v)
				c.SetString(v + " fade-out fast")
			}
		}
	}

	if gotSome {
		go func() {
			time.Sleep(200 * time.Millisecond)
			for _, v := range elements {
				if v != element {
					// print("hiding", v)
					doc.QuerySelector("." + v).Class().Add("hidden")
				}
			}
		}()
	}

}
