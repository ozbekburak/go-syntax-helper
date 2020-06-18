package main

import "fmt"

func main() {
	sum := 0

	// go'da sadece for döngü yapısı var.
	// diğer dillerden alışık olduğumuz while benzeri döngüler yok.
	for i := 0; i < 3; i++ {
		sum += i

	}
	fmt.Println("Sum:", sum)

	// başlangıç ve bitiş değeri vermeden aşağıdaki şekilde kullanabiliriz.
	// while gibi davrandığını söyleyebiliriz.
	decrease := 3
	for decrease > 0 {
		decrease = decrease - 1
		fmt.Println("Decreasing...", decrease)
	}

	// for loop'ta parantez() kullanmadığımıza dikkat edin.

}
