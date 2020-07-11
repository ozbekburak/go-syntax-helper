package main

// reference: https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var count int
	flag.IntVar(&count, "n", 5, "number of lines to read from the file")
	flag.Parse()

	var in io.Reader
	if filename := flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
		defer f.Close()

		in = f
	} else {
		// program herhangi bir argüman verildiğini tespit etmezse,
		// standart inputtan okuma yapar.
		// echo "gregor" | go run 49_flag_positional_arguments.go -n 1
		// gregor çıktısını verecektir. argüman almayan flag, standart inputtan okuma yapar ve onu işler!
		in = os.Stdin
	}

	buf := bufio.NewScanner(in)
	for i := 0; i < count; i++ {
		if !buf.Scan() {
			break
		}
		fmt.Println(buf.Text())
	}

	if err := buf.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading: err", err)
	}
}
