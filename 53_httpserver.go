package main

import (
	"fmt"
	"net/http"
)

// handlerlar http.handler interfaceini kullanan nesnelerdir
// aşağıdaki hello fonksiyonumuz handler olarak servis veriyor
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {

}
