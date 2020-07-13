package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("https://dev.to")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println(response.Status)
	fmt.Println(*response.Request)
	//fmt.Printf("Response : %v", *response)

	scanner := bufio.NewScanner(response.Body)
	fmt.Println(scanner.Text())
	// fmt.Println("\n\nScanner: ", scanner.Text())
	// fmt.Println(scanner.Text())
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(scanner.Text())
	// }
}
