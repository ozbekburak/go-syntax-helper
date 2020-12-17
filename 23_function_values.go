package main

import (
	"fmt"
	"math"
)

// a tour of go örneğinin aynısı kullanılmıştır.
// argüman olarak fonksiyon alıyoruz, float64 tipinde veri geri döndürüyoruz.
// argüman olarak aldığımız fonksiyonun dönüş değerini de belirttiğimize dikkat edelim.
func compute(fn func(float64, float64) float64, x, y int) float64 {
	return fn(x, y)
}

func main() {
	hypotenuse := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("Hypotenuse of 5 and 12:", hypotenuse(5, 12))

	fmt.Println("Hypotenuse of 3 and 4 (function -> function): ", compute(hypotenuse(3, 4)))
}
