package main

import (
	"fmt"
	"runtime"
)

// switch statementı diğer dillere çok benzese de go dilinde bazı küçük farklar var.
// go sadece sağlanan durumu(case) çalıştırır. Bundan dolayı her case'in sonuna break yazmamıza gerek kalmaz.
func main() {
	fmt.Print("Go language runs on: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("I wish it is not Windows! :)", os)
	}
}

// go-tour örneğini aynen kullandım, runtime kütüphanesinin küçük bir kullanımını da eklemiş olmak için.
