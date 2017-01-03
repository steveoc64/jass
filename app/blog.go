package main

import (
	"fmt"
	"strconv"
	"time"

	"./shared"

	"github.com/go-humble/router"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

var FocusedBlogElement = 0
var JBOffset = 0
var LastBlogViewed = 0
var blogItemHeight = 0
var lastY = 0
var lastH = 0

func blog(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	// Load up em templates
	sTemplate := MustGetTemplate("jass-blog")
	sTemplate.ExecuteEl(doc.QuerySelector(".jass-blog"), &Session)

	JBOffset = (int)(doc.QuerySelector(".header-pad").(*dom.HTMLDivElement).OffsetHeight())
	// print("JBO", JBOffset)

	if LastBlogViewed != 0 {
		// i := doc.QuerySelector(fmt.Sprintf(`[name="blog-%d"]`, LastBlogViewed)).(*dom.HTMLDivElement)
		// divOffset := getDivOffset(i)
		// w.ScrollTo(0, divOffset-JBOffset-2)
		// i.Focus()
		FocusedBlogElement = LastBlogViewed
	}

	for _, v := range Session.Blogs {
		// set background images on each blog-item
		// print("looking for blog-", v.ID)
		i := doc.QuerySelector(fmt.Sprintf(`[name="blog-image-%d"]`, v.ID)).(*dom.HTMLDivElement)
		if i != nil {
			bgi := fmt.Sprintf("url(/img/models/%s)", v.Image)
			// print("got it, set BGI", bgi)
			i.Style().SetProperty("background-image", bgi, "")

			i.AddEventListener("click", false, func(evt dom.Event) {
				id, _ := strconv.Atoi(evt.Target().GetAttribute("data-id"))
				// w := dom.GetWindow()

				divOffset := getDivOffset(i)
				if id == FocusedBlogElement {
					print("clicked on active element")
					newThing := fmt.Sprintf("/blog/%d", id)
					LastBlogViewed = id
					Session.Navigate(newThing)
				} else {
					print("clicked on new blog item")
					w.ScrollTo(0, divOffset-JBOffset-2)
					// i.Focus()
					FocusedBlogElement = id
				}
			})
		} else {
			print("not found")
		}

		// and add a clickhandler onto the titlebar
	}

	fadeIn("jass-blog")
	noButtons()

	if scrollFunc != nil {
		w.RemoveEventListener("scroll", false, scrollFunc)
	}
	scrollFunc = w.AddEventListener("scroll", false, blogScroller)

	// Get the height of the first blog element
	blogItemHeight = (int)(doc.QuerySelector("[name=blog-1]").(*dom.HTMLDivElement).OffsetHeight())
	highlightItem(1)
}

var scrollFunc func(*js.Object)

func blogScroller(evt dom.Event) {
	w := dom.GetWindow()
	y := w.ScrollY()

	// print("window scroll event", y, blogItemHeight, y/blogItemHeight)
	if blogItemHeight == 0 {
		return
	}
	theItem := (y / blogItemHeight) + 1

	highlightItem(theItem)
	lastY = y
}

func highlightItem(i int) {
	w := dom.GetWindow()
	doc := w.Document()

	if i != lastH {
		el := doc.QuerySelector(fmt.Sprintf("[name=blog-%d]", lastH))
		if el != nil {
			el.Class().Remove("blog-highlight")
		}
		el = doc.QuerySelector(fmt.Sprintf("[name=blog-image-%d]", lastH))
		if el != nil {
			el.Class().Remove("highlight")
		}
		el = doc.QuerySelector(fmt.Sprintf("[name=blog-title-%d]", lastH))
		if el != nil {
			el.Class().Remove("highlight")
		}
		print("highlight item", i)
		elname := fmt.Sprintf("[name=blog-%d]", i)
		el = doc.QuerySelector(elname)
		el.Class().Add("blog-highlight")
		elname = fmt.Sprintf("[name=blog-image-%d]", i)
		el = doc.QuerySelector(elname)
		el.Class().Add("highlight")
		elname = fmt.Sprintf("[name=blog-title-%d]", i)
		el = doc.QuerySelector(elname)
		el.Class().Add("highlight")
		lastH = i
	}
}

func getBlogs() {
	print("blogs starts as", Session.Blogs)
	Session.Blogs = []shared.Blog{}
	GetJSON("/api/blog", &Session.Blogs, func() {
		print("/api/blog complete", Session.Blogs)
	})
}

func blogItem(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	id, _ := strconv.Atoi(context.Params["id"])
	// print("in blog item", id)
	theBlog := Session.GetBlog(id)
	print("the blog is", theBlog)

	sTemplate := MustGetTemplate("jass-blog-article")
	print("the template is ", sTemplate)
	el := doc.QuerySelector(".jass-blog-article")
	print("el = ", el)
	sTemplate.ExecuteEl(doc.QuerySelector(".jass-blog-article"), theBlog)

	doc.QuerySelector(".jass-blog").Class().Add("hidden")
	w.ScrollTo(0, 0)

	fadeIn("jass-blog-article")
	noButtons()

	go func() {
		time.Sleep(2 * time.Second)
	}()

	doc.QuerySelector(".jass-blog-article").AddEventListener("click", false, func(evt dom.Event) {
		Session.Navigate("/blog")
	})
}
