package main

import "fmt"

// go, diğer programlama dillerinin aksine sınıflara sahip değildir
// onların yerine methodları kullanırız. method, özel bir argüman alan fonksiyondur diyebiliriz

// Writer sevdiğimiz yazarların doğum ve ölüm tarihlerini tutan bir structtır
type Writer struct {
	name string
	born int
	died int
}

// Age methodu, Writer struct tipinde özel bir argüman alarak, yazarımızın yaş bilgisini döndürür.
func (w Writer) Age() int {
	return w.died - w.born
}

func main() {
	writer1 := Writer{
		name: "Jules Payot",
		born: 1859,
		died: 1940,
	}
	writer2 := Writer{
		name: "Emily Bronte",
		born: 1818,
		died: 1848,
	}
	fmt.Printf("%s was a French writer. Died at %d\n", writer1.name, writer1.Age())
	fmt.Printf("%s was a English writer. Died at %d\n", writer2.name, writer2.Age())
}
