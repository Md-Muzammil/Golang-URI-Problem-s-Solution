package main

import (
	"fmt"
)

func main() {

	var salary, value, total float64
	var name rune
	fmt.Scanf("%s %lf %lf", &name, &salary, &value)

	total = salary + value*.15
	fmt.Printf("TOTAL = R$ %.2lf\n", total)

}
