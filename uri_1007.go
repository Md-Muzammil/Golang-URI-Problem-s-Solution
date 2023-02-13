package main

import "fmt"

func main() {

	var A, B, C, D, difference int

	fmt.Scanf("%d %d %d %d", &A, &B, &C, &D)

	difference = ((A * B) - (C * D))

	fmt.Printf("DIFERENCA = %d\n", difference)

}
