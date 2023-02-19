package main
import ("fmt")


var remaining( n,  k)int {
    int r = 0;
    for (int i = 1; i < n; i++)
        r = (r + k) % i

    return r
}

func main(){
    var n, x, y, j, num, pulo int
    while(1){
           fmt.Scanf("%d",&n)
           if(n == 0) break
           y = 1
           for(;;){
                   if(remaining(n,y) != 11) y++
                   else break
           }
           fmt.Printf("%d\n",y)
    }

