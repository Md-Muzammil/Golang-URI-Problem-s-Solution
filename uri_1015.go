package main

import (
	"fmt"
)

func main() {
	var x1, x2, y1, y2, x, y, temp, distance float64
	fmt.Scanf("%lf %lf", &x1, &y1)
	fmt.Scanf("%lf %lf", &x2, &y2)
	x = x2 - x1
	y = y2 - y1
	temp = x*x + y*y
	distance = sqrt(temp)
	fmt.Printf("%.4lf\n", distance)

}
