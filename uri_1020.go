package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Scanf("%d", &n)

	var a = n / 365
	var rA = n % 365
	var rM = rA % 30
	var m = rA / 30
	var d = rM % 30

	fmt.Printf("%d ano(s)n%d mes(es)n%d dia(s)n", a, m, d)
}
