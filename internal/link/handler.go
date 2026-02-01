package link

import (
	"go/adv-demo/configs"
	"net/http"
)

type LinkHandlerDeps struct {
	// Config *configs.Config
}

type LinkHandler struct {
	Config *configs.Config
}

func LinkHandlerConstructor(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		// Config: deps.Config,
	}

	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /link/{alias}", handler.Read())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) Read() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}