package utils

import (
	"fmt"
	"log"

	"github.com/bete7512/telegram-cms/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendSignupEmail(firstName string, email string) error {
	from := mail.NewEmail("Telegram CMS", "no-reply@hubbits.co")
	subject := "Signup Verification"
	to := mail.NewEmail(firstName, email)
	plainTextContent := "click the link below to verify your account"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.SENDGRID_API_KEY)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
		if response.StatusCode != 202 {
			return fmt.Errorf("failed to send email")
		}
	}
	return nil
}
