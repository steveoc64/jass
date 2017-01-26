package main

import (
	"fmt"
	"strconv"

	"./shared"
	"github.com/go-humble/locstor"
	"github.com/go-humble/router"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func getSessionID() {
	uid, err := locstor.GetItem("uid")
	if err == nil {
		Session.UserID, err = strconv.Atoi(uid)
		print("UID", Session.UserID)
	}
	sid, err := locstor.GetItem("sid")
	if err == nil {
		Session.SID, err = strconv.Atoi(sid)
		print("SID", Session.SID)
	}

	if Session.SID == 0 {
		GetJSON("/api/session", &Session.SID, func() {
			print("Gen SID", Session.SID)
			locstor.SetItem("sid", fmt.Sprintf("%d", Session.SID))
		})
	}
}

type GlobalSessionData struct {
	UserID               int
	Router               *router.Router
	AppFn                map[string]router.Handler
	Context              *router.Context
	ID                   map[string]int
	SID                  int
	URL                  string
	RedrawOnResize       bool
	MobileSensitive      bool
	OrientationSensitive bool
	wasMobile            bool
	LastWidth            int
	Orientation          string
	wasSubmobile         bool
	Products             []shared.Product
	Category             []shared.Category
	Blogs                []shared.Blog
	CartTotal            float64
	CartItemCount        int
	CartItems            []shared.Item
	ScrollFunc           func(*js.Object)
	SelectCat            string
	SelectCatID          int
}

func (s *GlobalSessionData) CategoryOK(cat_id int) bool {
	if s.SelectCatID == 0 {
		return true
	}
	return s.SelectCatID == cat_id
}

func (s *GlobalSessionData) NextBlog(id int) int {
	for i, v := range s.Blogs {
		if v.ID == id {
			if (i + 1) >= len(s.Blogs) {
				return -1 // no next blog
			}
			return s.Blogs[i+1].ID
		}
	}
	return -1
}

func (s *GlobalSessionData) PrevBlog(id int) int {
	for i, v := range s.Blogs {
		if v.ID == id {
			if i == 0 {
				return -1 // no prev blog
			}
			return s.Blogs[i-1].ID
		}
	}
	return -1
}

func (s *GlobalSessionData) GetURL() string {
	return "https://jass.com.au"
}

func (s *GlobalSessionData) GetCartTotal() string {
	// print("cart total", s.CartTotal)
	if s.CartTotal == 0.0 {
		return ""
	}
	return fmt.Sprintf("$ %.0f", s.CartTotal)
}

func (s *GlobalSessionData) GetCartItemCount() string {
	switch s.CartItemCount {
	case 0:
		return ""
	case 1:
		return "1 Item"
	default:
		return fmt.Sprintf("%d Items", s.CartItemCount)
	}
}

func (s *GlobalSessionData) AddToCart(item *shared.Item) {
	if item != nil {
		// check if already in cart, if so, increment qty
		for i, v := range s.CartItems {
			if v.SKU == item.SKU {
				s.CartItemCount++
				s.CartTotal += item.Price
				s.CartItems[i].Qty++
				return
			}
		}

		// not in the cart yet
		s.CartItemCount++
		s.CartTotal += item.Price
		s.CartItems = append(s.CartItems, *item)
	}

}

func (s *GlobalSessionData) GetBlog(id int) *shared.Blog {
	for i, v := range s.Blogs {
		if v.ID == id {
			return &s.Blogs[i]
		}
	}
	return nil
}

var Session GlobalSessionData

func (s *GlobalSessionData) Navigate(url string) {
	w := dom.GetWindow()

	// kill off any scroller function
	if s.ScrollFunc != nil {
		print("remove scroll handler")
		w.RemoveEventListener("scroll", false, s.ScrollFunc)
		s.ScrollFunc = nil
	}

	// print("Navigate to", url)
	// On navigate, clear out any subscriptions on events
	s.Context = nil
	s.Router.Navigate(url)
	s.URL = url
	s.RedrawOnResize = false

	dom.GetWindow().Document().QuerySelector(".jass-footer").Class().Remove("hidden")
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
