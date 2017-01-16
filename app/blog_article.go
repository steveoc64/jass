package main

import (
	"fmt"
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

var prevBlog = -1
var nextBlog = -1

func showBlogItem(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	Session.RedrawOnResize = true

	if Session.ScrollFunc != nil {
		w.RemoveEventListener("scroll", false, Session.ScrollFunc)
		Session.ScrollFunc = nil
	}

	id, _ := strconv.Atoi(context.Params["id"])
	theBlog := Session.GetBlog(id)

	ldTemplate("jass-blog-article", ".jass-blog-article", theBlog)
	// print("loaded template into jass-blog-article")

	prevBlog = Session.PrevBlog(id)
	nextBlog = Session.NextBlog(id)
	if prevBlog == -1 {
		doc.QuerySelector(".fa-chevron-left").Class().Add("none-such")
	}
	if nextBlog == -1 {
		doc.QuerySelector(".fa-chevron-right").Class().Add("none-such")
	}

	doc.QuerySelector(".jass-blog").Class().Add("hidden")
	w.ScrollTo(0, 0)

	// banh := jQuery(".blog-article-name").Height()
	// print("banh =", banh)
	fadeIn("jass-blog-article")
	noButtons()

	// set top margin on blog-image
	// jQuery(".blog-article").SetCss("margin-top", banh)

	doc.QuerySelector(".gotop").AddEventListener("click", false, func(dom.Event) {
		jQuery(".blog-article").Call("scrollTop", 0)
	})

	doc.QuerySelector(".blog-article-name").AddEventListener("click", false, func(evt dom.Event) {
		t := evt.Target()
		c := t.Class()
		if t.TagName() == "SPAN" {
			c = t.ParentElement().Class()
		}
		if c.Contains("gotop") && !c.Contains("shrink1") {
			print("bb")
			Session.Navigate("/blog")
		}

	})

	doc.QuerySelector(".blog-article").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		t := evt.Target()
		c := t.Class()
		print("clikked on", t.TagName(), t.Class().String())
		switch t.TagName() {
		case "I":
			// print("clicked on icon ... stay here")
			if c.Contains("fa-chevron-right") {
				// print("goto next blog")
				if nextBlog != -1 {
					Session.Navigate(fmt.Sprintf("/blog/%d", nextBlog))
					return
				}
			} else if c.Contains("fa-chevron-left") {
				if prevBlog != -1 {
					Session.Navigate(fmt.Sprintf("/blog/%d", prevBlog))
					return
				}
			}

		case "DIV":
			// print("clicked in general - go back")
			if c.Contains("jass-logo-small") {
				Session.Navigate("/blog")
				return
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
