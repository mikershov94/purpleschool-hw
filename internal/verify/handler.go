package verify

import "net/http"

type VerifyHandler struct{}

func NewVerifyHander(router *http.ServeMux) {
	handler := &VerifyHandler{}

	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req http.Request) {
		w.Write({})
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req http.Request) {
		w.Write({})
	}
}
