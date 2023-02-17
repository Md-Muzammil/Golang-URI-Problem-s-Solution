package main
import ("fmt")

func main(){

	var n, m int
    
    while(
		fmt.Scanf("%u %u", &n, &m) != EOF){
                    
        fmt.Printf("%u\n", n ^ m)
 }
}