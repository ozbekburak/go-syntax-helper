package main

import "fmt"

// go'da birden fazla result döndürebiliyoruz!
func swap(x, y string) (string, string) {
	return y, x
}

// go'da dönüş değerlerine isim verebiliriz. (named return values)
// argümansız return, named return value'larımızı döndürür. "naked" return olarak bilinir.
// küçük fonksiyonlarda kullanmamız öneriliyor, uzun fonksiyonlarda okunulabilirliği öldürebilir.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b := swap("swap", "me")
	fmt.Println(a, b)
	fmt.Println(split(17))
}
