package main
import ("fmt")


var cont int
var fibonacci( n) int{
    cont++
    if(n == 0) return 0
    if(n == 1) return 1
    else return fibonacci(n-1) + fibonacci(n-2)
}
func main(){
    var n, o int
    fmt.Scanf("%d",&n);
    for(int i = 0; i < n; i++){
            fmt.Scanf("%d",&o)
            cont = 0
            fmt.Printf("fib(%d) = %d calls = %d\n",o,cont-1,fibonacci(o))
    }

}


