package main

import ("fmt") 

func main(){
    var a,b,c,d int
    fmt.Scanf("%d%d", &a, &b)
    if(a<b)
    {
        for(c=a+1; c<b; c++)
        {
            if(c%5==2 || c%5==3)
                fmt.Printf("%d\n",c)
        }
    }
    else if(a>b)
    {
        for(c=b+1; c<a; c++)
        {
            if(c%5==2 || c%5==3)
                fmt.Printf("%d\n",c)
        }
    }
}

