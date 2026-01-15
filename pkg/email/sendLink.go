package email

import (
	"fmt"
	"go/validation-api/configs"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendLink(config configs.EmailConfig, hash string) error {
	e := email.NewEmail()
	e.From = "<" + config.Email + ">"
	e.To = []string{"<" + config.Email + ">"}
	e.Subject = "Michael Ershov"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte(fmt.Sprintf(`<a href="http://localhost:3000/verify/%s">http://localhost:3000/verify/%s</a>`, hash, hash))

	err := e.Send(config.Address + ":587", smtp.PlainAuth("", config.Email, config.Password, config.Address))

	return err
}