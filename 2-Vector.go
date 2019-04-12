// Leer un vector de dimensión N y luego hallar el
// promedio de los elementos en él que no sean
// cero. 

package main
import "fmt"
func main() {
	var numero, n, suma,cantidad int // para el tamaño del vector
    fmt.Println("El tamaño del vector");    
    fmt.Scanf("%v\n",&n);
// declaracion de array    
    vector := make([]int, n)
    
    for i:= 0 ; i < n; i++ {
        fmt.Printf("\nIntroduce el numero que ira en la posicion %v: ",i)
		fmt.Scanf("%v\n",&numero)
		vector[i] = numero
		if vector[i] != 0 {
			fmt.Println("Es valido: %d",vector[i])
			suma+=vector[i]
			cantidad++
		}		
    }
    
    promedio:= suma/cantidad
    fmt.Printf("Suma de elementos que !CERO: %d \n Promedio: %d",suma,promedio)   
}
