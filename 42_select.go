package main

import (
	"fmt"
	"time"
)

// gobyexample.com örneği kullanılmıştır
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "channel one"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "channel two"
	}()

	for i := 0; i < 2; i++ {
		// go dilinde select, birden çok kanal operasyonunda (channel operation) beklememizi sağlar
		// uygulamanın tamamlanması için iki değeri de eşzamanlı olarak bekliyoruz, tamamlanan değer ekrana yazdırılıyor
		// bizim caseimizde 1 saniye sleep koyduğumuz channel2 nin gözüktükten 1 saniye sonra channel1 in gözükmesini bekliyoruz
		select {
		case messageOne := <-channel1:
			fmt.Println(messageOne, "received")

		case messageTwo := <-channel2:
			fmt.Println(messageTwo, "received")
		}
	}
}
