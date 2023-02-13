package main

import (
	"fmt"
)

func main() {
	var NUMBER, hours int
	var amount, SALARY float64
	fmt.Scanf("%d %d %f", &NUMBER, &hours, &amount)
	SALARY = hours * amount
	fmt.Printf("NUMBER = %d\n", NUMBER)
	fmt.Printf("SALARY = U$ %0.2f\n", SALARY)

}
