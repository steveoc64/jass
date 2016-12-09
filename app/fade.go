package main

import (
	"time"

	"honnef.co/go/js/dom"
)

func fadeIn() {
	print("fading in")

	w := dom.GetWindow()
	doc := w.Document()
	// time.Sleep(1 * time.Second)
	// 2 seconds after load

	doc.QuerySelector(".jass-logo-container").Class().SetString("jass-logo-container fade-out fast")
	// time.Sleep(200 * time.Millisecond)
	// .. wait for it

	fadeInSplashBox := func() {
		print("fade in splash box")
		doc.QuerySelector(".jass-splash-box").Class().SetString("jass-splash-box fade-in fast")
		doc.QuerySelector(".option-bar").Class().SetString("option-bar fade-in one")
		doc.QuerySelector("#option1").Class().SetString("button button-outline option-button fade-in two")
		doc.QuerySelector("#option2").Class().SetString("button button-outline option-button fade-in three")

		// doc.QuerySelector(".jass-splash-image").(*dom.HTMLImageElement).Src = "img/models/model-000l.jpg"
		// doc.QuerySelector(".jass-splash-image").Class().SetString("jass-splash-image fade-in fast")
		// .. wait for the splash image to fade in
	}

	fadeInSplashBox()
	time.Sleep(500 * time.Millisecond)
	doc.QuerySelector(".jass-logo-container").Class().Add("hidden")

	// doc.QuerySelector(".jass-splash-image").Class().Add("green-base")

	// and finally display the top logo
	doc.QuerySelector(".jass-logo-top").Class().Remove("hidden")
	doc.QuerySelector(".hamburger").Class().Remove("hidden")
	doc.QuerySelector(".jass-logo-top").AddEventListener("click", false, func(evt dom.Event) {
		println("clicked on logo")
		w.ScrollTo(0, 0)
		fadeInSplashBox()
		if !doc.QuerySelector(".jass-model-cycle").Class().Contains("hidden") {
			go func() {
				doc.QuerySelector(".jass-model-cycle").Class().SetString("jass-model-cycle fade-out fast")
				time.Sleep(200 * time.Millisecond)
				doc.QuerySelector(".jass-model-cycle").Class().Add("hidden")
			}()
		}
		if !doc.QuerySelector(".jass-sale-items").Class().Contains("hidden") {
			go func() {
				doc.QuerySelector(".jass-sale-items").Class().SetString("jass-sale-items fade-out fast")
				time.Sleep(200 * time.Millisecond)
				doc.QuerySelector(".jass-sale-items").Class().Add("hidden")
			}()
		}
	})

	doc.QuerySelector("#option1").AddEventListener("click", false, func(evt dom.Event) {
		print("clicked on option 1")
		Session.Navigate("/shop")
	})
	doc.QuerySelector("#option2").AddEventListener("click", false, func(evt dom.Event) {
		print("clicked on option 2")
		Session.Navigate("/discover")
	})
}
