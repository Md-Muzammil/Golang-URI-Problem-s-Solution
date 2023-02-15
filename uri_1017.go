package main

import (
	"fmt"
)

func main() {

	var a, b, r float64

	fmt.Scanf("%lf %lf", &a, &b)

	r = (a * b) / 12

	fmt.Printf("%.3lf\n", r)
}
