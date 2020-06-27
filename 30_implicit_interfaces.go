package main

import "fmt"

type Singer interface {
	WriteASong()
}

type AmateurSinger struct {
	Name string
}

// AmateurSinger struct'ının, Singer arayüzüne ait bir methodu kullandığını görüyoruz.
// Böyle durumda açıkca, bu type'ın (AmateurSinger struct'ı yani) bu arayüzün methodunu kullandığını belirtmemize gerek yok
// Diğer dillerde aşina olabileceğiniz, "implements" anahtar kelimesine go dilinde ihtiyaç yok!
func (a AmateurSinger) WriteASong() {
	fmt.Println("This song belongs to a amateur singer!")
}

func main() {
	var someoneAmateur Singer
	someoneAmateur = AmateurSinger{Name: "AmateurGuy"}
	someoneAmateur.WriteASong()
}
