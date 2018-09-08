package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://graph.facebook.com/v3.1/685148168507218/feed?"

	payload := strings.NewReader("message=Incendio%20En%20Mi%20Casa%20Alerta&access_token=EAADNGF3mxwMBAOtNbGNu9ZC0lgNQpjQkO5nzCvC4CHZCZBXNuUJOSwxp1ZBMFXa1sHMhahnQZAWaY7HHCfJ7BxGv9Ao63otuqdH7v6U7pjIwuOjU8oZAT5EZBa4i55QcF5JZBJetZCYfi90wrH6O7li82Ja8DreWZAJmi3kbUQT541qhdWrCYY6hneIkaZCLUprsxirM7sTZAQF4HwZDZD")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
