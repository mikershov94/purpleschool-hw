package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type RandomHandler struct{}

func RandomHandlerConstructor(router *http.ServeMux) {
	handler := &RandomHandler{}
	router.HandleFunc("/random", handler.Random())
}

func (handler *RandomHandler) Random() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		num := rand.Intn(6) + 1
		strInt := fmt.Sprintf("%d", num)

		w.Write([]byte(strInt))
		return
	}
}
