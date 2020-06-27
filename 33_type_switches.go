package main

import "fmt"

func typeSwitches(i interface{}) {
	switch whichType := i.(type) {
	case float64:
		fmt.Println("Float type:", whichType)
	case bool:
		fmt.Println("Bool type:", whichType)
	default:
		fmt.Println("Unknown type:", whichType)
	}
}

func main() {
	typeSwitches(3.14)
	typeSwitches(true)
	typeSwitches("whoami")
}
