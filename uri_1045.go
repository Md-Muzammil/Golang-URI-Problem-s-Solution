package main
import ("fmt")

func main() {


	var A, B, C, D, E float64
    fmt.Scanf("%lf %lf %lf", &A, &B, &C)
    if (A < B){
        A = A + B
        B = A - B
        A = A - B
    }
    if (A < C){
        A = A + C
        C = A - C
        A = A - C
    }
    if (B < C){
        B = B + C
        C = B - C
        B = B - C
    }
    D = A * A
    E = (B * B) + (C * C)
    if(A >= B+C)
	fmt.Printf("NAO FORMA TRIANGULO\n")
    else{
        if(D == E)
		fmt.Printf("TRIANGULO RETANGULO\n")
        if(D > E)
		fmt.Printf("TRIANGULO OBTUSANGULO\n")
        if(D < E)
		fmt.Printf("TRIANGULO ACUTANGULO\n")
        if(A == B && B == C)
		fmt.Printf("TRIANGULO EQUILATERO\n")
        else if(A == B || B == C || A == C)
		fmt.Printf("TRIANGULO ISOSCELES\n")
    }
    



}

