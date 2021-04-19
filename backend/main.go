package main

import (
	"log"
	"net/http"
	"os"

	"github.com/blainsmith/go-svelte-rpc-starter/backend/rpc"
	"github.com/blainsmith/go-svelte-rpc-starter/backend/rpc/rpcserver"
	"github.com/pacedotdev/oto/otohttp"
)

func main() {
	port := os.Getenv("PORT")
	staticDir := os.Getenv("STATIC_DIR")

	httpServer := otohttp.NewServer()

	greeterService := rpc.GreeterService{}
	rpcserver.RegisterGreeterService(httpServer, &greeterService)

	mathService := rpc.MathService{}
	rpcserver.RegisterMathService(httpServer, &mathService)

	http.Handle("/oto/", httpServer)
	http.Handle("/", http.FileServer(http.Dir(staticDir)))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
