package main

import (
	"fmt"
)

func main() {

	var a, b, c, media float64

	fmt.Scanf("%f%f%f", &a, &b, &c)

	media = (a / 10 * 2) + (b / 10 * 3) + (c / 10 * 5)

	fmt.Println("EDIA =", media)
}
