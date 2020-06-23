package main

import "fmt"

// mapler keyleri, value'lara mapler. map'in de slice gibi zero value'su nil'dir.
var galatasaray map[string]string

func main() {
	galatasaray = make(map[string]string)
	galatasaray["goalkeeper"] = "Muslera"
	galatasaray["defender"] = "Lemina"
	galatasaray["striker"] = "Falcao"
	fmt.Println("Goalkeeper of the Galatasaray:", galatasaray["goalkeeper"])
	fmt.Println("Defence player of the Galatasaray:", galatasaray["defender"])
	fmt.Println("Striker of the Galatasaray:", galatasaray["striker"])
}
