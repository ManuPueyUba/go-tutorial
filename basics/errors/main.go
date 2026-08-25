package main

import (
	"errors"
	"fmt"
)

// En Go los errores son VALORES normales, no excepciones. Por convención,
// una función que puede fallar devuelve (resultado, error) y el error va
// último. Si error es nil, no hubo problema.

var ErrDivisionPorCero = errors.New("no se puede dividir por cero")

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionPorCero
	}
	return a / b, nil
}

// fmt.Errorf con %w "envuelve" (wrap) otro error, preservando la cadena
// de causas para poder inspeccionarla después con errors.Is / errors.As.
func procesarPedido(id int) error {
	_, err := dividir(100, 0)
	if err != nil {
		return fmt.Errorf("procesando pedido %d: %w", id, err)
	}
	return nil
}

// se pueden definir tipos de error propios implementando el método Error()
type ErrorValidacion struct {
	Campo string
}

func (e *ErrorValidacion) Error() string {
	return fmt.Sprintf("el campo %q es inválido", e.Campo)
}

func validarEdad(edad int) error {
	if edad < 0 {
		return &ErrorValidacion{Campo: "edad"}
	}
	return nil
}

func main() {
	// patrón estándar: if err != nil
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}

	_, err = dividir(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// errors.Is: ¿este error ES (o envuelve) este error específico?
	err = procesarPedido(42)
	if errors.Is(err, ErrDivisionPorCero) {
		fmt.Println("Detectamos que la causa raíz fue división por cero")
	}
	fmt.Println("Mensaje completo:", err)

	// errors.As: ¿este error ES (o envuelve) un error de ESTE TIPO?
	// sirve para recuperar campos extra del error concreto
	err = validarEdad(-5)
	var errVal *ErrorValidacion
	if errors.As(err, &errVal) {
		fmt.Println("Campo inválido:", errVal.Campo)
	}

	// defer: ejecuta la función al FINAL, cuando la función actual retorna
	// (sin importar por dónde retorne). Se usa mucho para liberar recursos.
	ejemploDefer()

	// panic / recover: para errores realmente excepcionales (bugs, estado
	// imposible). No es el mecanismo normal de manejo de errores en Go.
	resultadoSeguro := dividirSeguro(10, 0)
	fmt.Println("resultadoSeguro:", resultadoSeguro)
}

func ejemploDefer() {
	fmt.Println("1: entrando a la función")
	defer fmt.Println("4: esto se imprime al salir (defer 2)")
	defer fmt.Println("3: esto se imprime al salir (defer 1)")
	fmt.Println("2: haciendo trabajo")
	// los defer se apilan: si hay varios, se ejecutan en orden LIFO
}

func dividirSeguro(a, b float64) (resultado float64) {
	defer func() {
		// recover() detiene un panic en curso, solo funciona dentro de un defer
		if r := recover(); r != nil {
			fmt.Println("recuperado de un panic:", r)
			resultado = 0
		}
	}()

	if b == 0 {
		panic("división por cero no permitida acá")
	}
	return a / b
}
