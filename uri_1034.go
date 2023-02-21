package main

import ("fmt")



var pd[N] int

var blocos(int valores[], int n, int m) int{
    var i, j, k int
    
    if(m%valores[n-1] == 0) return m/valores[n-1]
    
    memset(pd,999999,sizeof(pd))
    pd[0] = 0
    
    for(i=0; i < n; i++){
        k = valores[i]
        for(j = k; j <= m; j++){
             pd[j] = std::min(pd[j], pd[j-k]+1)
        }
    }
    
    return pd[m]
}


func main(){
    var valores[26] int
    var t, n, m int
    var i, j int
    
    fmt.Scanf("%d", &t)
    for(i = 0; i < t; i++){
          fmt.Scanf("%d %d", &n, &m)
          for (j = 0; j < n; j++) 
		  fmt.Scanf("%d", &valores[j])
          qsort(valores, n, sizeof(int), compare)
          fmt.Printf("%d\n", blocos(valores,n,m))
    }


