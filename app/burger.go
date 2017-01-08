package main

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func initBurger() {
	w := dom.GetWindow()
	doc := w.Document()

	ldTemplate("slidemenu", "#slidemenu-div", &Session)
	sc := doc.QuerySelector("#slidemenu").Class()
	sc.Remove("cbp-spmenu-open")

	doc.QuerySelector("body").AddEventListener("click", false, func(evt dom.Event) {
		hc := doc.QuerySelector(".hamburger").Class()
		sc := doc.QuerySelector("#slidemenu").Class()
		if sc.Contains("cbp-spmenu-open") {
			sc.Remove("cbp-spmenu-open")
			hc.Remove("is-active")
		}
	})

	doc.QuerySelector(".hamburger").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		c := doc.QuerySelector(".hamburger").Class()
		c.Toggle("is-active")
		if c.Contains("is-active") {
			// print("burger time")
			openBurger()
		} else {
			// print("no more burger")
			closeBurger()
		}
		evt.StopPropagation()
	})

	doc.QuerySelector("#menu-fragrances").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		closeBurger()
		Session.Navigate("/shop")
	})
	// doc.QuerySelector("#menu-discover").AddEventListener("click", false, func(evt dom.Event) {
	// 	evt.PreventDefault()
	// 	closeBurger()
	// 	Session.Navigate("/discover")
	// })
	doc.QuerySelector("#menu-merchandise").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		// w.Open("https://shop.polymer-project.org/list/ladies_outerwear", "merch", "")
		js.Global.Get("location").Set("href", "https://shop.polymer-project.org/list/ladies_outerwear")
	})

	doc.QuerySelector("#menu-about").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		w.Open("https://www.facebook.com/worldofjass", "facebook", "")
	})

	doc.QuerySelector("#menu-ambassadors").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		w.Open("https://www.youtube.com/watch?v=AkZZbcfOJJM&list=PLczWL7gMyRhr7ow79N_YHJiwCV6r9nE5i", "ambassadors", "")
		// w.Open("https://theworldofjass.wordpress.com", "worldofjass", "")
	})

	// doc.QuerySelector("#menu-instagram").AddEventListener("click", false, func(evt dom.Event) {
	// 	evt.PreventDefault()
	// 	w.Open("https://www.instagram.com/worldofjass", "instagram", "")
	// })

	doc.QuerySelector("#menu-blog").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		Session.Navigate("/blog")
		// w.Open("https://theworldofjass.wordpress.com", "blog", "")
	})

	// doc.QuerySelector("#menu-twitter").AddEventListener("click", false, func(evt dom.Event) {
	// 	evt.PreventDefault()
	// 	w.Open("https://twitter.com/worldofjass", "twitter", "")
	// })

	doc.QuerySelector("#menu-contact").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		closeBurger()
		Session.Navigate("/contact")
	})
}

func closeBurger() {
	w := dom.GetWindow()
	doc := w.Document()

	if s := doc.QuerySelector(".cbp-spmenu"); s != nil {
		s.Class().Remove("cbp-spmenu-open")
	}
	doc.QuerySelector(".hamburger").Class().Remove("is-active")
}

func openBurger() {
	w := dom.GetWindow()
	doc := w.Document()

	if s := doc.QuerySelector(".cbp-spmenu"); s != nil {
		s.Class().Add("cbp-spmenu-open")
	}
	doc.QuerySelector(".hamburger").Class().Add("is-active")
}
