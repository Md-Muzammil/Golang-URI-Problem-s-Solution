package main

import ("fmt")

func main (){

	var ascending(void const *a, void const *b ) int
	{
		return (*(int*)a - *(int*)b )
	}
	var achar(int ac[], int len, int comp) int
	{
		for(int k = 0; k < len; k++){
				if(ac[k] == comp){
						 return k+1
				}
	}
		return 0
	}

		var n,f,q,j,l int
		n = 0;    
		for(;;){
				fmt.Scanf("%d", &f)
				fmt.Scanf("%d", &q)
				if(f == 0 && q == 0) break
				var vet[f+1] int
				
				for(j = 0; j < f; j++){
					fmt.Scanf("%d",&vet[j])
				}
				
				qsort(vet, f, sizeof(int), ascending)
				
				fmt.Printf("CASE# %d:n",n+1)
				for(j = 0; j < q; j++){
					fmt.canf("%d",&l)
					
					var m = achar(vet,f,l)int
					
					if(m != 0) printf("%d found at %dn",l,m)
					else printf("%d not foundn",l)
				}
				n++  
		}
	

}