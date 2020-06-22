package main

import (
	"fmt"
	"math"
)

var sqrt = []int{0, 1, 2, 4, 9, 16, 25, 36, 49, 64, 81, 100}

func main() {
	// range, slice veya mapler ile kullanılır. foreach gibi çalışır, veri elemanları üzerinde gezer
	// geriye iki değer döner, birincisi index ikincisi ise o indexteki elemanın kopyasıdır
	// kafa karışıklığı olmaması için açıkça belirttim (index ve value olarak) genel kullanımı "i, v" şeklindedir
	for index, value := range sqrt {
		fmt.Printf("%d. number: %d, its square root:%v\n", index, value, math.Sqrt(float64(value)))
	}

	for i := range sqrt { // for i, _ := range sqrt (aynı, lint kullanıyorsanız bunu kullanmanızı isteyecektir)
		fmt.Println("Value ile ilgilenmiyorum, index:", i)
	}

	for _, value := range sqrt {
		fmt.Println("Index ile ilgilenmiyorum, value:", value)
	}

}
