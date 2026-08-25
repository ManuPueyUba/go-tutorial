package main

// SOCKETS: un socket es el "enchufe" que usa el sistema operativo para
// comunicar dos procesos a través de la red (o incluso en la misma
// máquina). TCP da una conexión confiable y ordenada tipo "stream de
// bytes" entre dos puntos; UDP manda paquetes sueltos, más rápido pero
// sin garantías de orden ni de entrega.
//
// Este archivo es un servidor TCP "echo": acepta conexiones y devuelve
// en mayúsculas lo que cada cliente le manda, línea por línea.
//
// Corré esto en una terminal:  go run server.go
// Y en otra el cliente:        go run client.go

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// net.Listen abre un socket y se pone a escuchar conexiones entrantes
	// en el puerto 9000 de todas las interfaces (":9000")
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("error escuchando:", err)
		return
	}
	defer listener.Close()
	fmt.Println("servidor escuchando en :9000")

	for {
		// Accept() bloquea hasta que llega una nueva conexión
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error aceptando conexión:", err)
			continue
		}
		// cada conexión se maneja en su propia goroutine: así el
		// servidor puede atender a MUCHOS clientes a la vez sin que
		// uno lento bloquee a los demás
		go manejarConexion(conn)
	}
}

func manejarConexion(conn net.Conn) {
	defer conn.Close()
	remoto := conn.RemoteAddr()
	fmt.Println("nueva conexión de", remoto)

	// bufio.Scanner lee línea por línea de la conexión (que implementa io.Reader)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println("recibido de", remoto, ":", linea)

		if linea == "salir" {
			fmt.Fprintln(conn, "chau!")
			return
		}

		respuesta := strings.ToUpper(linea)
		fmt.Fprintln(conn, respuesta) // escribir en la conexión = mandarle datos al cliente
	}
	fmt.Println("conexión cerrada por", remoto)
}
