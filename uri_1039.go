package main
import ("fmt")
func vbbn (){

var r1,x1,y1,r2,x2,y2 int
var distancia float64

while(
	fmt.Scanf("%d %d %d %d %d %d",&r1,&x1,&y1,&r2,&x2,&y2) != EOF){
				distancia = sqrt(pow(x2-x1,2) + pow(y2-y1,2))
				if(r1 >= distancia+r2)
				fmt.Printf("RICO\n")
				else
				fmt.Printf("MORTO\n")
}





}





