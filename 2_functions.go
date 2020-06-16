package main

import "fmt"

// direkt int kullanabilirken, float şeklinde veri tipi kullanamıyoruz.
// seçeneklerimiz: float32, float64
func add(x float32, y float32) float32 {
	return x / y
}

func main() {
	fmt.Println(add(42.5, 13.1))
}
