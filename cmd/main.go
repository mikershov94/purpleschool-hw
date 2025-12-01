package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("Server starting on port 3000")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %v\r\n", err)
		return
	}
}
