package main

// COLA BLOQUEANTE (blocking queue): una cola con capacidad limitada donde
// - Put() bloquea si la cola está LLENA, hasta que alguien saque un elemento
// - Get() bloquea si la cola está VACÍA, hasta que alguien meta un elemento
//
// Es la estructura clásica para el patrón productor/consumidor.
// Acá van DOS implementaciones para comparar dos formas de pensar
// concurrencia en Go: la idiomática (channels) y la manual (mutex + cond,
// muy parecida a como se haría con threads en Java/C++).

import (
	"fmt"
	"sync"
)

// ---------- Implementación 1: con un channel con buffer ----------
// Esta es la forma IDIOMÁTICA en Go. Un channel con buffer ES,
// básicamente, una cola bloqueante ya hecha: el propio runtime se
// encarga de bloquear al que envía si está llena, y al que recibe
// si está vacía. No hace falta reinventar nada.

type ColaChannel struct {
	items chan int
}

func NuevaColaChannel(capacidad int) *ColaChannel {
	return &ColaChannel{items: make(chan int, capacidad)}
}

func (c *ColaChannel) Put(v int) {
	c.items <- v // bloquea automáticamente si el buffer está lleno
}

func (c *ColaChannel) Get() int {
	return <-c.items // bloquea automáticamente si está vacío
}

func (c *ColaChannel) Cerrar() {
	close(c.items)
}

// ---------- Implementación 2: con Mutex + sync.Cond ----------
// Esta es la forma "manual", más parecida a cómo se implementaría en un
// lenguaje con threads e hilos del sistema operativo desnudos. Sirve para
// entender QUÉ hace un channel por dentro, y es útil cuando necesitás
// lógica de sincronización más compleja que un channel no te da directo.
//
// sync.Cond permite que una goroutine se "duerma" (Wait) liberando el
// lock, hasta que otra la despierte (Signal/Broadcast) tras recuperar el
// lock. Es el mecanismo clásico de "condition variable".

type ColaCond struct {
	mu        sync.Mutex
	cond      *sync.Cond
	items     []int
	capacidad int
}

func NuevaColaCond(capacidad int) *ColaCond {
	c := &ColaCond{capacidad: capacidad}
	c.cond = sync.NewCond(&c.mu)
	return c
}

func (c *ColaCond) Put(v int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// "for", no "if": al despertar hay que re-chequear la condición,
	// porque puede que otra goroutine haya llenado la cola de nuevo
	// entre que nos despertaron y que recuperamos el lock
	for len(c.items) == c.capacidad {
		c.cond.Wait() // libera el lock, duerme, y lo re-adquiere al despertar
	}
	c.items = append(c.items, v)
	c.cond.Signal() // despierta a UN goroutine esperando (típicamente un Get)
}

func (c *ColaCond) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	for len(c.items) == 0 {
		c.cond.Wait()
	}
	v := c.items[0]
	c.items = c.items[1:]
	c.cond.Signal() // despierta a UN goroutine esperando (típicamente un Put)
	return v
}

// ---------- demo: productor / consumidor ----------

func demoConChannel() {
	fmt.Println("--- demo con ColaChannel ---")
	cola := NuevaColaChannel(3)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() { // productor
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			fmt.Println("produciendo:", i)
			cola.Put(i)
		}
	}()

	wg.Add(1)
	go func() { // consumidor
		defer wg.Done()
		for i := 0; i < 5; i++ {
			v := cola.Get()
			fmt.Println("consumiendo:", v)
		}
	}()

	wg.Wait()
}

func demoConCond() {
	fmt.Println("--- demo con ColaCond ---")
	cola := NuevaColaCond(3)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			fmt.Println("produciendo:", i)
			cola.Put(i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			v := cola.Get()
			fmt.Println("consumiendo:", v)
		}
	}()

	wg.Wait()
}

func main() {
	demoConChannel()
	demoConCond()
}
