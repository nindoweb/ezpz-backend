package notification

import (
	"fmt"
	"net/smtp"

	"github.com/spf13/viper"
)

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func SendMail(subject string, to []string, message string) error {
	SMTP := fmt.Sprintf("%s:%d", viper.GetString("MAIL_HOST"), viper.GetInt("MAIL_PORT"))
	return smtp.SendMail(SMTP, smtp.PlainAuth("", 
		viper.GetString("MAIL_USERNAME"),
		viper.GetString("MAIL_PASSWORD"), 
		viper.GetString("MAIL_HOST")), 
		viper.GetString("MAIL_FROM"), to, []byte(message))
}
