package main

import ("fmt")

func main(){


var t int
 
fmt.Scanf("%d", &t);

var h = t / 3600 
 t= t - ( h * 3600);

 var m = t / 60 
 var t = t - ( m * 60) int

fmt.Printf("%d:%d:%d\n" , h , m , t )
}



