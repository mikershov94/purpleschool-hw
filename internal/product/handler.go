package product

import (
	"go/validation-api/configs"
	"net/http"
)

type ProductHandlerDeps struct {
	Config *configs.Config
	Repo   *ProductRepository
}

type ProductHandler struct {
	Config *configs.Config
	Repo   *ProductRepository
}

func ProductHandlerConstructor(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{
		Config: deps.Config,
		Repo:   deps.Repo,
	}

	router.HandleFunc("POST /product/create", handler.Create())
}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

// Update
// Delete
// GetById
