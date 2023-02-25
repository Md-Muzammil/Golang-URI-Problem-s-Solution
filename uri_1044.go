package main
import ("fmt")
func main(){

	var A, B int
    fmt.Scanf("%d %d", &A, &B)
    if (B % A == 0 || A % B == 0)
    {
        fmt.Printf("Sao Multiplosn")
    }
    else
    {
        fmt.Printf("Nao sao Multiplosn")
    }



}

