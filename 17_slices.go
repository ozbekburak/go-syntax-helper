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

	// slice literals, a tour of go örneğinin aynısı kullanıldı
	a := []int{1, 3, 5, 7, 9}
	fmt.Printf("type of a: %T\n", a)

	b := []bool{true, false, true, false, true, false}
	fmt.Println("b: ", b)
	fmt.Printf("type of a: %T\n", b)

	s := []struct {
		i int
		j bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// aşağıdaki gösterimlerin aynı olduğunu hatırlayalım:
	myArray := [3]float64{1.5, 2.5, 3.5}
	fmt.Println("This is my array:", myArray)
	fmt.Println("Aşağıdaki gösterimler de aynı çıktıyı verecektir\n------------")
	fmt.Println(myArray[0:])
	fmt.Println(myArray[0:3])
	fmt.Println(myArray[:3])
	fmt.Println(myArray[:])

	/*
		slice length and capacity
		length: slice'ın içerdiği eleman sayısı
		capacity: slice'ın, ilk elemanından itibaren, bağlı olduğu array'in alabileceği eleman sayısıdır
		slice'ımızın uzunluğunu (length) tekrar slicelayarak (re-slice) genişletebiliriz
	*/
	population := []float64{40.5, 65.7, 70.0, 75.1, 81.2}
	fmt.Printf("---\nPopulation last 5 years: %v\nlength: %d, capacity %d\n", population, len(population), cap(population))

	// uzunluğu sıfır slice oluşturma
	population = population[:0]
	fmt.Printf("\nNo population: %v\nlength: %d, capacity %d\n", population, len(population), cap(population))

	// uzunluğu artırılmış slice
	population = population[:4]
	fmt.Printf("\nFirst 4 years population: %v\nlength: %d, capacity %d\n", population, len(population), cap(population))

	// İlk iki değeri düşürülen slice, kapasite değerine de dikkat edelim
	population = population[2:]
	fmt.Printf("\nFirst 4 years population: %v\nlength: %d, capacity %d\n", population, len(population), cap(population))
}
