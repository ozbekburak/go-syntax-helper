package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// slicelar için json encoding örneği
	mySlice := []string{"monday", "tuesday", "wednesday"}
	mySliceEncoded, _ := json.Marshal(mySlice)
	fmt.Println(string(mySliceEncoded))

	// mapler için json encoding örneği
	myMap := map[string]string{"book": "the sorrows of young werther", "author": "goethe"}
	myMapEncoded, _ := json.Marshal(myMap)
	fmt.Println(string(myMapEncoded))
}
