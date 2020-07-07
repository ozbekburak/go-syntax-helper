package main

// below example taken from https://blog.golang.org/json
import (
	"encoding/json"
	"fmt"
)

// interface{} (boş interface) sıfır methodlu bir interface tanımlar. Go'daki her bir
// tip, en azından bir tane boş method uygular bundan dolayı da boş interface için gerekli
// koşulları sağlar

func main() {
	// aşağıda b değişkeninde saklanan JSON verisini düşünelim
	b := []byte(`{"Name": "Wednesday", "Age":6, "Parents":["Gomez", "Morticia"]}`)

	// veri yapısını bilmeden Unmarshal ile onu boş bir interface'e(interface{}) decode edebiliriz
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("Something bad happened: ", err)
	}
	// bu noktada f, keyleri string olan ve valueları boş interface
	// değerleri olarak depolanan bir map olacaktır
	fmt.Println("decoded: ", f.(map[string]interface{}))

	for key, value := range f.(map[string]interface{}) {
		switch v := value.(type) {
		case string:
			fmt.Println(key, "is string, and its value:", v)
		case float64:
			fmt.Println(key, "is float, and its value:", v)
		case bool:
			fmt.Println(key, "is bool, and its value:", v)
		case []interface{}:
			fmt.Println(key, "is an array, and its value:", v)
			for i, u := range v {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(key, "is of a type that i don't know!")
		}
	}
}

// P.S :

// bool for JSON booleans,
// float64 for JSON numbers,
// string for JSON strings, and
// nil for JSON null.
