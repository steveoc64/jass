package main

import (
	"time"

	"honnef.co/go/js/dom"
)

func initBurger() {
	w := dom.GetWindow()
	doc := w.Document()

	doc.QuerySelector(".hamburger").AddEventListener("click", false, func(evt dom.Event) {
		c := doc.QuerySelector(".hamburger").Class()
		c.Toggle("is-active")
		if c.Contains("is-active") {
			print("burger time")
			openBurger()
		} else {
			print("no more burger")
			closeBurger()
		}
	})
}

func closeBurger() {
	w := dom.GetWindow()
	doc := w.Document()

	sc := doc.QuerySelector("#slidemenu").Class()
	sc.Remove("fade-in")
	sc.Add("fade-out")
	go func() {
		time.Sleep(200 * time.Millisecond)
		sc.Remove("cbp-spmenu-open")
	}()
}

func openBurger() {
	w := dom.GetWindow()
	doc := w.Document()

	sTemplate := MustGetTemplate("slidemenu")
	sTemplate.ExecuteEl(doc.QuerySelector("#slidemenu-div"), &Session)
	sc := doc.QuerySelector("#slidemenu").Class()
	sc.Add("cbp-spmenu-open")
	sc.Remove("fade-out")
	sc.Add("fade-in")
}
