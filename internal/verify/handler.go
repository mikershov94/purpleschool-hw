package verify

import (
	"encoding/json"
	"fmt"
	"go/validation-api/configs"
	"go/validation-api/pkg/email"
	"go/validation-api/pkg/req"
	"go/validation-api/pkg/res"
	"net/http"
	"os"
	"strings"
)

type VerifyHandler struct {
	Config configs.EmailConfig
}

func VerifyHandlerConstructor(router *http.ServeMux, deps configs.EmailConfig) {
	handler := &VerifyHandler{
		Config: deps,
	}

	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[SendRequest](&w, r)
		if err != nil {
			return
		}

		hash, err := email.Hash(body.Email)
		if err != nil {
			fmt.Println("Ошибка хэширования")
			return
		}

		err = email.SendLink(handler.Config, body.Email, hash)
		if err != nil {
			fmt.Println("Не удалось отправить письмо")
			fmt.Println(err)
			return
		}

		data := VerifyData{
			Email: body.Email,
			Hash:  hash,
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Ошибка при маршалинге данных")
			return
		}
		err = os.WriteFile("verify_data.json", jsonData, 0644)
		if err != nil {
			fmt.Println("Ошибка при создании файла")
			return
		}

		w.WriteHeader(201)
		fmt.Println("Письмо отправлено")
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		filename := "verify_data.json"
		verifyData, err := ReadVerifyData(filename)
		if err != nil {
			fmt.Println("Ошибка при чтении файла:", err)
			res.Json(w, false, 503)
			return
		}

		target := strings.Replace(req.URL.Path, "/verify/", "", 1)

		if !email.HashIsValid(verifyData.Hash, target) {
			fmt.Println("Ошибка верификации")
			res.Json(w, false, 402)
			return
		}

		_, err = os.Stat("verify_data.json")
		if err == nil {
			err := os.Remove("verify_data.json")
			if err != nil {
				fmt.Println("Ошибка при удалении файла:", err)
				res.Json(w, false, 503)
				return
			}
		}

		w.WriteHeader(200)
		fmt.Println("Перешел по ссылке")
		res.Json(w, true, 200)
	}
}
