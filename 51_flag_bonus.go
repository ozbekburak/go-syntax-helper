package main

import (
	"flag"
	"fmt"
)

// this example taken from gobyexample.com/command-line-flags

func main() {

	// basit flag tanımlamaları string, integer ve boolean veri tipleri için geçerlidir
	// flagler pointer döner değer değil
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("num", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// önceden tanımlanmış değişkenleri de kullanabiliriz flag atamalarımızı yaparken
	// ancak, pointer geçmemiz gerektiğine dikkat edelim
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word: ", *wordPtr)
	fmt.Println("numb: ", *numbPtr)
	fmt.Println("fork: ", *boolPtr)
	fmt.Println("svar: ", svar)
	fmt.Println("tail: ", flag.Args())

}
