package main
import ("fmt")

func main(){

	var a, b, c, d int
    var dif int
    fmt.Scanf("%d %d %d %d", &a, &c, &b, &d)
    dif = ((b*60)+d) - ((a*60)+c)
    if(dif<=0) dif += 24*60
    fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", dif/60, dif%60)

}