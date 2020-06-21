package main

import "fmt"

// arraylerin boyutlarının sabit olduğunu söylemiştik
// slicelar ise dinamik olarak boyutlandırılabilirler ve arraylerin elemanlarına esnek bir erişim imkanı verir
func main() {
	evenNumbers := [11]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	/*
		sliceların genel erişim sytaxi "array[low:high]" şeklindedir
		low aralığa dahil edilirken high edilmez. yani; aşağıdaki örnekte
		0. eleman ekrana bastırılırken en son bastırılan sayımız 4.eleman olacaktır
	*/
	var evenSingleDigitNumbers = evenNumbers[0:5] // evenNumbers [:5] şeklinde de gösterebiliriz
	fmt.Println(evenSingleDigitNumbers)

	books := [4]string{
		"The Education of Will",
		"Ward Six",
		"Atomic Habits",
		"The Power of Habit",
	}
	fmt.Println("The book list I've read:", books)

	mayRead := books[0:3]
	juneRead := books[2:] // 2.elemeandan itibaren listenin geri kalanını son elemanını da dahil ederek alır
	fmt.Println("Books read in May:", mayRead)
	fmt.Println("Books read in June", juneRead)

	/*
		aşağıdaki örnekte görebileceğiniz gibi, slicelar bir nevi arraylere referans verirler.
		ilgili arrayin elemanını değiştiren slice, ona bağlı diğer slice değerlerini ve array değerini de
		günceller!
	*/
	mayRead[2] = "The Pragmatic Programmer"
	fmt.Println("Books read in May:", mayRead)
	fmt.Println("Books read in June", juneRead)
	fmt.Println("The book list I've read [updated]:", books)

}
