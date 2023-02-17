package main
import ("fmt")


var mdc( x, y) int
{
    var aux = 0 int
    for(;y%x != 0;){
                      aux = x
                      x = y%x
                      y = aux
            }
            return x
}

package main(){
    
    var n, x, y, aux int
    
    fmt.Scanf("%d",&n)
    
    for(int i = 0; i < n; i++){
            fmt.Scanf("%d",&x)
            fmt.Scanf("%d",&y)
            
            if(x > y){
                aux = x
                x = y
                y = aux
            }
            
            fmt.Printf("%d\n", mdc(x,y))
    }
}  

