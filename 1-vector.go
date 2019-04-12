/*Leer un vector de N componentes, sumar sus
componentes e imprimir la suma y el promedio
de los números.
*/
package main

import "fmt"

func main() {
    var numero, n, suma int // para el tamaño del vector
    fmt.Println("El tamaño del slice");    
    fmt.Scanf("%v\n",&n);
// declaracion de array    
    vector := make([]int, n)
    
    for i:= 0 ; i < n; i++ {
        fmt.Printf("\nIntroduce el numero que ira en la posicion %v: ",i)
        fmt.Scanf("%v\n",&numero)
        vector[i] = numero
        suma+=vector[i]
    }
    
    promedio:= suma/len(vector)
    fmt.Printf("Resultado: de Suma: %d y Promedio: %d",suma,promedio)   
}
