package main

import (
	"./shared"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getItems() bool {
	print("getting items")
	req, err := http.NewRequest("GET", "/api/items", nil)
	// req.Header.Add("Authorization", "Basic "+basicAuth(u, p))
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("did not get acceptable status code: %v body: %q", resp.Status, string(body))
		return false
	} else {
		print("yay - got items")
	}
	Session.Items = []shared.Item{}
	err = json.NewDecoder(resp.Body).Decode(&Session.Items)
	print("items appears to be", Session.Items)
	if err != nil {
		return false
	}

	return true
}
