package main

import (
	"fmt"
	"go/validation-api/configs"
	"go/validation-api/internal/verify"
	"go/validation-api/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.DbConstructor(conf)
	router := http.NewServeMux()
	verify.NewVerifyHandler(router, conf.Email)

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
