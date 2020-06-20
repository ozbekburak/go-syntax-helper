package main

import "fmt"

// değişkenlerin paket veya fonksiyon seviyesinde olabildiğini görüyoruz.
// fonksiyonların dışında tanımlanan bütün ifadeler (statement) keyword ile başlamalıdır. (var, func etc) ":=" kullanılamaz.

var linux, windows bool      // without initializer
var year, age int = 1995, 25 // with initializer
var name = "Burak"           // initializer var ise, değişken tipi otomatik kendisi alır.

// değişken tanımları bloklar halinde yapılabilir.
var (
	isRaining bool   = false
	movieName string = "Who Am I"
)

func main() {
	var apple int
	city := "İstanbul" // short assignment.

	fmt.Println(apple, linux, windows)
	fmt.Println(name, year, age)
	fmt.Println(city)
}
