package main

import (
	"fmt"
)

func main() {
	var a, b float64

	fmt.Scanf("%lf %lf", &a, &b)

	fmt.Printf("%.3lf km/l\n", a/b)

}
