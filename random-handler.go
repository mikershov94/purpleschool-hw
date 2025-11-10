package main

import (
	"fmt"
	"net/http"
)

type RandomHandler struct{}

func RandomHandlerConstructor(router *http.ServeMux) {
	handler := &RandomHandler{}
	router.HandleFunc("/random", handler.Random())
}

func (handler *RandomHandler) Random() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("random")
	}
}
