package main

import (
	"fmt"

	"./shared"
	"github.com/go-humble/router"
	"honnef.co/go/js/dom"
)

func blog(context *router.Context) {
	w := dom.GetWindow()
	doc := w.Document()

	// Load up em templates
	sTemplate := MustGetTemplate("jass-blog")
	sTemplate.ExecuteEl(doc.QuerySelector(".jass-blog"), &Session)

	for _, v := range Session.Blogs {
		// set background images on each blog-item
		// print("looking for blog-", v.ID)
		i := doc.QuerySelector(fmt.Sprintf(`[name="blog-image-%d"]`, v.ID))
		if i != nil {
			bgi := fmt.Sprintf("url(/img/models/%s)", v.Image)
			print("got it, set BGI", bgi)
			i.(*dom.HTMLDivElement).Style().SetProperty("background-image", bgi, "")
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
