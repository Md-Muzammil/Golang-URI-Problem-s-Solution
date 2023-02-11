package main

import (
	"fmt"
)

func main() {
	var R, A float64
	fmt.Scanf("%f", &R)
	A = 3.14159 * (R * R)
	fmt.Printf("A=%.4f\n", A)

}
