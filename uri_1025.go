package main

import ("fmt")

func main (){

	var n,f,q,j,l int
    n = 0;    
    for(;;){
            fmt.Scanf("%d", &f)
            fmt.Scanf("%d", &q)
            if(f == 0 && q == 0) break
            int vet[f+1]
            
            for(j = 0; j < f; j++){
                fmt.Scanf("%d",&vet[j])
            }
            
            qsort(vet, f, sizeof(int), ascending)
            
            fmt.Printf("CASE# %d:n",n+1)
            for(j = 0; j < q; j++){
                fmt.Scanf("%d",&l)
                
                var m = achar(vet,f,l) int
                
                if(m != 0)
				fmt.Printf("%d found at %dn",l,m)
                else
				fmt.Printf("%d not foundn",l)
            }
            n++;  
    }
	

}