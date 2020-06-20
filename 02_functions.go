package main

import "fmt"

// direkt "int" veri tipimiz varken, diğer dillerden alışmış olabileceğimiz "float" yok.
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
