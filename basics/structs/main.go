package main

import "fmt"

// Un struct agrupa datos relacionados. Es lo más parecido que tiene Go
// a una "clase" sin métodos, aunque los métodos se agregan por afuera.
type Persona struct {
	Nombre string
	Edad   int
}

// Los métodos se declaran con un "receiver" antes del nombre de la función.
// Este es un receiver por VALOR: recibe una copia de la Persona.
func (p Persona) Saludar() string {
	return fmt.Sprintf("Hola, soy %s y tengo %d años", p.Nombre, p.Edad)
}

// Este es un receiver por PUNTERO: puede modificar la Persona original.
// Regla práctica: si el método modifica el struct, usá puntero.
func (p *Persona) Cumplir() {
	p.Edad++
}

// Composición: Go no tiene herencia, pero se puede "embeber" un struct
// dentro de otro. Los campos y métodos de Empleado quedan promovidos.
type Empleado struct {
	Persona // embebido (sin nombre de campo)
	Puesto  string
	Salario float64
}

func main() {
	// Zero value: todos los campos arrancan en su valor por defecto (0, "", false, nil...)
	var vacia Persona
	fmt.Println("Zero value:", vacia) // {  0}

	// Struct literal con nombres de campo (recomendado, es explícito)
	juan := Persona{Nombre: "Juan", Edad: 30}
	fmt.Println(juan)
	fmt.Println(juan.Saludar())

	// Struct literal posicional: hay que respetar el orden de los campos.
	// Es frágil si el struct cambia, mejor evitarlo salvo structs muy chicos.
	ana := Persona{"Ana", 25}
	fmt.Println(ana.Saludar())

	// Llamar un método con receiver puntero sobre una variable normal:
	// Go automáticamente hace (&juan).Cumplir(), no hace falta escribirlo.
	juan.Cumplir()
	fmt.Println("Después de cumplir:", juan.Edad) // 31

	// Puntero explícito a un struct
	pAna := &ana
	pAna.Edad = 26 // pAna.Edad es azúcar sintáctico de (*pAna).Edad
	fmt.Println(ana.Edad)

	// new(Persona) crea un puntero a un struct con zero values
	otra := new(Persona)
	otra.Nombre = "Pedro"
	fmt.Println(*otra)

	// Comparación: dos structs son iguales si TODOS sus campos son iguales
	// (solo funciona si todos los campos son comparables, ej: no tienen slices/maps)
	fmt.Println(Persona{"Ana", 26} == ana) // true

	// Composición / embedding
	emp := Empleado{
		Persona: Persona{Nombre: "Marta", Edad: 40},
		Puesto:  "Backend Developer",
		Salario: 1500000,
	}
	// Nombre y Edad y Saludar() quedan "promovidos": se accede como si fueran propios
	fmt.Println(emp.Nombre, emp.Puesto)
	fmt.Println(emp.Saludar())

	// Structs anónimos: útiles para datos de uso único (ej: un resultado temporal)
	punto := struct {
		X, Y int
	}{X: 1, Y: 2}
	fmt.Println(punto)

	// Slice de structs, algo muy común
	personas := []Persona{
		{Nombre: "Lucía", Edad: 22},
		{Nombre: "Diego", Edad: 35},
	}
	for _, p := range personas {
		fmt.Println(p.Saludar())
	}
}
