package main

import "fmt"

// arraylerin boyutlarının sabit olduğunu söylemiştik
// slicelar ise dinamik olarak boyutlandırılabilirler ve arraylerin elemanlarına esnek bir erişim imkanı verir
func main() {
	evenNumbers := [11]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	// sliceların genel erişim sytaxi "array[low:high]" şeklindedir
	// low aralığa dahil edilirken high edilmez. yani; aşağıdaki örnekte
	// 0. eleman ekrana bastırılırken en son bastırılan sayımız 4.eleman olacaktır
	var evenSingleDigitNumbers = evenNumbers[0:5] // evenNumbers [:5] şeklinde de gösterebiliriz
	fmt.Println(evenSingleDigitNumbers)

}
