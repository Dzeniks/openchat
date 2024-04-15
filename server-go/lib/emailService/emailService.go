package emailService

import (
	"net/smtp"
	env "server-go/lib/dotEnv"
)

func SendEmail(to []string, message []byte) error {
	// Sender data.
	from := env.DotEnv.SenderEmail
	password := env.DotEnv.SenderPassword

	// smtp server configuration.
	smtpHost := env.DotEnv.SmtpHost
	smtpPort := env.DotEnv.SmtpPort

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}
	return nil
}
