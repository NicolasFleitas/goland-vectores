// Hallar el producto escalar de los vectores A y B
// de dimensiones N respectivamente. 

package main
import "fmt"

func main(){	
	var numero, n, resultado int 
    fmt.Println("El tamaño del slice");    
    fmt.Scanf("%v\n",&n);
// declaracion de array 
	v1 := make([]int, n)
	v2 := make([]int, n)
	
	//Carga del primer vector
	fmt.Println("Vector 1")
    for i:= 0 ; i < n; i++ {
        fmt.Printf("\nValor para la posicion %v: ",i)
        fmt.Scanf("%v\n",&numero)
		v1[i] = numero		
	}
	fmt.Println("Vector 2")
	// Carga del segundo vector
	for i:= 0 ; i < n; i++ {
        fmt.Printf("\nValor para la posicion %v: ",i)
        fmt.Scanf("%v\n",&numero)
		v2[i] = numero		
	}

	for i:= 0 ; i < n; i++ {
		resultado += v1[i] * v2[i];
	}
	fmt.Printf("Resultado escalar %v",resultado)
}

