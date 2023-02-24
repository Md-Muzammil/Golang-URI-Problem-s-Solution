package main
import ("fmt")

func mainss (){

	var a, b, c, d, e, i, j, k float64
    var avg, avrg float64
    fmt.Scanf("%f %f %f %f", &a, &b, &c, &d)
    avg = ((a*2) + (b*3) + (c*4) + d) / 10
    fmt.Printf("Media: %.1lf\n", avg)
    if ( avg < 5.0 ){
        fmt.Printf("Aluno reprovado.\n")
    }
    else if( avg >= 7.0 ){
        fmt.Printf("Aluno aprovado.\n")
    }
    else{
        fmt.Printf("Aluno em exame.\n")
        fmt.Scanf("%f",&e)
        fmt.Printf("Nota do exame: %.1f\n",e)
        avrg = (avg+e) / 2
        if ( avrg >= 5.0 ){
            fmt.Printf("Aluno aprovado.\n")
        }
        else {
			fmt.Printf ("Aluno reprovado.\n")
        }
        	fmt.Printf("Media final: %.1lf\n",avrg)
    }


}




