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

// tüm HTTP istek headerlarını okuyup response body'de görüntüleyen header handlerı
func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
