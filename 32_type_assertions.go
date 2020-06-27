package main

import "fmt"

func main() {
	var myLovelyInterface interface{} = "#blacklivesmatter"

	// Aşağıdaki statementlar bize, arayüz değeri olan myLovelyInterface'in tuttuğu
	// ayrık typeları ve onlara ait değerleri atadığımız değişkenleri ve çıktıları gösterir
	// arayüzün özel bir tip döndürdüğünü test etmek için type assertion dediğimiz şekilde 2 değer dönülür.
	// birisi ulaşmak istediğimiz değer (underlying value) diğeri mantıksal ifade değeri (boolean value)dir.
	// bize assertion'ın doğru mu yanlış mı olduğunu söyler. (ok: true, ok: yanlış)
	// boolean value'yu kullanmazsak, yanlış tip istediğimzde "panic" ile karşılaşırız
	// panic kısaca bir şeylerin yanlış gittiğini gösteren bult-in bir fonksiyondur
	a := myLovelyInterface.(string)
	fmt.Println(a)

	b, ok := myLovelyInterface.(string)
	fmt.Println(b, ok)

	c, ok := myLovelyInterface.(int)
	fmt.Println(c, ok)

	d := myLovelyInterface.(float64)
	fmt.Println(d) // panic
}
