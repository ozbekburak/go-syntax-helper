package main

import "fmt"

// go programlama dilinin içerisinde built-in olarak gelen make fonksiyonu ile slice oluşturabiliriz

func main() {

	// burada go dili bize aslında dinamik boyutlandırılmış bir array oluşturup, onu işaret eden bir slice döndürmüş oluyor
	dynamicArray := make([]float64, 5) // 5 : length. len(dynamicArray)=5 !
	display("dynamicArray:", dynamicArray)

	// üçüncü bir argüman ile kapasite belirleyebiliriz
	dynamicArrayWithCapacity := make([]float64, 1, 5)
	display("dynamicArrayWithCapacity: ", dynamicArrayWithCapacity)

	// slices diğer sliceslar da olmak üzere her tipten veriyi içerebilir,
	slice1 := []string{"gregor", "samsa"}
	slice2 := []string{"1", "2"}
	parentSlice := [][]string{
		slice1,
		slice2,
	}
	fmt.Println("Slices of slices: ", parentSlice)
}

func display(nameOfSlice string, sliceItself []float64) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		nameOfSlice, len(sliceItself), cap(sliceItself), sliceItself)
}
