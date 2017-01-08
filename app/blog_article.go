package main

import (
	"strconv"

	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

func blogItem(context *router.Context) {
	if len(Session.Blogs) == 0 {
		GetJSON("/api/blog", &Session.Blogs, func() {
			// print("/api/blog complete", Session.Blogs)
			print("/api/blog complete")
			showBlogItem(context)
		})
	} else {
		showBlogItem(context)
	}
}

func showBlogItem(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	Session.RedrawOnResize = true

	if Session.ScrollFunc != nil {
		w.RemoveEventListener("scroll", false, Session.ScrollFunc)
		Session.ScrollFunc = nil
	}

	id, _ := strconv.Atoi(context.Params["id"])
	// print("in blog item", id)
	theBlog := Session.GetBlog(id)
	// print("the blog is", theBlog)

	ldTemplate("jass-blog-article", ".jass-blog-article", theBlog)
	// print("loaded template into jass-blog-article")

	doc.QuerySelector(".jass-blog").Class().Add("hidden")
	w.ScrollTo(0, 0)
	fadeIn("jass-blog-article")
	noButtons()

	doc.QuerySelector(".jass-blog-article").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		t := evt.Target()
		print("clikked on", t.TagName(), t.Class().String())
		switch t.TagName() {
		case "I":
			// print("clicked on icon ... stay here")
		default:
			// print("clicked in general - go back")
			if t.Class().Contains("gotop") {
				print("clikked on gotop")
				el := jQuery("blog-article")
				print("el", el)
				jQuery(".blog-article").Call("scrollTop", 0)
			} else if t.Class().Contains("jass-logo-small") {
				Session.Navigate("/blog")
			}
		}
	})

	if Session.ScrollFunc != nil {
		w.RemoveEventListener("scroll", false, Session.ScrollFunc)
		Session.ScrollFunc = nil
	}

	doc.QuerySelector(".blog-article").AddEventListener("scroll", false, blogArticleScroller)
	articleState = 0

	// Add social buttons
	addSocialButtons(theBlog.GetURL(), theBlog.Name)
}

var articleState = 0
var lastAY = 0
var blogArticleImage = jQuery

func blogArticleScroller(evt dom.Event) {
	w := dom.GetWindow()
	doc := w.Document()

	y := jQuery(".blog-article").ScrollTop()
	theClass := doc.QuerySelector(".blog-article").Class()
	nameClass := doc.QuerySelector(".blog-article-name").Class()
	// print("scroll =", y)

	if y == 0 {
		theClass.Remove("faded")
		theClass.Remove("faded2")
		nameClass.Remove("shrink1")
		nameClass.Remove("shrink2")
	} else if y < 80 {
		if articleState > 0 {
			theClass.Remove("faded")
			theClass.Remove("faded2")
			nameClass.Remove("shrink1")
			nameClass.Remove("shrink2")
		}
		articleState = 0
	} else if y < 240 {
		switch articleState {
		case 0:
			theClass.Add("faded")
			nameClass.Add("shrink1")
			articleState = 1
		case 1:
			if y < lastAY {
				theClass.Remove("faded")
				theClass.Remove("faded2")
				// nameClass.Remove("shrink1")
				// nameClass.Remove("shrink2")
				articleState = 0
			}
		case 2:
			theClass.Remove("faded2")
			// nameClass.Remove("shrink2")
			articleState = 1
		}
	} else {
		switch articleState {
		case 0:
			theClass.Add("faded")
			theClass.Add("faded2")
			nameClass.Add("shrink")
			nameClass.Add("shrink2")
		case 1:
			theClass.Add("faded2")
			nameClass.Add("shrink2")
		case 2:
			if y < lastAY {
				// scrolled backwards
				theClass.Remove("faded2")
			}
			// do nothing
		}
		articleState = 2
	}
	lastAY = y
}
