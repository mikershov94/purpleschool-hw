package verify

type VerifyData struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

type SendRequest struct {
	Email string `json:"email" validate:"email"`
}