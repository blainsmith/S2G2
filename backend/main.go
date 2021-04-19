package main

import (
	"log"
	"net/http"

	"github.com/blainsmith/go-svelte-rpc-starter/backend/rpc"
	"github.com/blainsmith/go-svelte-rpc-starter/backend/rpc/rpcserver"
	"github.com/pacedotdev/oto/otohttp"
)

func main() {
	httpServer := otohttp.NewServer()

	greeterService := rpc.GreeterService{}
	rpcserver.RegisterGreeterService(httpServer, &greeterService)

	mathService := rpc.MathService{}
	rpcserver.RegisterMathService(httpServer, &mathService)

	http.Handle("/oto/", httpServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
