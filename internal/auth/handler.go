package auth

import (
	"encoding/json"
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/res"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AuthHandlerDeps struct {
	Config *configs.Config
}

type AuthHandler struct {
	Config *configs.Config
}

func AuthHandlerConstructor(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("/auth/login", handler.Login())
	router.HandleFunc("/auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			res.Json(w, err.Error(), 402)
			return
		}

		validate := validator.New()
		err = validate.Struct(payload)
		if err != nil {
			res.Json(w, err.Error(), 402)
			return
		}

		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("register")
		return
	}
}
