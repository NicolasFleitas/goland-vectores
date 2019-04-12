// Hallar el máximo elemento de un vector de N
// componentes y la posición que ocupa dentro
// del mismo. Si hay varios máximos entonces
// imprimir el primero que aparezca. 

package main

import "fmt"

func main() {
    var numero, n,posicion int // para el tamaño del vector
    fmt.Println("El tamaño del slice");    
    fmt.Scanf("%v\n",&n);
// declaracion de array    
    vector := make([]int, n)
    
    for i:= 0 ; i < n; i++ {
        fmt.Printf("\nValor para la posicion %v: ",i)
        fmt.Scanf("%v\n",&numero)
		vector[i] = numero		
	}

	//buscar el maximo elemento y su posicion opcion 1
	max:=vector[0]
	for i:= 1 ; i < n; i++ {        
		if(vector[i]>max) {
			max = vector[i];
			posicion = i;
		}
	}

    fmt.Printf("El maximo elemento es %v\nSe encuentra en la posicion %v",max,posicion)   
}
