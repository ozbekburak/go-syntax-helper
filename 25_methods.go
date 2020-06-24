package main

import (
	"fmt"
	"math"
)

// go, diğer programlama dillerinin aksine sınıflara sahip değildir
// onların yerine methodları kullanırız. method, özel bir argüman alan fonksiyondur diyebiliriz

// Writer sevdiğimiz yazarların doğum ve ölüm tarihlerini tutan bir structtır
type Writer struct {
	name string
	born int
	died int
}

// Age methodu, Writer struct tipinde özel bir argüman alarak, yazarımızın yaş bilgisini döndürür.
/*
	aşağıdaki methodu

	func Age(w Writer) int {
		...
	}

	şeklinde normal bir fonksiyon olarak da tanımlayabilirdik.
	Methodların, alıcı (receiver) argümana sahip birer fonksiyon olduğunu unutmayalım!
*/
func (w Writer) Age() int {
	return w.died - w.born
}

// MyInt gibi method receiverlarını methodlar ile aynı pakette tanımlamalıyız, strcut dışında da
// aşağıdaki gibi methodumuz için receiver tanımlayabiliriz, ancak farklı bir pakette tanımlanmış
// (buna built-in paketleri de (int gibi) dahil) receiverları kullanamayız
type MyInt int

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

	inPackageReceiver := MyInt(math.Pow(3, 2))
	fmt.Println(inPackageReceiver)
}
