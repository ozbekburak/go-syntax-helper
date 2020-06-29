package main

import (
	"fmt"
	"image"
)

func main() {
	// Sırasıyla inceleyelim
	// image paketimiz
	// Rect(x0, y0, x1, y1 int) fonksiyonu aslında Rectangle{Pt(x0, y0), Pt(x1, y1)} basitleştirilmiş hali
	// image.Rect(0, 0, 100, 100) ile Verilen koordinatlarda Rectangle döndürüyoruz
	// bu Rectangle'ı NewRGBA 'ye gönderiyoruz.
	// https://golang.org/pkg/image/#NewRGBA buradan incelersek Rectangle'ı argüman olarak kabul ediyor
	// geriye *image.RGBA dönüyor
	// Aşağıda da bounds, at, colormodel ile Image interfaceinin kullandığı methodların çıktısını görebiliriz
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
	fmt.Println(m.ColorModel)
}
