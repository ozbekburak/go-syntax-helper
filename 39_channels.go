package main

import "fmt"

// kanallar, kabaca eşzamanlı goroutinleri bağlayan yapılardır (borulardır: gobyexample tanımı)
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// kanallarda veri, ok yönünde ilerler, yani sum 'dan c'ye
	c <- sum // toplamı channela gönderiyoruz
}

func main() {
	// aşağıdaki kod, işi iki goroutine dağıtır. (slicedaki sayıların toplanması)
	// her iki goroutine de çalışmasını tamamladığında toplam sonuca ulaşırız
	s := []int{1, 2, 3, 4, 5}

	myChannel := make(chan int)
	go sum(s[:len(s)/2], myChannel)
	go sum(s[len(s)/2:], myChannel)
	x, y := <-myChannel, <-myChannel // myChanneldan 'den alır
	fmt.Printf("First job: %d\nSecond job: %d\nTotal: %d\n", x, y, x+y)

}
