package main

// ¿QUÉ ES UNA GOROUTINE?
//
// Un THREAD (hilo) del sistema operativo es una unidad de ejecución que el
// SO planifica sobre los cores de la CPU. Crear uno es "caro": ocupa varios
// MB de stack y cambiar de contexto entre threads involucra al kernel.
//
// Una GOROUTINE es una función que corre de forma concurrente, gestionada
// por el RUNTIME de Go, no directamente por el sistema operativo. Es mucho
// más liviana que un thread: arranca con un stack de ~2KB (que crece según
// hace falta), y el runtime multiplexa muchísimas goroutines sobre un
// número chico de threads del SO reales. Este esquema se llama M:N
// scheduling (M goroutines corriendo sobre N threads del SO).
//
// Por eso es normal en Go tener miles o millones de goroutines vivas al
// mismo tiempo, algo impensado con threads del SO tradicionales.
//
// La variable de entorno / runtime.GOMAXPROCS controla cuántos threads del
// SO usa el runtime para correr goroutines EN PARALELO (por defecto, uno
// por core lógico). Concurrencia (varias tareas progresando, intercaladas)
// no es lo mismo que paralelismo (varias tareas corriendo literalmente al
// mismo tiempo en distintos cores) — Go te da ambas cosas con goroutines.

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func saludar(nombre string) {
	fmt.Println("Hola desde", nombre)
}

func main() {
	fmt.Println("Cores lógicos disponibles (GOMAXPROCS):", runtime.GOMAXPROCS(0))

	// "go" antes de una llamada la lanza como goroutine: NO espera a que termine
	go saludar("goroutine 1")
	go saludar("goroutine 2")

	// si main() termina, el programa termina YA, aunque haya goroutines
	// corriendo. Por eso este sleep (mala práctica, es solo para demostrar):
	time.Sleep(100 * time.Millisecond)
	fmt.Println("--- fin del ejemplo con sleep (no lo hagas así en código real) ---")

	// la forma correcta de esperar goroutines es sync.WaitGroup
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) // avisamos que hay una goroutine más por esperar
		go func(id int) {
			defer wg.Done() // avisa que esta goroutine terminó
			fmt.Println("worker", id, "trabajando")
		}(i) // OJO: pasamos "i" como argumento, no lo capturamos por closure
	}
	wg.Wait() // bloquea hasta que todos los wg.Done() se hayan llamado
	fmt.Println("todos los workers terminaron")

	// ¿Por qué pasamos "i" como argumento en vez de usarlo directo en el closure?
	// Si el closure capturara "i" directamente, todas las goroutines
	// compartirían la MISMA variable "i", y para cuando se ejecuten podría
	// valer cualquier cosa (típicamente el último valor). Pasarlo como
	// parámetro crea una copia propia para cada goroutine.

	// CONDICIÓN DE CARRERA (race condition): dos o más goroutines acceden
	// a la misma variable al mismo tiempo, y al menos una la escribe.
	// El resultado es no determinístico. Este contador está roto a propósito:
	contador := 0
	var wg2 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			contador++ // NO es atómico: leer, sumar y escribir son 3 pasos
		}()
	}
	wg2.Wait()
	fmt.Println("contador final (probablemente NO sea 1000):", contador)
	// Corré esto varias veces, o con: go run -race main.go
	// para que Go detecte la carrera automáticamente.
	// La solución (sync.Mutex, channels, atomic) está en ../locks y ../channels
}
