package main

import "fmt"

// Una interface define un conjunto de métodos. Cualquier tipo que tenga
// esos métodos la implementa AUTOMÁTICAMENTE, sin decir "implements"
// (esto se llama "duck typing" o tipado estructural).
type Figura interface {
	Area() float64
	Perimetro() float64
}

type Rectangulo struct {
	Ancho, Alto float64
}

func (r Rectangulo) Area() float64      { return r.Ancho * r.Alto }
func (r Rectangulo) Perimetro() float64 { return 2 * (r.Ancho + r.Alto) }

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64      { return 3.1416 * c.Radio * c.Radio }
func (c Circulo) Perimetro() float64 { return 2 * 3.1416 * c.Radio } // Si comentamos este método, Circulo deja de implementar Figura y no se puede pasar a describir().

// esta función no sabe ni le importa si recibe un Rectangulo o un Circulo,
// solo le importa que cumpla la interface Figura. Esto es polimorfismo.
func describir(f Figura) {
	fmt.Printf("Área: %.2f, Perímetro: %.2f\n", f.Area(), f.Perimetro())
}

// fmt.Stringer es una interface del stdlib: si un tipo tiene String() string,
// fmt.Println y compañía la usan automáticamente para imprimirlo.
type Persona struct {
	Nombre string
	Edad   int
}

func (p Persona) String() string {
	return fmt.Sprintf("%s (%d años)", p.Nombre, p.Edad)
}

func main() {
	r := Rectangulo{Ancho: 3, Alto: 4}
	c := Circulo{Radio: 5}

	describir(r)
	describir(c)

	// un slice de la interface puede contener distintos tipos concretos
	figuras := []Figura{r, c, Rectangulo{Ancho: 1, Alto: 1}}
	total := 0.0
	for _, f := range figuras {
		total += f.Area()
	}
	fmt.Println("Área total:", total)

	// any (alias de interface{}) acepta CUALQUIER tipo, no tiene métodos
	var cualquiera any
	cualquiera = 42
	fmt.Println(cualquiera)
	cualquiera = "ahora soy un string"
	fmt.Println(cualquiera)

	// type assertion: "afirmo que el valor dentro de la interface es de tipo X"
	var f Figura = c
	circuloOriginal, ok := f.(Circulo)
	if ok {
		fmt.Println("era un círculo de radio:", circuloOriginal.Radio)
	}

	// si la afirmación es incorrecta y no usás la forma con ", ok", hace panic
	_, ok = f.(Rectangulo)
	fmt.Println("¿es un rectángulo?", ok) // false

	// type switch: para manejar varios tipos posibles de forma prolija
	for _, fig := range figuras {
		switch v := fig.(type) {
		case Rectangulo:
			fmt.Println("es un rectángulo de", v.Ancho, "x", v.Alto)
		case Circulo:
			fmt.Println("es un círculo de radio", v.Radio)
		default:
			fmt.Println("figura desconocida")
		}
	}

	// fmt.Stringer en acción: Println llama a String() automáticamente
	p := Persona{Nombre: "Ana", Edad: 25}
	fmt.Println(p) // Ana (25 años), no {Ana 25}

	// la interface vacía (any) se usa para funciones genéricas tipo fmt.Println,
	// pero hoy en día para eso conviene usar Generics (type parameters) cuando aplica
}
