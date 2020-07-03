package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	tock := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-tock:
			fmt.Println("tock.")
			return
			// default case'i diğer hiçbir case hazır değilse o zaman çalışır
		default:
			fmt.Println("	.")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
