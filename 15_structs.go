package main

import "fmt"

// struct lar aslında belirttiğimiz veri türlerine sahip veri alanlarının oluşturduğu bir koleksiyondur. (collection)
// başka veri tiplerini kullanarak kendi veri tipimizi oluşturabiliriz. programımıza modülerlik katar.
type Person struct {
	Name    string
	Age     int
	Address string
	Weight  float64
}

func main() {

	myPerson := Person{"Tyler Durden", 25, "İstanbul", 75.4}
	fmt.Println("All information about person:", myPerson)

	// struct alanlarına "." ile erişebiliriz
	fmt.Println("Just show the name: ", myPerson.Name)

	// struct alanlarını yine aynı şekilde güncelleyebiliriz
	myPerson.Age = 35
	fmt.Printf("Updated age: %v\n", myPerson.Age)

	/*
		struct alanlarına pointerlar aracılığı ile de erişebiliriz.
		aşağıda dikkat etmemiz gereken nokta şu:
		"pointToMyPerson" aslında structımızın adresine işaret ediyor.
		ancak, (*pointToMyPerson).Name zahmetli bir notasyon olarak görüldüğü için
		dil bize "pointToMyPerson.name" yazmaya izin veriyor
	*/
	pointToMyPerson := &myPerson
	fmt.Println("I point to your name:", pointToMyPerson.Name)

	// struct literals
	// aşağıdaki gösterimle, birden çok değişkeni aynı anda tanımlama şeklini de hatırlamış olduk
	var (
		person1       = Person{"Kaan", 18, "Ankara", 60.5} // Person tipinde
		person2       = Person{Age: 20}                    // diğer alanların nasıl set edileceğini inceleyin, ayrıca alanların sırasının bir önemi olmadığı da ortada
		person3       = Person{}
		personPointer = &Person{"Kaan", 18, "Ankara", 60.5} // *Person tipinde
	)

	fmt.Println(person1, "--", person2, "--", person3, "--", personPointer)

	fightclub := struct {
		Actors []string
		Year   int
		Rating float64
	}{
		Actors: []string{"Brad Pitt", "Edward Norton"},
		Year:   1999,
		Rating: 9.0,
	}
	fmt.Println("My Lovely anonymous struct: ", fightclub)
}
