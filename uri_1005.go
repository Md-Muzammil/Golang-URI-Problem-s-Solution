package main

import (
	"fmt"
)

func main() {

	var a, b, av float64

	fmt.Scanf("%f %f", &a, &b)

	av = ((a * 3.5) + (b * 7.5)) / 11

	fmt.Printf("MEDIA = %.5f\n", av)
}
