package main

import (
	"fmt"
	"strconv"

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

	ww := w.InnerWidth()
	blogCols = 1
	if ww >= 800 {
		blogCols = 2
	}
	if ww >= 1280 {
		blogCols = 3
	}
	print("blog cols =", blogCols)

	// Load up em templates
	ldTemplate("jass-blog", ".jass-blog", &Session)

	JBOffset = (int)(doc.QuerySelector(".header-pad").(*dom.HTMLDivElement).OffsetHeight())
	// print("JBO", JBOffset)

	if LastBlogViewed != -1 {
		highlightItem(LastBlogViewed - 1)
		i := doc.QuerySelector(fmt.Sprintf(`[name="blog-%d"]`, LastBlogViewed)).(*dom.HTMLDivElement)
		divOffset := getDivOffset(i)
		w.ScrollTo(0, divOffset-JBOffset-2)
		i.Focus()
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
					// FocusedBlogElement = id
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

	// Get the height of the first blog element
	blogItemHeight = (int)(doc.QuerySelector("[name=blog-1]").(*dom.HTMLDivElement).OffsetHeight())
	// print("blogitemheight", blogItemHeight)
	blogItemHeight += 16 // 8px margin
	highlightItem(0)
}

func blogScroller(evt dom.Event) {
	w := dom.GetWindow()
	y := w.ScrollY()

	// print("window scroll event", y, blogItemHeight, y/blogItemHeight)
	if blogItemHeight == 0 {
		return
	}
	theItem := ((y + (blogItemHeight / 2)) / blogItemHeight) * blogCols

	if theItem != lastH {
		evt.PreventDefault()
		highlightItem(theItem)
		lastH = theItem
	}
	lastY = y
}

func highlightItem(i int) {
	w := dom.GetWindow()
	doc := w.Document()

	// print("highlightitem", i, Session.Blogs[i].ID)

	el := doc.QuerySelector(fmt.Sprintf("[name=blog-%d]", Session.Blogs[lastH].ID))
	if el != nil {
		el.Class().Remove("highlight")
	}
	el = doc.QuerySelector(fmt.Sprintf("[name=blog-%d]", Session.Blogs[i].ID))
	if el != nil {
		el.Class().Add("highlight")
	}

	FocusedBlogElement = Session.Blogs[i].ID
}

func getBlogs() {
	Session.Blogs = []shared.Blog{}
	GetJSON("/api/blog", &Session.Blogs, func() {
		print("/api/blog complete - getBlogs", Session.Blogs)
	})
}

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
	print("in blog item", id)
	theBlog := Session.GetBlog(id)
	print("the blog is", theBlog)

	// ldTemplate("jass-blog", ".jass-blog", &Session)

	ldTemplate("jass-blog-article", ".jass-blog-article", theBlog)
	print("loaded template into jass-blog-article")

	doc.QuerySelector(".jass-blog").Class().Add("hidden")
	w.ScrollTo(0, 0)
	fadeIn("jass-blog-article")
	noButtons()

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// }()

	// print("looking for .blog-article-name")
	// ela := doc.QuerySelector(".blog-article-name")
	// print("ela =", ela)
	// if ela == nil {
	// 	print("cant find article name")
	// 	blogHeaderEnds = 80
	// } else {
	// 	blogHeaderEnds = getDivEnd(doc.QuerySelector(".blog-article-name"))
	// }
	// navEnds = getDivEnd(doc.QuerySelector(".navigation"))
	// // print("article name ends at", blogHeaderEnds)

	doc.QuerySelector(".jass-blog-article").AddEventListener("click", false, func(evt dom.Event) {
		evt.PreventDefault()
		t := evt.Target()
		switch t.TagName() {
		case "I":
			// print("clicked on icon ... stay here")
		default:
			// print("clicked in general - go back")
			Session.Navigate("/blog")
		}
	})

	if Session.ScrollFunc != nil {
		w.RemoveEventListener("scroll", false, Session.ScrollFunc)
		Session.ScrollFunc = nil
	}

	doc.QuerySelector(".blog-article").AddEventListener("scroll", false, blogArticleScroller)
	// Session.ScrollFunc = w.AddEventListener("scroll", false, blogArticleScroller)
	// blogArticleTitle = doc.QuerySelector(".blog-article-title").(*dom.HTMLDivElement)
	// blogArticleTitleTop = getDivOffset(blogArticleTitle)
	articleState = 0
}

// var blogArticleTitle = &dom.HTMLDivElement{}
// var blogArticleTitleTop = 0
// var blogHeaderEnds = 0
// var navEnds = 0
var articleState = 0
var lastAY = 0

func blogArticleScroller(evt dom.Event) {
	w := dom.GetWindow()
	doc := w.Document()

	y := jQuery(".blog-article").ScrollTop()
	theClass := doc.QuerySelector(".blog-article").Class()
	// print("scroll =", y)

	// print("scroll article", y, articleState)
	if y < 80 {
		if articleState > 0 {
			theClass.Remove("faded")
			theClass.Remove("faded2")
		}
		articleState = 0
	} else if y < 240 {
		switch articleState {
		case 0:
			theClass.Add("faded")
			articleState = 1
		case 1:
			if y < lastAY {
				theClass.Remove("faded")
				theClass.Remove("faded2")
				articleState = 0
			}
		case 2:
			theClass.Remove("faded2")
			articleState = 1
		}
	} else {
		switch articleState {
		case 0:
			theClass.Add("faded")
			theClass.Add("faded2")
		case 1:
			theClass.Add("faded2")
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
