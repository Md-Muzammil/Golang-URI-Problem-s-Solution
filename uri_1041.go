package main
import ("fmt")

func main (){


	var x, y float64


    fmt.Scanf("%lf %lf", &x, &y)

    if (x == 0.0 && y == 0.0)


    {

        fmt.Printf("Origem\n")

    }

    else if (x == 0.0 && y != 0.0)

    {

        fmt.Printf("Eixo Y\n")

    }

    else if (y == 0.0 && x != 0.0)

    {
        fmt.Printf("Eixo X\n")
    }

    else if (x > 0.0)

    {

        if (y > 0.0)

        {
            fmt.Printf("Q1\n")
        }

    else 
		
		fmt.Printf("Q4\n")
    }

    else if (y > 0.0)

    {
        fmt.Printf("Q2\n")
    }

    else
	fmt.Printf("Q3\n")

}




