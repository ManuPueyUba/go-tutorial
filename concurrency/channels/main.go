package main

// Filosofía de Go: "Don't communicate by sharing memory;
// share memory by communicating."
//
// En vez de que varias goroutines lean/escriban la misma variable protegida
// por un lock, la idea es que se pasen datos entre ellas a través de un
// CHANNEL: una especie de tubería thread-safe por la que fluyen valores.

import (
	"fmt"
	"time"
)

func main() {
	// channel sin buffer (unbuffered): el que envía (send) se bloquea hasta
	// que otra goroutine reciba (receive), y viceversa. Es un "rendezvous".
	mensajes := make(chan string)

	go func() {
		mensajes <- "hola desde la goroutine" // enviar: bloquea hasta que alguien reciba
	}()
	msg := <-mensajes // recibir: bloquea hasta que llegue algo
	fmt.Println(msg)

	// channel CON buffer: enviar no bloquea hasta que el buffer se llena
	buffer := make(chan int, 3)
	buffer <- 1
	buffer <- 2
	buffer <- 3
	// buffer <- 4 // esto bloquearía: el buffer está lleno (capacidad 3)
	fmt.Println(<-buffer, <-buffer, <-buffer)

	// dirección de un channel: se puede restringir a "solo enviar" o
	// "solo recibir" en la firma de una función, para que el compilador
	// nos cuide de usarlo mal
	generar := func(out chan<- int, n int) { // out: solo se puede enviar
		for i := 0; i < n; i++ {
			out <- i
		}
		close(out) // close() avisa "no voy a mandar nada más"
	}
	consumir := func(in <-chan int) { // in: solo se puede recibir
		// range sobre un channel recibe valores hasta que se cierra
		for v := range in {
			fmt.Println("recibido:", v)
		}
	}
	numeros := make(chan int)
	go generar(numeros, 5)
	consumir(numeros)

	// recibir con comma-ok: permite distinguir "el channel está cerrado"
	// de "recibí el zero value"
	cerrado := make(chan int)
	close(cerrado)
	v, ok := <-cerrado
	fmt.Println(v, ok) // 0 false -> channel cerrado y vacío

	// select: espera sobre VARIOS channels a la vez, sigue con el primero
	// que esté listo (como un switch, pero para operaciones de channel)
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(50 * time.Millisecond)
		c1 <- "resultado de c1"
	}()
	go func() {
		time.Sleep(20 * time.Millisecond)
		c2 <- "resultado de c2"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

	// select con timeout: patrón muy común para no bloquearse para siempre
	lento := make(chan string)
	select {
	case res := <-lento:
		fmt.Println(res)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timeout: nadie respondió a tiempo")
	}

	// patrón "done channel": una goroutine avisa a otra que pare,
	// cerrando un channel (varias goroutines pueden recibir de un
	// channel cerrado sin bloquearse, todas reciben el zero value)
	done := make(chan struct{}) // struct{} no ocupa memoria, se usa como señal pura
	go func() {
		fmt.Println("worker: trabajando...")
		time.Sleep(30 * time.Millisecond)
		fmt.Println("worker: listo, avisando que terminé")
		close(done)
	}()
	<-done
	fmt.Println("main: el worker terminó")
}
