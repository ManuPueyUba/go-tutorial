package main

// Cliente TCP: se conecta al servidor de server.go, manda lo que
// escribas por stdin, y muestra lo que el servidor responde.
//
// Necesita el servidor corriendo: go run main.go en sockets/server/ (en otra terminal)
// Después, desde esta carpeta: go run main.go

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// net.Dial abre una conexión TCP hacia el servidor
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("no se pudo conectar:", err)
		return
	}
	defer conn.Close()
	fmt.Println("conectado al servidor, escribí algo (o 'salir' para cortar):")

	// leemos las respuestas del servidor en una goroutine aparte,
	// para poder seguir leyendo lo que el usuario tipea al mismo tiempo.
	// readerDone se cierra cuando el scanner corta (el servidor cerró
	// la conexión, ej: después de "chau!")
	readerDone := make(chan struct{})
	go func() {
		defer close(readerDone)
		lector := bufio.NewScanner(conn)
		for lector.Scan() {
			fmt.Println("servidor dice:", lector.Text())
		}
		if err := lector.Err(); err != nil {
			fmt.Println("error leyendo del servidor:", err)
		}
	}()

	entrada := bufio.NewScanner(os.Stdin)
	for entrada.Scan() {
		linea := entrada.Text()
		fmt.Fprintln(conn, linea) // mandar la línea al servidor
		if linea == "salir" {
			break
		}
	}
	if err := entrada.Err(); err != nil {
		fmt.Println("error leyendo stdin:", err)
	}

	// esperamos un toque a que llegue la última respuesta del servidor
	// antes de que main() termine y el defer cierre la conexión
	select {
	case <-readerDone:
	case <-time.After(300 * time.Millisecond):
	}
}
