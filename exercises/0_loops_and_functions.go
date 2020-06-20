package main

import (
	"fmt"
	"math"
)

// Soru: sqrt() yani karekök alma fonksiyonunun kütüphane olmadan gerçeklenmesi.
// Elde edilen sonucun, matematik kütüphanesi ile karşılaştırılması.
func Sqrt(x float64) float64 {
	// z için başlangıç değer float tipinde 1, isterseniz z := float64(1) şeklinde de gösterebilirsiniz.
	z := 1.0
	for i := 1; i <= 100; i++ {
		fmt.Println("Our z:", z)
		z -= (z*z - x) / (2 * x)
		// karekök almada birçok yöntem kullanabilirsiniz, ben tahminimizin karesinin 1üzeri eksi 10 kadar uzaklıkta
		// olmasının iyi bir tahmin olacağını düşünerek bu döngüyü yazdım.
		if z*z+0.0000000001 > x {
			fmt.Println("We got this!\nIteration order:", i)
			return z
		}
	}

	return z
}

func main() {
	ourGuess := Sqrt(2)
	libraryValue := math.Sqrt(2)
	// kütüphane değerinin daha yakın olacağını biliyoruz, libraryValue - ourGuess her türlü pozitif sonuç dönecektir.
	// ben math.Abs kütüphanesini de denemek için bu şekilde yazdım.
	differenceBetweenGuessAndActual := math.Abs(ourGuess - libraryValue)
	fmt.Println("Closest point that we found:", ourGuess)
	fmt.Println("The math.sqrt() gives us:", libraryValue)
	fmt.Printf("Difference : %.15f\n", differenceBetweenGuessAndActual)
}
