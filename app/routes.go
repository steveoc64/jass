package main

import (
	"errors"

	"./shared"

	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

func fixLinks() {
	Session.Router.InterceptLinks()
}

// Load a template and attach it to the specified element in the doc
func loadTemplate(template string, selector string, data interface{}) error {
	w := dom.GetWindow()
	doc := w.Document()

	t, err := GetTemplate(template)
	if t == nil {
		print("Failed to load template", template)
		return errors.New("Invalid template")
	}
	if err != nil {
		print(err.Error())
		return err
	}

	el := doc.QuerySelector(selector)
	// print("loadtemplate", template, "into", selector, "=", el)
	if el == nil {
		print("Could not find selector", selector)
		return errors.New("Invalid selector")
	}
	// print("looks ok adding template", t, "to", el, "with", data)
	if err := t.ExecuteEl(el, data); err != nil {
		print(err.Error())
		return err
	}
	Session.Router.InterceptLinks()
	return nil
}

func enableRoutes(Rank int) {

	// print("enabling routes for rank", Rank, "session", Session, Session.GetRank())

	Session.AppFn = map[string]router.Handler{
		"mainpage": mainPage,
	}
	w := dom.GetWindow()
	doc := w.Document()

	if el := doc.QuerySelector("#show-image"); el != nil {
		// print("Adding click event for photo view")
		el.AddEventListener("click", false, func(evt dom.Event) {
			el.Class().Remove("md-show")
			// doc.QuerySelector("#show-image").Class().Remove("md-show")
		})
	}
}

func initRouter() {
	print("initRouter")
	Session.Context = nil
	Session.ID = make(map[string]int)

	// Include public routes
	Session.Router = router.New()
	Session.Router.ShouldInterceptLinks = true
	Session.Router.HandleFunc("/", defaultRoute)
	Session.Router.HandleFunc("/fragrance", fragrance)
	Session.Router.HandleFunc("/shop", shop)
	Session.Router.HandleFunc("/merchandise", merchandise)
	Session.Router.HandleFunc("/discover", discover)
	Session.Router.Start()

}

func defaultRoute(context *router.Context) {
	print("Nav to Default Route")
}

func loadRoutes(Rank int, Routes []shared.UserRoute) {

	if Session.Router != nil {
		Session.Router.Stop()
	}
	Session.Router = router.New()
	Session.Router.ShouldInterceptLinks = true
	enableRoutes(Rank)

	for _, v := range Routes {
		// print("processing ", v.Route, v)
		if f, ok := Session.AppFn[v.Func]; ok {
			// print("found a function called", v.Func)
			// print("adding route", v.Route, v.Func)
			Session.Router.HandleFunc(v.Route, f)
		}
	}
	Session.Router.Start()
}
