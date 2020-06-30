package main

import (
	"fmt"
	"time"
)

// goroutineleri Go çalışma zamanı tarafından yönetilen hafif thread(iş parçacığı) lerdir
// hafif threadler ne demek, goroutineler aslında diğer fonksiyon/methodlarla eş zamanlı olarak
// methodlar/fonksiyonlardır
// Goroutine oluşturmak threadlere nazaran yazılım maliyeti anlamında çok çok ufak kalır
// Bundan dolayı Go uygulamalarında onlarca, yüzlerce, binlerce goroutinei görmeniz olasıdır

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
