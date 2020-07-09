package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// slicelar için json encoding örneği
	mySlice := []string{"monday", "tuesday", "wednesday"}
	mySliceEncoded, _ := json.Marshal(mySlice)
	fmt.Println(string(mySliceEncoded))

	// mapler için json encoding örneği
	myMap := map[string]string{"book": "the sorrows of young werther", "author": "goethe"}
	myMapEncoded, _ := json.Marshal(myMap)
	fmt.Println(string(myMapEncoded))

	/*
		__aşağıdaki örnekler gobyexample.com/json dan alınmıştır__
	*/

	// encoding/json paketi özel veri tiplerinizi otomatik olarak encode eder (örneğin aşağıdaki response1 custom bir type bizim için)
	// encode edilmiş çıktılara yalnızca export edilmiş alanları dahil eder (yani büyük harfle tanımlanmış)
	// ve varsayılan olarak bu adları JSON anahtarları olarak kullanır

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// struct alanlarında encode edilmiş verinin json keylerini etiketler (tag) ile
	// customize edebiliyoruz. (bkz: __response2 struct__)
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// aşağıda hangi veri tipinin geleceğini bilmediğimiz json tipini decode ediyoruz
	byt := []byte(`{"num":6.13, "strs":["a", "b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	// decoded mapteki değerleri kullanmak için önce uygun tip dönüşümlerini gerçekleştirmemiz gerekir
	// num'ı float64'e dönüşütürülmesi gösterilmiştir

	/*
		TODO: Yeni go sürümlerinde araştır. dönüşüm gerçekleştirmeden veriye eriştim?
	*/
	num := dat["num"].(float64)
	fmt.Println(num)

	// iç içe geçmiş veriler (nested data) birçok dönüşüm gerektirir
	strs := dat["strs"].([]interface{})
	str1 := strs[0]
	fmt.Println(str1)

}
