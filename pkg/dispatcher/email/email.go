package email

import (
	"net/smtp"
)

type email struct {
	username  string
	password  string
	host      string
	recipient string
}

func (e *email) Send(message string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		e.username,
		e.password,
		e.host,
	)

	err := smtp.SendMail(
		e.host+":25",
		auth,
		e.username,
		[]string{e.recipient},
		[]byte(message),
	)

	return err
}
