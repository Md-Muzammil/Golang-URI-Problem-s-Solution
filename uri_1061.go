package main
import ("fmt"
		"reflect"
)

func main(){
{
		var a[4], b[2],c[2],d[2],e[2],f[4]string
		var i,j,k,l,m,n,p,q,r,s,t,u,v,w,x,y,z,h[5] int
		fmt.canf("%s%d", a,&i)
		fmt.Scanf("%d%s%d%s%d", &j,&b,&k,&c,&l)
		fmt.Scanf("%s%d", &f,&p)
		fmt.Scanf("%d%s%d%s%d", &q, &d,&r,&e,&s)
		m=60-l+s
		n=(60-k-1+r)*60
		t=(24-1-j+q)*3600
		u=(p-i-1)*86400
		v=m+n+t+u
	
		h[0]=v/86400
		fmt.Printf("%d dia(s)n", h[0])
	
		h[0]=v%86400
		h[1]=h[0]/3600
		fmt.Printf("%d hora(s)n", h[1])
	
		h[1]=h[0]%3600
		h[2]=h[1]/60
		fmt.Printf("%d minuto(s)n", h[2])
	
		h[2]=h[1]%60
		fmt.Printf("%d segundo(s)n", h[2])
	}
	
}


