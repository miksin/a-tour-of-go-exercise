package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var image [][]uint8 = make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		image[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			image[i][j] = uint8(i*i + j*j)
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
}
