package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	RandomHandlerConstructor(router)

	server := http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	fmt.Println("Server is listening on port 5000")
	server.ListenAndServe()
}
