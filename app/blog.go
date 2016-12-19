package main

import (
	"fmt"
	"strconv"

	"./shared"
	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

var FocusedBlogElement = 0

func blog(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	// Load up em templates
	sTemplate := MustGetTemplate("jass-blog")
	sTemplate.ExecuteEl(doc.QuerySelector(".jass-blog"), &Session)

	for _, v := range Session.Blogs {
		// set background images on each blog-item
		// print("looking for blog-", v.ID)
		i := doc.QuerySelector(fmt.Sprintf(`[name="blog-image-%d"]`, v.ID)).(*dom.HTMLDivElement)
		if i != nil {
			bgi := fmt.Sprintf("url(/img/models/%s)", v.Image)
			print("got it, set BGI", bgi)
			i.Style().SetProperty("background-image", bgi, "")

			// add a clickhandler to the blog image
			i.AddEventListener("click", false, func(evt dom.Event) {
				id, _ := strconv.Atoi(evt.Target().GetAttribute("data-id"))
				print("clicked on blog item", id)

				if id == FocusedBlogElement {
					print("clicked on active element")
					FocusedBlogElement = 0
					doc.QuerySelector(".jass-logo-small-box").(*dom.HTMLDivElement).Focus()
				} else {
					print("clicked on new blog item")
					i.Focus()
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
}

func getBlogs() {
	Session.Blogs = []shared.Blog{}
	GETJson("/api/blog", &Session.Blogs)

	print("blogs is", Session.Blogs)
}

func blogScrollHandler() {
	print("adding scroll handler to the blog")

	w := dom.GetWindow()
	doc := w.Document()
	jb := doc.QuerySelector(".jass-blog")
	jb.AddEventListener("scroll", false, func(evt dom.Event) {
		print("scroll on blog")
	})

}
