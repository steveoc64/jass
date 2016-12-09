package main

import (
	"./shared"
	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

type GlobalSessionData struct {
	UserID               int
	Router               *router.Router
	AppFn                map[string]router.Handler
	Context              *router.Context
	ID                   map[string]int
	URL                  string
	RedrawOnResize       bool
	MobileSensitive      bool
	OrientationSensitive bool
	wasMobile            bool
	LastWidth            int
	Orientation          string
	wasSubmobile         bool
	Items                []shared.Item
}

var Session GlobalSessionData

func (s *GlobalSessionData) Navigate(url string) {

	// print("Navigate to", url)
	// On navigate, clear out any subscriptions on events
	s.Context = nil
	s.Router.Navigate(url)
	s.URL = url
	s.RedrawOnResize = false
}

func (s *GlobalSessionData) Nav(url string) {
	s.Router.Navigate(url)
	s.URL = url
}

func (s *GlobalSessionData) Back() {
	// On navigate, clear out any subscriptions on events
	s.Context = nil
	s.URL = ""
	s.Router.Back()
}

func (s *GlobalSessionData) Reload(context *router.Context) {
	s.Router.Navigate(s.URL)
	return

	if context == nil {
		// print("reload to ", s.URL)
	} else {
		// print("reload from last url", s.URL)
		s.Router.Navigate(s.URL)
	}
}

func (s *GlobalSessionData) Mobile() bool {
	return dom.GetWindow().InnerWidth() < 740
}

func (s *GlobalSessionData) SubMobile() bool {
	return dom.GetWindow().InnerWidth() < 700
}

func (s *GlobalSessionData) Resize() {
	// print("resize event", dom.GetWindow().InnerWidth(), dom.GetWindow().InnerHeight(), Session.Mobile())
	if s.OrientationSensitive {
		w := dom.GetWindow()

		o := s.Orientation
		s.Orientation = "Landscape"
		if w.InnerHeight() > w.InnerWidth() {
			s.Orientation = "Portrait"
		}
		if s.Orientation != o {
			// print("Redraw due to orientation change")
			// dom.GetWindow().Alert("orientation change")
			s.Reload(s.Context)
			return
		}
	}

	doIt := false
	if s.RedrawOnResize {
		doIt = true
		// print("Redraw due to resize")
	}
	if !doIt && s.MobileSensitive {
		if s.Mobile() != s.wasMobile {
			doIt = true
			// print("Major Redraw due to change from mobile to non-mobile or vise versa")
			// dom.GetWindow().Alert("changed to mobile")
		}
		if s.SubMobile() != s.wasSubmobile {
			doIt = true
			// print("redraw due to change of orientation only, inside mobile mode")
			// dom.GetWindow().Alert("changed to submobile")
		}
	}

	s.wasMobile = s.Mobile()
	s.wasSubmobile = s.SubMobile()
	if doIt {
		s.Reload(s.Context)
	}
}
