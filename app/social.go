package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"

	"honnef.co/go/js/dom"
)

func addSocialButtons(url, name string) {
	w := dom.GetWindow()
	doc := w.Document()

	// print("adding social buttons")

	// Twitworld
	doc.QuerySelector(".share-twitter").AddEventListener("click", false, func(evt dom.Event) {
		print("clicked on the twitworld thing")
		w.Open(fmt.Sprintf(`%s?text=%s %s`,
			"https://twitter.com/intent/tweet",
			url,
			name),
			"twitter",
			"menubar=0,resizable=1,width=400,height=280")
	})

	// Faceworld
	doc.QuerySelector(".share-facebook").AddEventListener("click", false, func(evt dom.Event) {
		print("clicked on the faceworld thing")
		FB := js.Global.Get("FB")
		print("fb", FB)
		jQuery("loginbutton,#feedbutton").RemoveAttr("disabled")
		FB.Call("getLoginStatus", func(r interface{}) {
			print("in update status callback", r)
		})
		FB.Call("ui", js.M{
			"method": "share",
			"href":   url,
			"quote":  fmt.Sprintf("Chamelee Blog - %s\n%s", name, url),
		}, func(r interface{}) {
			print("completed UI call", r)
		})
	})

	// GWorld
	doc.QuerySelector(".share-google").AddEventListener("click", false, func(evt dom.Event) {
		print("clicked on the googleworld thing")
		w.Open("https://plus.google.com/share?url="+url,
			"gplus",
			"menubar=no,toolbar=no,resizable=yes,scrollbar=yes,width=600,height=480")
		// gapi := js.Global.Get("gapi")
		// print("google api is", gapi)
		// gapi.Get("plusone").Call("go")
	})
}
