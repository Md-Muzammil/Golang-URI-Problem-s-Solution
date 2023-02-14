package main

import (
	"fmt"
)

func main() {

	var a, b, c float64

	fmt.Scanf("%lf %lf %lf", &a, &b, &c)

	fmt.Printf("TRIANGULO: %.3lf\n", (a*c)/2)
	fmt.Printf("CIRCULO: %.3lf\n", c*c*3.14159)
	fmt.Printf("TRAPEZIO: %.3lf\n", ((a+b)/2)*c)
	fmt.Printf("QUADRADO: %.3lf\n", b*b)
	fmt.Printf("RETANGULO: %.3lf\n", a*b)
}
