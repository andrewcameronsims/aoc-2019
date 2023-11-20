package day8

import (
	"fmt"
	"math"
)

type Image struct {
	Layers [][][]int
	Height int
	Width  int
}

func NewImage(encoded []int, height int, width int) Image {
	numLayers := len(encoded) / height / width
	layers := make([][][]int, numLayers)

	for i := 0; i < numLayers; i++ {
		layer := make([][]int, height)

		for j := 0; j < height; j++ {
			row := make([]int, width)

			for k := 0; k < width; k++ {
				row[k] = encoded[(i*height*width)+(j*width)+k]
			}

			layer[j] = row
		}

		layers[i] = layer
	}

	img := Image{
		Height: height,
		Width:  width,
		Layers: layers,
	}

	return img
}

func Count2D(matrix [][]int, target int) int {
	var count int

	for _, row := range matrix {
		for _, cell := range row {
			if cell == target {
				count++
			}
		}
	}

	return count
}

func Solution(input []int) {
	fmt.Printf("Day One: %d\n", partOne(input))
	img := partTwo(input)

	fmt.Println("Day two")
	fmt.Println()

	for _, row := range img {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Print("\n")
	}
}

func partTwo(input []int) [][]string {
	img := NewImage(input, 6, 25)

	rendered := make([][]string, img.Height)

	for i := 0; i < img.Height; i++ {
		row := make([]string, img.Width)

		for j := 0; j < img.Width; j++ {
			for _, l := range img.Layers {
				cell := l[i][j]
				if cell == 0 {
					row[j] = "#"
					break
				}

				if cell == 1 {
					row[j] = " "
					break
				}
			}
		}

		rendered[i] = row
	}

	return rendered
}

func partOne(input []int) int {
	img := NewImage(input, 6, 25)

	leastZeros := math.MaxInt
	var leastZeroLayer [][]int

	for _, l := range img.Layers {
		zeros := Count2D(l, 0)
		if zeros < leastZeros {
			leastZeroLayer = l
			leastZeros = zeros
		}
	}

	ones := Count2D(leastZeroLayer, 1)
	twos := Count2D(leastZeroLayer, 2)

	return ones * twos
}
