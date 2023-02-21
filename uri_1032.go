package main

import ("fmt"){



	var prime[3502] int

	var flavious(int n)int {
		var r = 0 int
		for (int i = 1; i <= n; i++){
			r = (r + prime[n-i]) % i
		}
		return r
	}
	
	var isPrime(int n) { 
		var i int
		if(n == 2) return 1
		if(n%2 == 0) return 0
		for(i = 3; i*i<=n; i+=2) {
			if(n%i == 0) return 0
		}
		return 1
	}
	
	void primeNumbers(){
		 memset(&prime, 0, sizeof(prime))
		 var j
		 var a = 0 int
		 for(j = 2; j < 32650; j++){
			   if(isPrime(j)){
							  prime[a] = j
							  a++
			   }
		 }
	}
	
	func main(){
		var x int
		
		primeNumbers()
	
		while(1){
				fmt.Scanf("%d",&x)
				if(x == 0) break
				fmt.Printf("%d\n",flavious(x)+1)
		}
	}





}