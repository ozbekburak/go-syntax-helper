package main

import "fmt"

type Point3D struct {
	X, Y, Z int
}

func (coordinate *Point3D) ScaleMethod(s int) {
	coordinate.X = coordinate.X * s
	coordinate.Y = coordinate.Y * s
	coordinate.Z = coordinate.Z * s
}

func ScaleFunction(coordinate *Point3D, s int) {
	coordinate.X = coordinate.X * s
	coordinate.Y = coordinate.Y * s
	coordinate.Z = coordinate.Z * s
}

func main() {
	coordinate := Point3D{
		1, 0, 0,
	}
	// Aşağıdaki fonksiyon gibi, pointer göndermek zorunda değiliz. (&coordinate) veya (coordinate) olur
	// Go, coordinate.ScaleMethod(5) == (&coordinate).ScaleMethod(5) şeklinde bizim yerimize yorumlar
	coordinate.ScaleMethod(5)

	fmt.Println("Calling scale method:", coordinate)
	// Pointer argümanına sahip fonksiyonlara pointer göndermek zorundayız!
	ScaleFunction(&coordinate, 5)
	fmt.Println("Calling scale function after scale method:", coordinate)
}
