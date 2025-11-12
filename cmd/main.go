package main

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.AuthHandlerConstructor(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	fmt.Println("Server is listening on port 5000")
	server.ListenAndServe()
}
