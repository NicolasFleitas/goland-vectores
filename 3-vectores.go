// 3-Leer un vector de N componentes. Hallar e
// imprimir el promedio de los valores impares. 

package main

import "fmt"

func main() {
    var numero, n, suma,cantidad int // para el tamaño del vector
    fmt.Println("El tamaño del slice");    
    fmt.Scanf("%v\n",&n);
// declaracion de array    
    vector := make([]int, n)
    
    for i:= 0 ; i < n; i++ {
        fmt.Printf("\nValor para la posicion %v: ",i)
        fmt.Scanf("%v\n",&numero)
		vector[i] = numero
		if vector[i] % 2 != 0 {
			suma+=vector[i]
			cantidad++
		}
    }
    
    promedio:= suma/cantidad
    fmt.Print("Resultado de nros impares: de Suma: %d y Promedio: %d",suma,promedio)   
}
