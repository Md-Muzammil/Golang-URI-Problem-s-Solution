package main

import (
	"fmt"
)

func mainnh() {

	var salary, value, TOTAL float64
	var name string
	fmt.Scanf("%s %lf %lf", &name, &salary, &value)
	TOTAL = salary + value*.15
	fmt.Printf("TOTAL = R$ %.2lf\n", TOTAL)

}
