package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Random number:", get_random_number())
	fmt.Println("Random float:", get_random_float())
	fmt.Printf("The square root of 81 is: %d and 81^2 is: %d\n", int(math.Sqrt(81)), int(math.Pow(81, 2)))

	// convert a float to an int truncates de value it dosent round it
	// import pi
	value := 3.6
	fmt.Printf("The float %f as an int is: %d\n", value, int(value))

	// If you want to round a float to the nearest integer, you can use the math.Round function. For example:
	roundedValue := math.Round(value)
	fmt.Printf("The float %f rounded to the nearest integer is: %d\n", value, int(roundedValue))

	// If you want to round a float to a specific number of decimal places, you can use the following approach:
	decimalPlaces := 2
	multiplier := math.Pow(10, float64(decimalPlaces))
	roundedValueToDecimalPlaces := math.Round(value*multiplier) / multiplier
	fmt.Printf("The float %f rounded to %d decimal places is: %.2f\n", value, decimalPlaces, roundedValueToDecimalPlaces)

	// Si queremos buscar un maximo:
	fmt.Println("El maximo entre 3 y 7: es:", math.Max(3, 7)) // Compara 2 valores y devuelve el mayor de ellos

	fmt.Println("El minimo entre 3 y 7: es:", math.Min(3, 7)) // Compara 2 valores y devuelve el menor de ellos

	fmt.Println(" El valor absoluto de -15.5 es:", math.Abs(-15.5)) // Devuelve el valor absoluto de un número

	// Potencias y raíces
	fmt.Println("Raiz cuadrada de 81:", int(math.Sqrt(81)))
	fmt.Println("Raiz cubica de 81:", int(math.Cbrt(81)))
	fmt.Println("81 elevado a la 2:", int(math.Pow(81, 2)))

	// Logaritmos
	x := 100.0
	fmt.Println("Logaritmo natural de", x, ":", math.Log(x))
	fmt.Println("Logaritmo base 10 de", x, ":", math.Log10(x))
	fmt.Println("Logaritmo base 2 de", x, ":", math.Log2(x))

	// Constantes

	fmt.Println("Pi:", math.Pi)
	fmt.Println("E:", math.E)
	fmt.Println("MaxInt64:", math.MaxInt64)
	fmt.Println("MinInt64:", math.MinInt64)

	// Trigonometría
	fmt.Println("Seno de", x, ":", math.Sin(x))
	fmt.Println("Coseno de", x, ":", math.Cos(x))
	fmt.Println("Tangente de", x, ":", math.Tan(x))

	// Verificaciones especiales
	// importar inf
	pos_inf := math.Inf(1)  // infinito positivo
	neg_inf := math.Inf(-1) // infinito negativo
	fmt.Println("Infinito positivo:", pos_inf)
	fmt.Println("Infinito negativo:", neg_inf)

	fmt.Println("¿El valor", pos_inf, "es NaN?", math.IsNaN(pos_inf))
	fmt.Println("¿El valor", neg_inf, "es NaN?", math.IsNaN(neg_inf))

	fmt.Println("¿El valor", pos_inf, "es infinito positivo?", math.IsInf(pos_inf, 1))
	fmt.Println("¿El valor", pos_inf, "es infinito negativo?", math.IsInf(pos_inf, -1))
	fmt.Println("¿El valor", neg_inf, "es infinito positivo?", math.IsInf(neg_inf, 1))
	fmt.Println("¿El valor", neg_inf, "es infinito negativo?", math.IsInf(neg_inf, -1))
	fmt.Println("¿El valor", pos_inf, "es infinito de cualquier signo?", math.IsInf(pos_inf, 0))
	fmt.Println("¿El valor", neg_inf, "es infinito de cualquier signo?", math.IsInf(neg_inf, 0))

	fmt.Println("¿Es NaN?", math.IsNaN(x))    // ¿es "Not a Number"? (ej: resultado de 0.0/0.0)
	fmt.Println("¿Es Inf?", math.IsInf(x, 0)) // ¿es infinito? (positivo o negativo)

}

func get_random_number() int {
	return rand.Intn(100)
}

func get_random_float() float64 {
	return rand.Float64()
}
