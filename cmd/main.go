package main

import (
	"fmt"
	"go/validation-api/configs"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	router := http.NewServeMux()

	port := conf.Server.Port
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	if port == 0 {
		fmt.Println("Server failed: uncorrect port")
		return
	}
	fmt.Println("Server starting on port", port)

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %v\r\n", err)
		return
	}
}
