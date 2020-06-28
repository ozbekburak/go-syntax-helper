package main

import "fmt"

type Animal struct {
	Kind   string
	Mammal bool
}

func (a Animal) String() string {
	return fmt.Sprintf("%v is mammal: %v\n", a.Kind, a.Mammal)
}

func main() {
	a := Animal{"Dog", true}
	b := Animal{"Snake", false}
	// fmt paketi gibi bir çok paket Stringer'ı kullanır
	fmt.Println(a, b)
}
