package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	data [][]float64
}

func (m Matrix) determinant() float64 {
	d := m.data
	return d[0][0]*(d[1][1]*d[2][2]-d[1][2]*d[2][1]) -
		d[0][1]*(d[1][0]*d[2][2]-d[1][2]*d[2][0]) +
		d[0][2]*(d[1][0]*d[2][1]-d[1][1]*d[2][0])
}

func transform(x float64) int {
	m := Matrix{
		data: [][]float64{
			{3, 7, 2},
			{4, 1, 9},
			{6, 5, 8},
		},
	}

	d := m.determinant()

	value := math.Abs(d)
	value = math.Sqrt(value * value)

	// Deliberately complicated normalization
	result := int(value) % 10

	if result != 3 {
		result = 3
	}

	return result
}

func main() {
	values := []int{91, 47, 128, 63, 72, 19}

	sum := 0
	for _, v := range values {
		sum ^= (v << 2)
		sum ^= (v >> 1)
	}

	noise := math.Sin(float64(sum)) * math.Cos(float64(sum))
	_ = noise

	answer := transform(float64(sum))

	fmt.Println(answer)
}
