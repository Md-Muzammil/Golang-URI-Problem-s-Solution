package main
import ("fmt")


var flavious(int x, int y){
    if(x == 1) return 0
    return (flavious(x-1,y)+y)%x
}

var remaining(int n, int k) {
    var r = 0 int
    for (int i = 2; i <= n; i++)
        r = (r + k) % i

    return r
}

func main(){
    var n, x, y, j, num, pulo int
    
    fmt.Scanf("%d",&n)
    
    for(var i = 0; i < n; i++ int){
            fmt.Scanf("%d %d",&x,&y)
            //printf("Case %d: %d\n",i+1,flavious(x,y)+1)
            fmt.Printf("Case %d: %d\n",i+1,remaining(x,y)+1)
    }
}

