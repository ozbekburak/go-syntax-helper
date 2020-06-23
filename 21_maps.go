package main

import "fmt"

// mapler keyleri, value'lara mapler. map'in de slice gibi zero value'su nil'dir.
var galatasaray map[string]string

// WriterAndDeathDate represents book writers and their death dates
type WriterAndDeathDate struct {
	Name      string
	DeathDate int
}

func main() {
	galatasaray = make(map[string]string)
	galatasaray["goalkeeper"] = "Muslera"
	galatasaray["defender"] = "Lemina"
	galatasaray["striker"] = "Falcao"
	fmt.Println("Goalkeeper of the Galatasaray:", galatasaray["goalkeeper"])
	fmt.Println("Defence player of the Galatasaray:", galatasaray["defender"])
	fmt.Println("Striker of the Galatasaray:", galatasaray["striker"])

	bestBooks := map[string]WriterAndDeathDate{
		"The Trial": WriterAndDeathDate{
			"Franz Kafka", 1924,
		},
		"Chess Story": WriterAndDeathDate{
			"Stefan Zweig", 1942,
		},
	}

	// sadece struct tipi olarak kullanıyorsak, key:value eşleştirmesi yaparken yeniden WriterAndDeathDate diye belirtmemize gerek kalmaz
	bestBooksDifferentRepresentation := map[string]WriterAndDeathDate{
		"The Trial":   {"Franz Kafka", 1924},
		"Chess Story": {"Stefan Zweig", 1942},
	}

	fmt.Println("Map of the best books and their writer names & death dates :", bestBooks)
	fmt.Println("Another representation same result:", bestBooksDifferentRepresentation)

}
