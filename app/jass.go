package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
	"honnef.co/go/js/dom"
)

var jQuery = jquery.NewJQuery
var codeVersion = "0.1.0"

func main() {
	w := dom.GetWindow()
	doc := w.Document()

	print("Booting up jass -", codeVersion)
	w.ScrollTo(0, 0)

	initRouter()
	initForms()
	initBurger()
	initFooter()
	getSessionID()

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

	showTopMenu()

	doc.QuerySelector("#option-shop").AddEventListener("click", false, func(evt dom.Event) {
		Session.Navigate("/shop")
	})
	doc.QuerySelector("#option-discover").AddEventListener("click", false, func(evt dom.Event) {
		Session.Navigate("/discover")
	})

	// Top row of options - shopping options
	doc.QuerySelector("[name=opt-fragrance]").AddEventListener("click", false, func(evt dom.Event) {
		Session.SelectCat = "Fragrance"
		Session.SelectCatID = 1
		Session.Navigate("/shop")
	})
	doc.QuerySelector("[name=opt-skincare]").AddEventListener("click", false, func(evt dom.Event) {
		Session.SelectCat = "Skincare"
		Session.SelectCatID = 2
		Session.Navigate("/shop")
	})
	doc.QuerySelector("[name=opt-merch]").AddEventListener("click", false, func(evt dom.Event) {
		Session.SelectCat = "Merchandise"
		Session.SelectCatID = 3
		Session.Navigate("/shop")
		// evt.PreventDefault()
		// js.Global.Get("location").Set("href", "https://shop.polymer-project.org/list/ladies_outerwear")
	})

	// Bottom row of options - blog and other links
	doc.QuerySelector("[name=opt-b]").AddEventListener("click", false, func(evt dom.Event) {
		Session.Navigate("/blog")
	})
	doc.QuerySelector("[name=opt-c]").AddEventListener("click", false, func(evt dom.Event) {
		Session.Navigate("/contact")
	})
	doc.QuerySelector("[name=opt-about]").AddEventListener("click", false, func(evt dom.Event) {
		Session.Navigate("/about")
	})
	doc.QuerySelector("[name=opt-a]").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		// w.Open("https://theworldofjass.wordpress.com", "worldofjass", "")
		w.Open("https://www.youtube.com/watch?v=AkZZbcfOJJM&list=PLczWL7gMyRhr7ow79N_YHJiwCV6r9nE5i", "ambassadors", "")
	})

	print("jQuery version -", jQuery().Jquery)

	// Now nav to wherever we should be
	loc := w.Location()
	// print("current loc", loc.Pathname)
	if loc.Pathname != "/" {
		// print("nav to", loc.Pathname)
		Session.Navigate(loc.Pathname)
	}
}

func getDivOffset(el dom.Element) int {
	if el == nil {
		print("getting offset of invalid element")
		return 0
	}
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

	el := doc.QuerySelector(selector)
	if el == nil {
		print("no such element", selector)
	} else {
		// print("load template", tmpl, "into", selector)
		sTemplate := MustGetTemplate(tmpl)
		err := sTemplate.ExecuteEl(el, data)
		if err != nil {
			print("Template Error", err.Error())
		}
	}
}
