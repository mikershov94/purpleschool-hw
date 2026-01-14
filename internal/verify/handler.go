package verify

import (
	"crypto/sha256"
	"fmt"
	"go/validation-api/configs"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type VerifyHandler struct {
	Config configs.EmailConfig
}

func NewVerifyHandler(router *http.ServeMux, deps configs.EmailConfig) {
	handler := &VerifyHandler{
		Config: deps,
	}

	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(201)

		hash := sha256.Sum256([]byte(handler.Config.Email))
		// fmt.Println()
		// TODO
		// хэш перевести в строку и вставить в HTML

		e := email.NewEmail()
		e.From = "<" + handler.Config.Email + ">"
		e.To = []string{"<" + handler.Config.Email + ">"}
		e.Subject = "Michael Ershov"
		e.Text = []byte("Text Body is, of course, supported!")
		e.HTML = []byte(fmt.Sprintf(`<a href="http://localhost:3000/verify/%x">http://localhost:3000/verify/</a>`, hash))
		err := e.Send(handler.Config.Address + ":587", smtp.PlainAuth("", handler.Config.Email, handler.Config.Password, handler.Config.Address))
		if err != nil {
			fmt.Println("Не удалось отправить письмо")
			fmt.Println(err)
			return
		}

		fmt.Println("Письмо отправлено")
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		fmt.Println("Перешел по ссылке")
	}
}
