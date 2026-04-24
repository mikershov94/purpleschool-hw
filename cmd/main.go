package main

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/link"
	"go/adv-demo/pkg/db"
	"go/adv-demo/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.DbConstructor(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepo := link.LinkRepositoryConstructor(db)

	// Handlers
	auth.AuthHandlerConstructor(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	link.LinkHandlerConstructor(router, link.LinkHandlerDeps{
		Repo: linkRepo,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8080")
	server.ListenAndServe()
}
