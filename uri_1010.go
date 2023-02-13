package main

import (
	"fmt"
)

func main() {
	var a, b int
	var c, d, res float64

	fmt.Scanf("%d %d %f%f", &a, &b, &c, &d)

	res = (a * c) + (b * d)

	fmt.Printf("VALOR A PAGAR: R$ %.2lf\n", res)

}
