package main

import (
	"time"

	"honnef.co/go/js/dom"
)

func fadeIn() {
	print("fading in")

	w := dom.GetWindow()
	doc := w.Document()
	time.Sleep(2 * time.Second)
	// 2 seconds after load

	doc.QuerySelector(".jass-logo").Class().SetString("jass-logo fade-out fast")
	time.Sleep(200 * time.Millisecond)
	// .. wait for it

	doc.QuerySelector(".jass-logo-box").Class().Add("hidden")

	doc.QuerySelector(".jass-splash-image").(*dom.HTMLImageElement).Src = "img/models/model-001.jpg"
	doc.QuerySelector(".jass-splash-image").Class().SetString("jass-splash-image fade-in fast")
	time.Sleep(500 * time.Millisecond)
	// .. wait for the splash image to fade in

	doc.QuerySelector(".jass-splash-image").Class().Add("green-base")

	// and finally display the top logo
	doc.QuerySelector(".jass-logo-top").Class().Remove("hidden")
	doc.QuerySelector(".jass-logo-top").AddEventListener("click", false, func(evt dom.Event) {
		println("clicked on logo")
		w.ScrollTo(0, 0)
	})
}
