package main

import "honnef.co/go/js/dom"

func doSplashPage() {
	w := dom.GetWindow()
	doc := w.Document()

	fadeIn("jass-splash-box")
	doc.QuerySelector(".jass-option-bar").Class().SetString("jass-option-bar fade-in one")
	doc.QuerySelector("#option1").Class().SetString("button button-outline jass-option-button fade-in two")
	doc.QuerySelector("#option2").Class().SetString("button button-outline jass-option-button fade-in three")

}

func showTopMenu() {
	w := dom.GetWindow()
	doc := w.Document()
	doc.QuerySelector(".jass-logo-top").Class().Remove("hidden")
	doc.QuerySelector(".hamburger").Class().Remove("hidden")
	doc.QuerySelector(".jass-logo-top").AddEventListener("click", false, func(evt dom.Event) {
		print("Clicked on logo")
		w.ScrollTo(0, 0)
		doSplashPage()
	})
}

// 	// and finally display the top logo
// 	doc.QuerySelector(".jass-logo-top").Class().Remove("hidden")
// 	doc.QuerySelector(".hamburger").Class().Remove("hidden")
// 	doc.QuerySelector(".jass-logo-top").AddEventListener("click", false, func(evt dom.Event) {
// 		println("clicked on logo")
// 		w.ScrollTo(0, 0)
// 		fadeInSplashBox()
// 		if !doc.QuerySelector(".jass-model-cycle").Class().Contains("hidden") {
// 			go func() {
// 				doc.QuerySelector(".jass-model-cycle").Class().SetString("jass-model-cycle fade-out fast")
// 				time.Sleep(200 * time.Millisecond)
// 				doc.QuerySelector(".jass-model-cycle").Class().Add("hidden")
// 			}()
// 		}
// 		if !doc.QuerySelector(".jass-sale-items").Class().Contains("hidden") {
// 			go func() {
// 				doc.QuerySelector(".jass-sale-items").Class().SetString("jass-sale-items fade-out fast")
// 				time.Sleep(200 * time.Millisecond)
// 				doc.QuerySelector(".jass-sale-items").Class().Add("hidden")
// 			}()
// 		}
// 	})

// 	doc.QuerySelector("#option1").AddEventListener("click", false, func(evt dom.Event) {
// 		print("clicked on option 1")
// 		Session.Navigate("/shop")
// 	})
// 	doc.QuerySelector("#option2").AddEventListener("click", false, func(evt dom.Event) {
// 		print("clicked on option 2")
// 		Session.Navigate("/discover")
// 	})
// }
