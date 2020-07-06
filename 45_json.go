package main

import (
	"encoding/json"
	"fmt"
)

// News structı haberleri temsil eden veri yapımızı gösteriyor
type News struct {
	ID     int
	Title  string
	Body   string
	Author string
	Time   int64
}

func main() {
	fmt.Println("Encoding and Decoding JSON in Go\n-----------")

	// News'e ait bir instance oluşturuyoruz
	n := News{1, "Covid-19 Report", "Last reports show that Covid-19 still threats our health", "Reporter X", 129503943940}

	// encode işlemi
	resultOfEncodedData, err := encodeToJSON(n)
	if err != nil {
		fmt.Println("Someting went wrong: ", err)
	}
	fmt.Printf("Your encoded json: %v\n", resultOfEncodedData)

	// decode işlemi
	resultOfDecodedData, error := decodeJSON(resultOfEncodedData)
	if error != nil {
		fmt.Println("Something went wrong: ", error)
	}
	fmt.Printf("-----------\nHey you successfully decode your json!: %v\n", *resultOfDecodedData)
}

// Aşağıdaki fonksiyon ile verimizin json-encoded halini oluşturuyoruz
// Json, genelde api lar ile çalışırken websitelerinin bize sunduğu veri formatı olarak karşımıza çıkar
func encodeToJSON(n News) ([]byte, error) {
	encodedNews, err := json.Marshal(n)
	if err != nil {
		return nil, err
	}

	return encodedNews, nil

}

// does not work for now
func decodeJSON(data []byte) (*News, error) {
	// JSON datamızı şimdi geri decode edelim
	// İlk önce decode edilecek verinin saklanacağı bir alan oluşturmamız lazım
	var decodedNews News

	// Eğer verimiz (data []byte) decodedNews e uygun bir JSON verisiyle, error
	// nil'e eşit olacaktır ve data'dan gelen veri decodedNews struct'ında saklanacaktır.
	err := json.Unmarshal(data, &decodedNews)
	if err != nil {
		return nil, err
	}
	return &decodedNews, nil
}
