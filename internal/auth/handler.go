package auth

import (
	"fmt"
	"go/adv-demo/configs"
	"net/http"
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

	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("register")
		return
	}
}
