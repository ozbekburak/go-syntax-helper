// Go, hataları hata değerleriyle ifade eder
package main

import (
	"fmt"
	"time"
)

// CustomError ile kendimiz bir hata structı oluşturduk
type CustomError struct {
	What string
	When time.Time
}

func (customError *CustomError) Error() string {
	return fmt.Sprintf("%s occured at %v", customError.What, customError.When)
}

func run() error {
	return &CustomError{
		"Something bad happened!",
		time.Now(),
	}
}

func main() {
	// dikkat edilmesi gereken noktalardan biri:
	// nil : success, non-nil : hatayı gösterir
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
