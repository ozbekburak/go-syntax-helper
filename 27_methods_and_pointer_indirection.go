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
	coordinate.ScaleMethod(5)
	fmt.Println("Calling scale method:", coordinate) // Aşağıdaki fonksiyon gibi, pointer göndermek zorunda değiliz. (&coordinate) veya (coordinate) olur
	ScaleFunction(&coordinate, 5)                    // Pointer argümanına sahip fonksiyonlara pointer göndermek zorundayız!
	fmt.Println("Calling scale function after scale method:", coordinate)
}
