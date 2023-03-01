package main

func main(){

	var x float64
    var i int
	var sum=0
    for(i=1;i<=6;i++){
        fmt.Scanf("%f",&x)
        if(x>0){
        sum=sum+1
        }
    }
    fmt.Printf("%d valores positivos\n",sum)


}