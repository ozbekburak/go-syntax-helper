package main

// reference: https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go
import (
	"flag"
	"fmt"
)

// Color type for colorizing our output
type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func main() {
	// flag değerini yakalamak için değişken tanımlıyoruz
	// flag için ayırdığımız değişken bize flagden dönen değerin adresini tutar
	// *useColor -> true, false değeri tutar yani.
	useColor := flag.Bool("color", false, "display colorized output need additional color argument \ne.g: go run flag.go -color red")

	flag.Parse()

	if *useColor {
		switch flag.Arg(0) {
		case "black":
			colorize(ColorBlack, "Hello, DigitalOcean!")
		case "red":
			colorize(ColorRed, "Hello, DigitalOcean!")
		case "green":
			colorize(ColorGreen, "Hello, DigitalOcean!")
		case "yellow":
			colorize(ColorYellow, "Hello, DigitalOcean!")
		case "blue":
			colorize(ColorBlue, "Hello, DigitalOcean!")
		default:
			fmt.Println("Hello, DigitalOcean!")
		}

		return
	}

	fmt.Println("Hello, DigitalOcean!")
}
