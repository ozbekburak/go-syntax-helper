package main

import (
	"fmt"
	"math"
)

// if statementını aşağıdaki gibi de yazabiliriz
// to-do: sıfırdan küçük sayıları kontrol et
func pow(number, numberTwo, powNum float64) (string, float64, float64) {
	// dikkat edilmesi gereken nokta, result değişkeni if statementı dışında kullanılamaz.
	if result := math.Pow(number, powNum); result < math.Pow(numberTwo, powNum) {
		return "numberOne lower than numberTwo", result, numberTwo
	} else {
		// burada result değişkeni kullanılamaz.
		return "numberOne greater than numberTwo:", number, numberTwo
	}

}

func main() {
	fmt.Println(
		pow(1, 4, 2),
	)

}
