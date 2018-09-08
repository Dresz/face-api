package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)


func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	url := "https://graph.facebook.com/v3.1/685148168507218/feed?"

	payload := strings.NewReader("message=Alerta&access_token=EAADNGF3mxwMBAOtNbGNu9ZC0lgNQpjQkO5nzCvC4CHZCZBXNuUJOSwxp1ZBMFXa1sHMhahnQZAWaY7HHCfJ7BxGv9Ao63otuqdH7v6U7pjIwuOjU8oZAT5EZBa4i55QcF5JZBJetZCYfi90wrH6O7li82Ja8DreWZAJmi3kbUQT541qhdWrCYY6hneIkaZCLUprsxirM7sTZAQF4HwZDZD")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	fmt.Fprint(w, "Post enviado!")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

