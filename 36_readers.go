package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	r := strings.NewReader("Hello, Reader!")
	// fmt.Println(r) // r 'nin adresini yazdırdığına dikkat edelim ve döndürdüğü değerleri inceleyleim
	b := make([]byte, 8)
	// fmt.Println(b)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("\n%q\n", b[0])
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
