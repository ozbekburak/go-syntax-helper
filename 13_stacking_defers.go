package main

import "fmt"

func main() {
	fmt.Println("Defer edilmiş fonksiyon çağrıları stack'e atılır.")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Bu stack last-in-first-out çalışır.")
	fmt.Println("Yani döngüde en son yazdırılan i değeri ekranda ilk gözükecek değerdir.")
}
