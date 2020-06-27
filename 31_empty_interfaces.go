package main

import "fmt"

// arayüz (interface) değerlerini (değer, tip) tutan demetler (tuple) gibi düşünebiliriz.
func main() {
	var emptyInterface interface{}
	describe(emptyInterface)

	emptyInterface = 50
	describe(emptyInterface)

	emptyInterface = "not empty"
	describe(emptyInterface)
}

func describe(emptyInterface interface{}) {
	fmt.Printf("(%v, %T)\n", emptyInterface, emptyInterface)
}
