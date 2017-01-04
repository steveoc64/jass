package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
	"honnef.co/go/js/dom"
)

var jQuery = jquery.NewJQuery

func main() {
	w := dom.GetWindow()
	doc := w.Document()

	w.ScrollTo(0, 0)

	initRouter()
	initForms()
	initBurger()

	Session.LastWidth = dom.GetWindow().InnerWidth()
	Session.Orientation = "Landscape"
	if dom.GetWindow().InnerHeight() > dom.GetWindow().InnerWidth() {
		Session.Orientation = "Portrait"
	}
	if Session.Mobile() {
		Session.wasMobile = true
	}
	if Session.SubMobile() {
		Session.wasSubmobile = true
	}

	js.Global.Set("resize", func() {
		Session.Resize()
	})

	// getItems()
	// getBlogs()
	doSplashPage()
	showTopMenu()

	doc.QuerySelector("#option-shop").AddEventListener("click", false, func(evt dom.Event) {
		Session.Navigate("/shop")
	})
	doc.QuerySelector("#option-discover").AddEventListener("click", false, func(evt dom.Event) {
		Session.Navigate("/discover")
	})

	// Top row of options - shopping options
	doc.QuerySelector("[name=opt-shop]").AddEventListener("click", false, func(evt dom.Event) {
		Session.Navigate("/shop")
	})
	doc.QuerySelector("[name=opt-merch]").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		js.Global.Get("location").Set("href", "https://shop.polymer-project.org/list/ladies_outerwear")
	})

	// Bottom row of options - blog and other links
	doc.QuerySelector("[name=opt-b]").AddEventListener("click", false, func(evt dom.Event) {
		Session.Navigate("/blog")
	})
	doc.QuerySelector("[name=opt-about]").AddEventListener("click", false, func(evt dom.Event) {
		Session.Navigate("/about")
	})
	doc.QuerySelector("[name=opt-a]").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		// w.Open("https://theworldofjass.wordpress.com", "worldofjass", "")
		w.Open("https://www.youtube.com/watch?v=AkZZbcfOJJM&list=PLczWL7gMyRhr7ow79N_YHJiwCV6r9nE5i", "ambassadors", "")
	})

	// print("Your current jQuery version is: " + jQuery().Jquery)
}

func getDivOffset(el dom.Element) int {
	retval := float64(0.0)
	pel := el.(dom.HTMLElement).OffsetParent()
	if pel != nil {
		for {
			retval += el.(dom.HTMLElement).OffsetTop()
			el = el.(dom.HTMLElement).OffsetParent()
			if el == nil {
				return int(retval)
			}
		}
	}
	return int(retval)
}

func getDivEnd(el dom.Element) int {
	retval := getDivOffset(el)
	retval += int(el.(dom.HTMLElement).OffsetHeight())
	return retval
}

func ldTemplate(tmpl string, selector string, data interface{}) {
	w := dom.GetWindow()
	doc := w.Document()

	sTemplate := MustGetTemplate(tmpl)
	sTemplate.ExecuteEl(doc.QuerySelector(selector), data)
}
