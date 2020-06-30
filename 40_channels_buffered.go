package main

import "fmt"

func main() {
	// kanalları ara belleğe alabiliyoruz.
	motivationChannel := make(chan string, 3)
	motivationChannel <- "Don't"
	motivationChannel <- "Give"
	motivationChannel <- "Up"
	// motivationChannel <- "Yes!" : buffer limiti(3) aşıldığı için deadlock hatası verecektir
	fmt.Println(<-motivationChannel)

}
