package main

import (
	"fmt"
)

func main() {

	var numeros [5]int
	numeros[0] = 10
	numeros[1] = 20

	// Se puede eliminar un indice?

	// numeros[1] = nil // Esto no es posible, ya que el tipo de dato es int y no puede ser nil

	// o directo con valores
	frutas := [3]string{"manzana", "banana", "pera"}
	fmt.Println(frutas)
	fmt.Println("El indice 0 de frutas es:", frutas[0]) // manzana
	fmt.Println(frutas[:])                              // [manzana banana pera]
	fmt.Println(frutas[1:])                             // [banana pera]

	// que Go cuente por vos
	dias := [...]string{"lun", "mar", "mier", "jue", "vie"}
	fmt.Println(len(dias)) // 5

	// slices

	// declarar y crear con valores
	mas_numeros := []int{1, 2, 3, 4, 5}
	fmt.Println(mas_numeros)

	// slice vacío, con make (largo 0, capacidad 3)
	otros := make([]int, 0, 3)
	fmt.Println(otros)
	// otros[0] = 1 // Esto va a dar error, ya que no hay elementos en el slice
	// para agregar elementos, se usa append
	otros = append(otros, 1)
	otros = append(otros, 2)
	otros = append(otros, 3)
	fmt.Println(otros)

	// agregar más elementos de los que tiene capacidad
	fmt.Println("Trato de agregar un 4to elemento")
	otros = append(otros, 4) // Se puede perfectamente agregar un 4to elemento, ya que Go va a crear un nuevo slice con el doble de capacidad y copiar los elementos del slice original
	fmt.Println(otros)
	fmt.Println("Largo:", len(otros), "Capacidad:", cap(otros), "Contenido:", otros)

	// slice con largo inicial 5 (todos en 0)
	ceros := make([]int, 5)
	fmt.Println(ceros)

	numeros1 := []int{10, 20, 30, 40, 50}

	fmt.Println(numeros1[1:3]) // [20 30]  -> índices 1 y 2 (el 3 no incluye)
	fmt.Println(numeros1[:2])  // [10 20]  -> desde el principio hasta índice 2
	fmt.Println(numeros1[2:])  // [30 40 50] -> desde índice 2 hasta el final

	frutas2 := []string{"manzana", "banana", "pera"}

	for i, fruta := range frutas2 {
		fmt.Println(i, fruta)
	}

	// si no te importa el índice
	for _, fruta := range frutas2 {
		fmt.Println(fruta)
	}

	original := []int{1, 2, 3, 4, 5}
	recorte := original[1:3]
	recorte[0] = 999
	recorte = append(recorte, 888) // Esto va a modificar el slice original, ya que recorte y original comparten la misma memoria hasta que se hace append y se crea un nuevo slice con más capacidad
	fmt.Println(recorte)           // [999 3 888]

	fmt.Println(original) // [1 999 3 4 5]  -> ¡se modificó el original!

	// Si necesitamos copias:

	copia := make([]int, len(original))
	copy(copia, original)

}
