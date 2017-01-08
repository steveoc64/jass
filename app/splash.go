package main

import "honnef.co/go/js/dom"

func doSplashPage() {
	fadeIn("jass-splash-box", "jass-options")
	showButtons("shop", "discover")
}

func showTopMenu() {
	w := dom.GetWindow()
	doc := w.Document()
	doc.QuerySelector(".jass-logo-top").Class().Add("fade-in")
	doc.QuerySelector(".jass-logo-top").Class().Remove("hidden")
	doc.QuerySelector(".loadspinner").Class().Add("fade-out")
	doc.QuerySelector(".hamburger").Class().Add("fade-in")
	doc.QuerySelector(".hamburger").Class().Remove("hidden")
	doc.QuerySelector(".jass-logo-top").AddEventListener("click", false, func(evt dom.Event) {
		print("Clicked on logo")
		Session.Navigate("/")
	})
}
