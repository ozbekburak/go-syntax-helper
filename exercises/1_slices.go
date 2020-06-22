package main

import "golang.org/x/tour/pic"

// localinizde çalıştırıyorsanız : go get -v "golang.org/x/tour/pic" unutmayın!
func Pic(dx, dy int) [][]uint8 {
	length := dy
	resultPic := make([][]uint8, length)
	for i := 0; i < length; i++ {
		// her eleman için bir slice oluşturuyoruz aslında
		resultPic[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			resultPic[i][j] = uint8((i + j) / 2 * 4)
		}
	}
	return resultPic
}

func main() {
	pic.Show(Pic)
}
