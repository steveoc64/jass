package main

import (
	"./shared"
	"github.com/go-humble/router"
)

func mainPage(context *router.Context) {
	go _mainPage("Main", nil, context)
}

func _mainPage(action string, msg *shared.NetData, context *router.Context) {
	println("main page")
}
