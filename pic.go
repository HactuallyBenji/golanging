package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	matrix := make([][]uint8, dy);

	for y := 0; y < dy; y++ {
		matrix[y] = make([]uint8, dx);
		for x := range matrix[y] {
			matrix[y][x] = uint8(x^y)
		}
	}
	
	return matrix
}

func main() {
	pic.Show(Pic)
}
