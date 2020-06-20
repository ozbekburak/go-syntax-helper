package main

import "fmt"

// defer: parent (deferı kapsayan) fonksiyon bitinceye kadar çalıştırmak istemediğimiz fonksiyonlar için kullanırız.
// örneğin, aşağıdaki örnekte hello ekrana yazdırıldıktan sonra world yazılacaktır.
func main() {
	defer fmt.Println("world")
	fmt.Println("hello")

}
