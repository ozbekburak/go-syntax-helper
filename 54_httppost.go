package main

// the examples are shown below mostly get it from https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	requestBody, error := json.Marshal(map[string]string{
		"name":    "Burak",
		"country": "Turkey",
	})
	if error != nil {
		log.Fatalln(error)
	}

	response, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
