package main
import (fmt)



func main(){
    var a, b, c float64
	fmt.Scanf(&a,&b,&c)

    if((pow(b,2) - 4*a*c) < 0 || 2*a == 0)
				fmt.Printf("Impossivel calcular\n")
    else{
        		 fmt.Printf("R1 = %.5lf\n",((-b) + sqrt((pow(b,2) - 4*a*c)))/(2*a))
         		fmt.Printf("R2 = %.5lf\n",((-b) - sqrt((pow(b,2) - 4*a*c)))/(2*a))
        }
    
}






