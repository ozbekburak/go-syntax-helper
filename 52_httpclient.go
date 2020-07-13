package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("http://www.reddit.com")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println(response.Status)
	fmt.Printf("Response : %v", *response)

	scanner := bufio.NewScanner(response.Body)
	fmt.Println(scanner.Text())
	for i := 0; i < 5; i++ {
		fmt.Println(scanner.Text())
	}
}
