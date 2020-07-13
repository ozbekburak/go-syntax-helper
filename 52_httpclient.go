package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("https://devurls.com")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Printf("Type of response : %T", response.Body)
}
