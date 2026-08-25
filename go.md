# GO

## Contenido del tutorial

Cada carpeta es un programa independiente y comentado en español. Se corre
parado en la carpeta con `go run main.go` (o `go run server.go` /
`go run client.go` en el caso de sockets, que son dos programas).

**basics/** — fundamentos del lenguaje
- `hello/` — programa mínimo
- `math/` — paquete `math`, números random
- `slices/` — arrays, slices, `append`, aliasing
- `structs/` — structs, métodos (value vs pointer receiver), embedding
- `for/` — el único loop de Go, `range`, labels
- `maps/` — `map[K]V`, comma-ok, mapas nil
- `pointers/` — `&`, `*`, paso por valor vs referencia
- `interfaces/` — interfaces implícitas, `any`, type assertions/switch, `fmt.Stringer`
- `errors/` — `error`, wrapping con `%w`, `errors.Is`/`As`, `defer`, `panic`/`recover`

**concurrency/** — concurrencia y comunicación entre procesos
- `goroutines/` — qué es una goroutine (vs. un thread del SO), `sync.WaitGroup`, condiciones de carrera
- `channels/` — comunicar en vez de compartir memoria, `select`, patrón done-channel
- `locks/` — `sync.Mutex`, `sync.RWMutex`, `sync/atomic`, `sync.Once`, `go run -race`
- `sockets/` — servidor y cliente TCP (echo server) con `net`
- `blocking-queue/` — cola bloqueante productor/consumidor: con channel vs. con `sync.Cond`

## Como correr un programa

Crear un archivo si tiene un main podemos correrlo haciendo:

```bash
go run hello.go
```
Esto lo corre sin compilar

## Compilarlo y después ejecutarlo
```fish
go build hello.go # Esto lo compila
./hello # Esto lo ejecuta
```

### 1. Inicializar un módulo (go mod init)

Cuando tu proyecto crece más allá de un archivo suelto, necesitás un módulo. Parate en la carpeta de tu proyecto y corré:

```fish
go mod init nombre-del-modulo
```

Por convención el nombre suele ser algo tipo github.com/tu-usuario/tu-proyecto (aunque no publiques nada, sirve como identificador único). Esto genera un go.mod, que es como el package.json de Node o el Cargo.toml de Rust: lleva registro de tu módulo y sus dependencias.

### 2. Estructura típica de un proyecto
```fish
miproyecto/
├── go.mod
├── main.go
├── internal/
│   └── config/
│       └── config.go
└── pkg/
    └── util/
        └── util.go
```

main.go: punto de entrada, tiene func main().
internal/: código privado de tu módulo, nadie de afuera puede importarlo (el compilador lo fuerza).
pkg/: código pensado para ser reutilizado/importado por otros.

Para proyectos chicos ni te preocupes por esto, es para cuando ya tenés varios archivos.

### 3. Paquetes (packages)

Cada carpeta es un paquete. Todo archivo .go dentro arranca con:

```fish
go
package nombrepaquete
```

package main es especial: le dice a Go que ese paquete es ejecutable (necesita func main()). Cualquier otro nombre es una librería que se importa.

## 4. Manejar dependencias externas
```fish
go get github.com/algun/paquete
```

Esto la agrega a tu go.mod y descarga el código. Después la importás normal:

```go
import "github.com/algun/paquete"
```

Para limpiar dependencias no usadas:

```fish
go mod tidy
```

## 5. Testing (viene incorporado, sin librerías externas)

Archivo algo_test.go:

```go
package main

import "testing"

func TestSuma(t *testing.T) {
    resultado := Suma(2, 3)
    if resultado != 5 {
        t.Errorf("esperaba 5, obtuve %d", resultado)
    }
}
```

Se corre con:

```fish
go test ./...
```

## 6. Formateo y linting automático

```fish
gofmt -w .
```