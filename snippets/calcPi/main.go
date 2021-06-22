package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(calcPi(1000000))
}

func calcPi(terms int) float64 {
	pi := 0.0
	t := float64(terms)
	for i := 0.0; i < t; i++ {
		pi += math.Pow(-1.0, i) * 4.0 / (1.0 + (2.0 * i))
	}
	return pi
}
