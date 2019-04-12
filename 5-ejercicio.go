// Leer 2 vectores A y B ambos de dimensión N y
// luego calcular e imprimir el vector C que sea la
// suma de A+B.

package main

import "fmt"

func main() {
    var numero, n int 
    fmt.Println("El tamaño del slice");    
    fmt.Scanf("%v\n",&n);
// declaracion de arrays   
	vector := make([]int, n)
	vectorB := make([]int, n)
	vsuma := make([]int, n)
	//Carga del primer vector
	fmt.Println("Vector A")
    for i:= 0 ; i < n; i++ {
        fmt.Printf("\nValor para la posicion %v: ",i)
        fmt.Scanf("%v\n",&numero)
		vector[i] = numero		
	}
	fmt.Println("Vector B")
	// Carga del segundo vector
	for i:= 0 ; i < n; i++ {
        fmt.Printf("\nValor para la posicion %v: ",i)
        fmt.Scanf("%v\n",&numero)
		vectorB[i] = numero		
	}
	//Suma de ambos vectores almacenados en un tercer vector
	for i:= 0 ; i < n; i++ {
		vsuma[i] = vector[i]+vectorB[i]
		fmt.Println(vsuma[i])
	}

    //fmt.Print("Resultado de nros impares: de Suma: %d y Promedio: %d",suma,promedio)   
}
