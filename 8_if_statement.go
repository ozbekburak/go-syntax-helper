package main

import "fmt"

func isEven(num int) bool {
	even := true
	// if statementlarında parantez kullanabiliriz ancak formatter kullanıyorsanız kaldıracaktır
	if num%2 == 0 {
		return even
	}
	return !even
}

func main() {
	fmt.Println("Is 10 even number?", isEven(10))
}
