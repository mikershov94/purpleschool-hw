package main

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/link"
	"go/adv-demo/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.DbConstructor(conf)
	router := http.NewServeMux()

	// Handlers
	auth.AuthHandlerConstructor(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	link.LinkHandlerConstructor(router, link.LinkHandlerDeps{})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8080")
	server.ListenAndServe()
}
