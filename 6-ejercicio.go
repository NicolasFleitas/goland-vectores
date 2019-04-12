// 6-Leer un vector X de N componentes y crear el
// vector C con los elementos pares de X. 
package main
import "fmt"

func main() {
    var numero, n int 
    fmt.Println("El tamaño del slice");    
    fmt.Scanf("%v\n",&n);	
	vectorX := make([]int, n)	

	var dimension int = 0 

    for i:= 0 ; i < n; i++ {
        fmt.Printf("\nValor para la posicion %v: ",i)
		fmt.Scanf("%v\n",&numero)
		vectorX[i] = numero		
		if vectorX[i] % 2 == 0 {
			dimension++
		}
	}
	vectorPar := make([]int, dimension)	// tamaño del vector con la cantidad de pares del vector X
	var j int
	for i:= 0 ; i < n; i++ {
		if vectorX[i] % 2 == 0 {			
			vectorPar[j] = vectorX[i]
			j++			
		}
	}

	fmt.Printf("\nEl vector resultante es\n")	
	for i:= 0 ; i < len(vectorPar); i++ {
		fmt.Println(vectorPar[i])
	}
    
}
