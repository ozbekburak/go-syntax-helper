package main

import "fmt"

// direkt int kullanabilirken, float şeklinde veri tipi kullanamıyoruz.
// seçeneklerimiz: float32, float64
func divide(x float32, y float32) float32 {
	return x / y
}

func subtract(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(divide(42.5, 13.1))
	fmt.Println(subtract(42, 13))
}
