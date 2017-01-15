package main

import (
	"honnef.co/go/js/dom"
)

func initFooter() {
	w := dom.GetWindow()
	doc := w.Document()

	ldTemplate("footer", ".jass-footer", &Session)

	doc.QuerySelector(".jass-footer").AddEventListener("click", false, func(evt dom.Event) {
		tag := evt.Target().TagName()
		c := evt.Target().Class()
		print("cliked in blog footer", tag)
		switch tag {
		case "I":
			if c.Contains("fa-twitter-square") {
				w.Open("https://twitter.com/worldofjass/", "twitter", "")
			} else if c.Contains("fa-facebook-square") {
				w.Open("https://www.facebook.com/worldofjass/", "facebook", "")
			} else if c.Contains("fa-google-plus-square") {
				w.Open("https://plus.google.com/114254388513381629021", "googleplus", "")
			} else if c.Contains("fa-youtube-square") {
				w.Open("https://www.youtube.com/watch?v=AkZZbcfOJJM&list=PLczWL7gMyRhr7ow79N_YHJiwCV6r9nE5i", "youtube", "")
			} else if c.Contains("fa-instagram") {
				w.Open("https://www.instagram.com/worldofjass/", "instagram", "")
			} else if c.Contains("fa-pinterest-square") {
				w.Open("https://au.pinterest.com/worldofjass/", "pinterest", "")
			} else if c.Contains("fa-chevron-down") {
				evt.CurrentTarget().Class().Remove("clikked")
			}
		case "SPAN", "IMG":
			evt.CurrentTarget().Class().Toggle("clikked")
		case "DIV":
			if c.Contains("backbtn") {
				c.Toggle("clikked")
			}
			href := evt.Target().GetAttribute("data-href")
			if href != "" {
				Session.Navigate(href)
			}
		}
	})

}
