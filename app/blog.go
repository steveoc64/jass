package main

import (
	"fmt"
	"strconv"
	"time"

	"./shared"

	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

var FocusedBlogElement = 0
var JBOffset = 0
var LastBlogViewed = -1
var blogItemHeight = 0
var lastY = 0
var lastH = 0
var blogCols = 1
var lastBlogScroll = time.Time{}
var scrollThreshold = 100 * time.Millisecond
var headerClass = &dom.TokenList{}
var boundBlogEvents = false

func blog(context *router.Context) {
	Session.Blogs = []shared.Blog{}
	GetJSON("/api/blog", &Session.Blogs, func() {
		// print("/api/blog complete", Session.Blogs)
		print("/api/blog complete")
		showBlog()
	})
}

func showBlog() {
	w := dom.GetWindow()
	doc := w.Document()

	Session.RedrawOnResize = true
	lastBlogScroll = time.Now()

	ww := w.InnerWidth()
	blogCols = 1
	if ww >= 800 {
		blogCols = 2
	}
	if ww >= 1280 {
		blogCols = 3
	}
	// print("blog cols =", blogCols)

	// Load up em templates
	ldTemplate("jass-blog", ".jass-blog", &Session)

	JBOffset = (int)(doc.QuerySelector(".header-pad").(*dom.HTMLDivElement).OffsetHeight())
	// print("JBO", JBOffset)

	header := doc.QuerySelector(".blog-header")
	headerClass = header.Class()
	headerClass.Remove("showme")
	header.AddEventListener("click", false, func(dom.Event) {
		print("clikked on blogheader")
		// doc.QuerySelector(".blog-container").Underlying().Set("scrollTop", 0)
		w.ScrollTo(0, 0)
	})

	if LastBlogViewed != -1 {
		highlightItem(LastBlogViewed - 1)
		i := doc.QuerySelector(fmt.Sprintf(`[name="blog-%d"]`, LastBlogViewed)).(*dom.HTMLDivElement)
		divOffset := getDivOffset(i)
		w.ScrollTo(0, divOffset-JBOffset-2)
		// i.Focus()
		FocusedBlogElement = LastBlogViewed
	}

	for _, v := range Session.Blogs {
		// set background images on each blog-item
		// print("looking for blog-image", v.ID)
		i := doc.QuerySelector(fmt.Sprintf(`[name="blog-image-%d"]`, v.ID)).(*dom.HTMLDivElement)
		if i != nil {
			bgi := fmt.Sprintf("url(/img/models/%s)", v.Image)
			// print("got it, set BGI", bgi)
			i.Style().SetProperty("background-image", bgi, "")

			doc.QuerySelector(fmt.Sprintf("[name=blog-%d]", v.ID)).AddEventListener("click", false, func(evt dom.Event) {
				evt.PreventDefault()
				id, err := strconv.Atoi(evt.Target().GetAttribute("data-id"))
				if err != nil {
					print("not clicked on specific blog thing")
					return
				}
				// w := dom.GetWindow()

				if id == FocusedBlogElement {
					// print("clicked on active element")
					newThing := fmt.Sprintf("/blog/%d", id)
					LastBlogViewed = id
					Session.Navigate(newThing)
				} else {
					theBlog := Session.Blogs[id-1].ID
					// print("clicked on new blog item", id, theBlog)
					highlightItem(theBlog - 1)
					// divOffset := getDivOffset(i)
					// w.ScrollTo(0, divOffset-JBOffset-2)
					// print("scrolls to ", divOffset, divOffset-JBOffset+10)
					// i.Focus()
					FocusedBlogElement = id
				}
			})

			i.AddEventListener("mouseover", false, func(evt dom.Event) {
				id, _ := strconv.Atoi(evt.Target().GetAttribute("data-id"))
				// print("hover over", id, "so turn the rest off")
				for _, vv := range Session.Blogs {
					if vv.ID != id {
						doc.QuerySelector(fmt.Sprintf("[name=blog-%d]", vv.ID)).Class().Remove("highlight")
					}
				}
			})
		} else {
			print("blog not found")
		}

		// and add a clickhandler onto the titlebar
	}

	fadeIn("jass-blog")
	noButtons()

	if Session.ScrollFunc != nil {
		w.RemoveEventListener("scroll", false, Session.ScrollFunc)
		Session.ScrollFunc = nil
	}

	Session.ScrollFunc = w.AddEventListener("scroll", false, blogScroller)
	// doc.QuerySelector(".blog-container").AddEventListener("scroll", false, blogScroller)

	// Links to social btns

	// Get the height of the first blog element
	blogItemHeight = 600
	if el := doc.QuerySelector("[name=blog-1"); el != nil {
		blogItemHeight = (int)(el.(*dom.HTMLDivElement).OffsetHeight())
		// print("calced height to be ", blogItemHeight)
	} else {
		// print("blogitemheight", blogItemHeight)
		blogItemHeight += 16 // 8px margin
	}
	// bih := jQuery("[name=blog-1]").Height()
	// print("bih = ", bih)
	// print("bih2 = ", blogItemHeight)
	highlightItem(0)

	doc.QuerySelector(".jass-logo-small").AddEventListener("click", false, func(evt dom.Event) {
		Session.Navigate("/")
	})
}

// var bt = &dom.HTMLDivElement{}

func blogScroller(evt dom.Event) {
	// bt = evt.Target().(*dom.HTMLDivElement)
	w := dom.GetWindow()
	w.RequestAnimationFrame(doBlogScroll)
	// doBlogScroll(1)
}

func doBlogScroll(t time.Duration) {
	// y := jQuery(".blog-container").ScrollTop()
	y := dom.GetWindow().ScrollY()

	if y == 0 {
		headerClass.Remove("showme")
	} else {
		headerClass.Add("showme")
	}

	// print("window scroll event", y, blogItemHeight, y/blogItemHeight)
	if blogItemHeight == 0 {
		return
	}
	theItem := ((y + (blogItemHeight / 2)) / blogItemHeight) * blogCols

	if theItem != lastH {
		// evt.PreventDefault()
		highlightItem(theItem)
		lastH = theItem
	}
	lastY = y
}

func highlightItem(i int) {
	w := dom.GetWindow()
	doc := w.Document()

	// print("highlightitem", i, Session.Blogs[i].ID)
	if i < 0 || i >= len(Session.Blogs) {
		return
	}

	if lastH >= 0 && lastH < len(Session.Blogs) {
		el := doc.QuerySelector(fmt.Sprintf("[name=blog-%d]", Session.Blogs[lastH].ID))
		if el != nil {
			el.Class().Remove("highlight")
		}
	}
	el := doc.QuerySelector(fmt.Sprintf("[name=blog-%d]", Session.Blogs[i].ID))
	if el != nil {
		el.Class().Add("highlight")
	}

	FocusedBlogElement = Session.Blogs[i].ID
}

// func getBlogs() {
// 	Session.Blogs = []shared.Blog{}
// 	GetJSON("/api/blog", &Session.Blogs, func() {
// 		print("/api/blog complete - getBlogs", Session.Blogs)
// 	})
// }
