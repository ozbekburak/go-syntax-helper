package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// 'A' ascii karakterinin akışını (stream) yayan bir reader türü yazmamızı istiyorlar
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 65 // A harfinin ascii karşılığı
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
