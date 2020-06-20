package main

import (
	"fmt"
)

func main() {
	var x int = 25
	var assignIntToFloat float64

	fmt.Printf("%T\n", x) // print "int"

	assignIntToFloat = float64(x)

	fmt.Printf("%T\n", assignIntToFloat) // print "float64"
}
