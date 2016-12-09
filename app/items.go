package main

import (
	"io/ioutil"
	"net/http"
)

func getItems() {
	req, err := http.NewRequest("POST", "/api/items", nil)
	// req.Header.Add("Authorization", "Basic "+basicAuth(u, p))
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("did not get acceptable status code: %v body: %q", resp.Status, string(body))
		return false
	}
}
