package main

import "fmt"

// Go pasa TODO por valor (copia) salvo que uses punteros explícitamente.
// Esto incluye structs, arrays, e incluso los valores dentro de slices/maps
// cuando los sacás con range (range también copia).

func incrementar(x int) {
	x++ // modifica la copia local, no afecta al original
}

func incrementarPuntero(x *int) {
	*x++ // "desreferencia" el puntero y modifica el valor original
}

type Contador struct {
	valor int
}

// receiver por valor: recibe una COPIA del struct
func (c Contador) SumarSinEfecto() {
	c.valor++ // no sirve de nada, se pierde al terminar la función
}

// receiver por puntero: opera sobre el struct original
func (c *Contador) Sumar() {
	c.valor++
}

func main() {
	// & obtiene la dirección de memoria de una variable (crea un puntero)
	x := 10
	p := &x
	fmt.Println("valor de x:", x)
	fmt.Println("dirección de x:", p)
	fmt.Println("valor apuntado por p:", *p) // * desreferencia: "el valor que está ahí"

	*p = 20 // modificar a través del puntero modifica x también
	fmt.Println("x después de *p = 20:", x)

	// pasar por valor: la función no puede modificar el original
	n := 5
	incrementar(n)
	fmt.Println("n después de incrementar (por valor):", n) // sigue en 5

	// pasar por puntero: la función SÍ puede modificar el original
	incrementarPuntero(&n)
	fmt.Println("n después de incrementarPuntero:", n) // 6

	// lo mismo pasa con structs
	c := Contador{valor: 0}
	c.SumarSinEfecto()
	fmt.Println("valor después de SumarSinEfecto:", c.valor) // 0

	c.Sumar()                                       // Go hace (&c).Sumar() automáticamente
	fmt.Println("valor después de Sumar:", c.valor) // 1

	// new(T) crea un puntero a un zero value de T
	pc := new(Contador)
	pc.Sumar()
	fmt.Println("pc.valor:", pc.valor) // 1

	// punteros a punteros existen pero son poco comunes
	pp := &p
	fmt.Println("valor final a través de doble puntero:", **pp)

	// el zero value de un puntero es nil
	var nulo *int
	fmt.Println(nulo == nil) // true
	// *nulo // esto PANIC: nil pointer dereference

	// ¿cuándo usar puntero vs valor?
	// - puntero: si necesitás modificar el original, o el struct es grande
	//   (copiar sería caro), o querés representar "puede no existir" (nil)
	// - valor: para datos chicos e inmutables, es más simple y evita bugs
	//   de aliasing (dos variables apuntando al mismo dato sin querer)
}
