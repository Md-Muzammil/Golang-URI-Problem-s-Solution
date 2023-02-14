package main

import (
	"fmt"
)

func main() {

	var R int
	var volume, pi float64
	pi = 3.14159
	fmt.Scanf("%d", &R)
	volume = (4 / 3.0 * pi * R * R * R)
	fmt.Printf("VOLUME = %.3lf\n", volume)
}
