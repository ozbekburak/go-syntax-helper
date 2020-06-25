package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale bir pointer receiver ile tanımlanmış bir methoddur
// pointer receiverları, değerleri değiştirebilirler
// Vertex'in önündeki * kaldırıp çıktıyı incelediğimizde, struct'ımızın ilk değerlerini koruduğunu görürürz
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
