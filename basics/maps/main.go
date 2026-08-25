package main

import "fmt"

func main() {
	// Un map es una tabla hash: clave -> valor.

	// mapa vacío con make
	edades := make(map[string]int)
	edades["Ana"] = 25
	edades["Juan"] = 30
	fmt.Println(edades) // map[Ana:25 Juan:30]

	// map literal
	capitales := map[string]string{
		"Argentina": "Buenos Aires",
		"Brasil":    "Brasília",
		"Chile":     "Santiago",
	}
	fmt.Println(capitales["Argentina"])

	// actualizar un valor: se sobreescribe, sin error si ya existía
	edades["Ana"] = 26
	fmt.Println(edades["Ana"])

	// leer una clave que no existe devuelve el zero value del tipo, SIN error
	fmt.Println(edades["No existe"]) // 0 (zero value de int)

	// comma-ok: la forma correcta de distinguir "no existe" de "existe y vale 0"
	valor, existe := edades["No existe"]
	fmt.Println(valor, existe) // 0 false

	edad, existe := edades["Ana"]
	fmt.Println(edad, existe) // 26 true

	// borrar una clave
	delete(edades, "Juan")
	fmt.Println(edades)

	// borrar una clave que no existe no da error, es un no-op
	delete(edades, "No existe")

	// len funciona igual que en slices
	fmt.Println("Cantidad de capitales:", len(capitales))

	// iterar con range: el orden de iteración NO está garantizado
	for pais, capital := range capitales {
		fmt.Println(pais, "->", capital)
	}

	// el zero value de un map es nil, no un map vacío
	var nulo map[string]int
	fmt.Println(nulo == nil) // true
	fmt.Println(nulo["x"])   // 0, LEER de un map nil es seguro

	// nulo["x"] = 1 // esto PANIC: assignment to entry in nil map
	// si vas a escribir, siempre inicializá con make o un literal

	// values pueden ser structs, slices, otros maps, etc.
	type Punto struct{ X, Y int }
	puntos := map[string]Punto{
		"origen": {X: 0, Y: 0},
		"a":      {X: 1, Y: 2},
	}
	fmt.Println(puntos["a"].X)

	// map de slices: útil para agrupar valores por clave
	grupos := map[string][]string{}
	grupos["frutas"] = append(grupos["frutas"], "manzana")
	grupos["frutas"] = append(grupos["frutas"], "banana")
	fmt.Println(grupos)

	// no se puede usar map, slice ni func como clave (no son comparables)
	// pero sí como valor. Como clave sí se puede usar: int, string, structs
	// comparables (sin slices/maps adentro), etc.
}
