package main

import "fmt"

// aşağıda array gösteriminin iki farklı şekli kullanılmıştır
// for içerisinde len() fonksiyonundan faydalanılarak array'imizin uzunluğuna nasıl erişebileceğimizi de gösterdik
// !!! Array'ler yeniden boyutlandırılamaz, bundan dolayı slice denen bir veri tipinden yardım alacağız ileride
func main() {
	var writers [4]string

	writers[0] = "Ahmet Hamdi Tanpınar"
	writers[1] = "İhsan Oktay Anar"
	writers[2] = "Franz Kafka"
	writers[3] = "Ernest Hemingway"

	for i := 0; i < len(writers); i++ {
		fmt.Printf("%v. writer: %s\n", i, writers[i])
	}

	positiveSingleDigits := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(positiveSingleDigits)
}
