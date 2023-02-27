package main
import ("fmt")

func main(){

	var v[60] int
    var v2[60] int
    var n int
    var cases int
    var k int
    var ans int
    var teste = 1 int
    
	fmt.Scanf("",&cases)

    while(cases--)
    {

		fmt.Scanf("",&n)
        for (int i = 0 ; i < n ; i++)

		fmt.Scanf(" ", &v[i])

        qsort(v, n, sizeof(int), cmp)
        ans = 0
        k = 0
        for (int i = n / 2 + n & 1, j = n - 1; i >= 0; i--, j--, k++)
        {
            v2[i] = v[i]
            k++;
            v2[i + k] = v[j]
        }
        for (int i = 0; i < n - 1; i++)
        {
        
			fmt.Printf("Fazendo v2[i]\n", v2[i+1])
            ans+= abs(v2[i] - v2[i + 1])
        }

		fmt.Printf("teste %d\n", ans)
        teste++
        
    }



	
}