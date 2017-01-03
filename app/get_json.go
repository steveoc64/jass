package main

import (
	"encoding/json"

	"github.com/gopherjs/jquery"
)

func GetJSON(url string, data interface{}, f func()) {
	jquery.GetJSON(url, func(d interface{}) {
		j, _ := json.Marshal(d)
		json.Unmarshal(j, data)
		go f()
	})
}
