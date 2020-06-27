package main

import "fmt"

type Athlete interface {
	SalaryAmount() string
	Hometown() string
}

type SoccerPlayer struct {
	Name     string
	Age      int
	Salary   float64
	Homeland string
}

type BasketballPlayer struct {
	Name     string
	Age      int
	Salary   float64
	Homeland string
}

func main() {

	player1 := Athlete(SoccerPlayer{
		Name:     "Cristiano Ronaldo",
		Age:      35,
		Salary:   31.000,
		Homeland: " Funchal, Madeira",
	})

	player2 := Athlete(SoccerPlayer{
		Name:     "Alex Morgan",
		Age:      30,
		Salary:   4.000,
		Homeland: "Diamond Bar, California",
	})

	player3 := Athlete(BasketballPlayer{
		Name:     "Lebron James",
		Age:      35,
		Salary:   37.440,
		Homeland: "Akron, Ohio",
	})

	player4 := Athlete(BasketballPlayer{
		Name:     "Diana Taurasi",
		Age:      38,
		Salary:   4.900,
		Homeland: "California",
	})

	fmt.Println("Info of : ", player1, ":", player1.SalaryAmount())
	fmt.Println("Info of : ", player2, ":", player2.SalaryAmount())
	fmt.Println("Info of : ", player3, ":", player3.Hometown())
	fmt.Println("Info of : ", player4, ":", player4.Hometown())

}

func (s SoccerPlayer) SalaryAmount() string {
	if s.Salary > 15.000 {
		return "He gets lots of money!"
	}
	return "It will be better!"

}

func (s SoccerPlayer) Hometown() string {
	return "Not important where he/she is from!"
}

func (b BasketballPlayer) SalaryAmount() string {
	if b.Salary > 15.000 {
		return "He gets lots of money!"
	}
	return "It will be better"

}

func (b BasketballPlayer) Hometown() string {
	return "Not important where he/she is from!"
}
