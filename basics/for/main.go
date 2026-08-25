package main

import "fmt"

func main() {
	// Go tiene un solo loop: "for". No hay while ni do-while,
	// pero "for" se puede escribir de varias formas.

	// 1) for clásico: init; condición; post
	for i := 0; i < 5; i++ {
		fmt.Println("clásico:", i)
	}

	// 2) for como "while": solo la condición
	n := 0
	for n < 3 {
		fmt.Println("como while:", n)
		n++
	}

	// 3) for sin nada: loop infinito, hay que cortarlo con break
	contador := 0
	for {
		if contador == 3 {
			break
		}
		fmt.Println("infinito:", contador)
		contador++
	}

	// 4) continue: saltea a la siguiente iteración
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue // saltea los pares
		}
		fmt.Println("impar:", i)
	}

	// 5) range sobre un slice: da índice y valor
	frutas := []string{"manzana", "banana", "pera"}
	for i, fruta := range frutas {
		fmt.Println(i, fruta)
	}

	// si no te importa el índice, se descarta con _
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}

	// si solo te importa el índice, se puede omitir el valor
	for i := range frutas {
		fmt.Println("índice:", i)
	}

	// 6) range sobre un map: da clave y valor (el orden NO está garantizado)
	edades := map[string]int{"Ana": 25, "Juan": 30}
	for nombre, edad := range edades {
		fmt.Println(nombre, "tiene", edad, "años")
	}

	// 7) range sobre un string: recorre RUNES (puntos de código Unicode), no bytes
	for i, r := range "café" {
		fmt.Printf("índice %d -> rune %c\n", i, r)
	}
	// ojo: "é" ocupa 2 bytes en UTF-8, por eso el índice salta de 3 a 5

	// 8) loops anidados con labels: break/continue pueden apuntar a un loop externo
external:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue external // corta la iteración externa, no la interna
			}
			fmt.Println("i:", i, "j:", j)
		}
	}

	// 9) for con múltiples variables en el post (poco común, pero válido)
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Println("i:", i, "j:", j)
	}
}
