package main

import (
	"bytes"
	"encoding/json"
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

	defer response.Body.Close()
}
