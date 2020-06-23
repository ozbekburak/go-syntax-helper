package main

import "fmt"

func main() {
	songDurations := make(map[string]float64)

	// map'e eleman ekleme ya da güncelleme
	songDurations["iwanttobreakfree"] = 3.43
	songDurations["runaway"] = 3.50

	// map'ten bir elemena erişme
	queenSongDuration := songDurations["iwanttobreakfree"]
	fmt.Println(queenSongDuration)

	bonjoviSongDuration := songDurations["runaway"]
	fmt.Println(bonjoviSongDuration)

	// eleman silme
	delete(songDurations, "runaway")
	fmt.Println("Value:", songDurations["runaway"])

	// bir key'in oluşturulup oluşturulmadığını aşağıdaki gibi test edebiliriz
	// key, map'te varsa, ok=true, yoksa ok=false.
	fmt.Println("\nTesting key-value pairs")
	value, ok := songDurations["iwanttobreakfree"]
	fmt.Println("--\nThe queen song duration: ", value, "\nI want to break free?:", ok)

	secondValue, ok := songDurations["runaway"]
	fmt.Println("--\nThe bonjovi song duration: ", secondValue, "\nrunaway?:", ok)

}
