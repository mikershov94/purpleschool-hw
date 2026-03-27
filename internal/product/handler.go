package product

import (
	"go/validation-api/configs"
	"go/validation-api/pkg/req"
	"go/validation-api/pkg/res"
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

	router.HandleFunc("POST /products", handler.Create())
}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductCreateRequest](&w, r)
		if err != nil {
			return
		}

		product := ProductConstructor(body.Name, body.Description, body.Image)

		createdProduct, err := handler.Repo.Create(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		res.Json(w, createdProduct, 201)
	}
}

// Update
// Delete
// GetById
