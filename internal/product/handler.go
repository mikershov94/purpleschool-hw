package product

import (
	"go/validation-api/configs"
	"go/validation-api/pkg/req"
	"go/validation-api/pkg/res"
	"net/http"
	"strconv"

	"gorm.io/gorm"
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
	router.HandleFunc("PATCH /products/{id}", handler.Update())
	router.HandleFunc("DELETE /products/{id}", handler.Delete())
	router.HandleFunc("GET /products/{id}", handler.GetById())
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

func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductUpdateRequest](&w, r)
		if err != nil {
			return
		}

		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		product, err := handler.Repo.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        body.Name,
			Description: body.Description,
			Image:       body.Image,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, product, 201)
	}
}

func (handler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = handler.Repo.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = handler.Repo.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(w, nil, 200)
	}
}

func (handler *ProductHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		product, err := handler.Repo.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		res.Json(w, product, 200)
	}
}
