package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// artık daha fazla verinin gönderilmeyeceğini söylüyoruz
	// sadece gönderici(sender) kanalı kapatabilir
	// genelde kanalları kapatmaya gerek duymayız, range gibi kullandığımız döngüleri
	// sonlandırmak için gerekir
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	// receiver(alıcı) kanalın daha veri alıp almayacağını test edebilir
	// kanal kapandıysa ve daha fazla veri alınmayacaksa ok = false
	v, ok := <-c
	fmt.Println("\n--\nReceiver: ", v, "\nAre there more value:", ok)
}
