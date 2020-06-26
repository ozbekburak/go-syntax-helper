package main

import "fmt"

// Writer arayüzü...
// interface kavramı, kullanıldığı tipler tarafından kullanılmasını zorunlu kıldığı methodları içerir diyebiliriz
type Writer interface {
	WriteABook(s string)
}

func main() {
	var isaacAsimov Writer
	isaacAsimov = Writer(Novelist{"Isaac Asimov"})
	robertFrost := Writer(Poet{"Robert Frost"})
	isaacAsimov.WriteABook("I, Robot")
	robertFrost.WriteABook("The Road Not Taken")
}

type Novelist struct {
	name string
}

type Poet struct {
	name string
}

func (novelist Novelist) WriteABook(s string) {
	fmt.Println(novelist.name, "wrote a book which name is:", s)
}

func (poet Poet) WriteABook(s string) {
	fmt.Println(poet.name, "wrote a poem which name is:", s)
}
