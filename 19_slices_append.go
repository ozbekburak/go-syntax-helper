package main

import "fmt"

func main() {
	var appendToSlice []float64
	displaySlice(appendToSlice)

	// nil slice'larda da append kullanabiliriz
	appendToSlice = append(appendToSlice, 5.5)
	displaySlice(appendToSlice)

	// slice ihtiyaç olduğu kadar büyüyebilir
	appendToSlice = append(appendToSlice, 1.5)
	displaySlice(appendToSlice)

	// aynı anda birden fazla eleman ekleyebiliriz
	appendToSlice = append(appendToSlice, 3.0, 4.0, 5.0)
	displaySlice(appendToSlice)
}

func displaySlice(appendedSlice []float64) {
	fmt.Printf("len=%d cap=%d %v\n", len(appendedSlice), cap(appendedSlice), appendedSlice)
}
