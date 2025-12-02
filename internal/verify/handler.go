package verify

import (
	"go/validation-api/configs"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type VerifyHandler struct {
	Config configs.EmailConfig
}

func NewVerifyHander(router *http.ServeMux, deps configs.EmailConfig) {
	handler := &VerifyHandler{
		Config: deps,
	}

	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(201)

		e := email.NewEmail()
		e.From = "Michael Ershov"
		e.To = []string{handler.Config.Email}
		e.Subject = "Michael Ershov"
		e.Text = []byte("Text Body is, of course, supported!")
		e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
		e.Send("smtp.gmail.com:587", smtp.PlainAuth("", handler.Config.Email, handler.Config.Password, handler.Config.Address))

	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
	}
}
