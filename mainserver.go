// Ya definido el servicio, crearemos el servidor RPC
// y registraremos el servicio usando HTTP como protocolo
// para que este escuchando las peticiones todo el tiempo.

package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"https://github.com/JCSleiman/rpc-server/tree/master/Service"
)

func main() {
	s := new(service.Areas)
	err := rpc.Register(s)
	if err != nil {
		log.Fatal("Service registration error:", err)
	}

	rpc.HandleHTTP()

	l, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		lof.Fatal("Listen error:", err)
	}

	log.Println("Servidor iniciado!")
	log.Println(l.Addr())
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
