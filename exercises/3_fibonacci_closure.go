package main

import "fmt"

// fibonacci fonksiyonu, int tipinde veri döndüren bir fonksiyon döndürür
/*
	Tutorialspoint Fibonacci Algorithm açıklaması :

	Fibonacci series generates the subsequent number by adding two previous numbers.
	Fibonacci series starts from two numbers −
	F0 & F1. The initial values of F0 & F1 can be taken 0, 1 or 1, 1 respectively.

	Ayrıca aşağıdaki koşulu sağlamalıdır
	Fn = Fn-1 + Fn-2

*/
func fibonacci() func() int {
	f0, f1 := 1, 1
	return func() int {
		fNum := f0
		f0 = f1
		f1 = fNum + f1
		fmt.Printf("f: %d, f0: %d, f1: %d\n", fNum, f0, f1)
		return fNum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(i+1, ".step:", f())
	}
}
