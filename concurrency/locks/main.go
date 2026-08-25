package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// En ../goroutines vimos que incrementar una variable compartida desde
// muchas goroutines da resultados inconsistentes (race condition). Acá
// vemos 3 formas de arreglarlo.

func conMutex() {
	// sync.Mutex (MUTual EXclusion): solo una goroutine a la vez puede
	// estar entre Lock() y Unlock(). Las demás se BLOQUEAN esperando.
	var mu sync.Mutex
	contador := 0

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			contador++ // sección crítica: solo una goroutine entra a la vez
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("con Mutex, contador final:", contador) // siempre 1000
}

// sync.RWMutex: distingue lectores de escritores. Muchos lectores pueden
// entrar al mismo tiempo (RLock), pero un escritor (Lock) es exclusivo y
// espera a que no haya nadie más adentro. Útil cuando hay MUCHAS lecturas
// y pocas escrituras, como una caché.
type Cache struct {
	mu   sync.RWMutex
	data map[string]int
}

func NuevaCache() *Cache {
	return &Cache{data: make(map[string]int)}
}

func (c *Cache) Get(clave string) (int, bool) {
	c.mu.RLock() // lock de lectura: no bloquea a otros lectores
	defer c.mu.RUnlock()
	v, ok := c.data[clave]
	return v, ok
}

func (c *Cache) Set(clave string, valor int) {
	c.mu.Lock() // lock de escritura: exclusivo
	defer c.mu.Unlock()
	c.data[clave] = valor
}

func conRWMutex() {
	cache := NuevaCache()
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		cache.Set("x", 42)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if v, ok := cache.Get("x"); ok {
			fmt.Println("leí x =", v)
		}
	}()

	wg.Wait()
}

func conAtomic() {
	// para operaciones simples sobre un solo número, sync/atomic es más
	// liviano que un Mutex: usa instrucciones atómicas de la CPU
	var contador atomic.Int64

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			contador.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println("con atomic, contador final:", contador.Load())
}

func conSyncOnce() {
	// sync.Once garantiza que una función se ejecute UNA sola vez,
	// sin importar cuántas goroutines la llamen. Típico para inicializar
	// un singleton de forma segura.
	var once sync.Once
	inicializar := func() {
		fmt.Println("inicializando (esto se imprime una sola vez)")
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(inicializar)
		}()
	}
	wg.Wait()
}

func main() {
	conMutex()
	conRWMutex()
	conAtomic()
	conSyncOnce()

	// Para DETECTAR condiciones de carrera automáticamente, Go trae el
	// "race detector" incorporado. Corré este mismo archivo con:
	//
	//   go run -race main.go
	//
	// y probá comentar el mu.Lock()/mu.Unlock() de conMutex para ver
	// cómo el detector lo marca en rojo con el stack trace exacto.
}
