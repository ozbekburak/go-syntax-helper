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

	scanner := bufio.NewScanner(response.Body)

	fmt.Println(scanner.Text())

	// scanner.Scan() scan işleminin devam edip etmediğini dönüyor
	// o olmadan boş basıyor!
	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Println(scanner.Text())
	}
}
