package main

import "encoding/json"

func main() {
	requestBody, error := json.Marshal(map[string]string{
		"name":    "Burak",
		"country": "Turkey",
	})
}
