package main

import (
	"fmt"
)

func main() {
	var age int

	// kullanıcıdan input okumanın birçok farklı yolu var.
	// C diline daha yakın bulduğum Scanf'i kullandım, siz farklı yöntemleri de araştırabilirsiniz.
	fmt.Print("\nHi! what is your age -> ")
	fmt.Scanf("%d", &age)

	// koşul kullanmadan switch 'ten faydalanmak bizi uzun if-else bloklarından kurtarabilir.
	switch {
	case age <= 25:
		fmt.Println("You are young, You will miss these days!")
	case age > 25 && age <= 50:
		fmt.Println("Reponsibilities loading... or loaded already?")
	default:
		fmt.Println("You've come a long way.")
	}

}
