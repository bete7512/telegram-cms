package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"

	"github.com/bete7512/telegram-cms/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendSignupEmail(firstName string, email string, redirectUri string) error {
    from := mail.NewEmail("Telegram CMS", "betekibebe@gmail.com")
    subject := "Signup Verification"
    to := mail.NewEmail(firstName, email)

    htmlFile, err := ioutil.ReadFile("./statics/signup_verification.html")
    if err != nil {
        return fmt.Errorf("failed to read HTML file: %w", err)
    }

    t, err := template.New("signup_verification").Parse(string(htmlFile))
    if err != nil {
        return fmt.Errorf("failed to parse HTML template: %w", err)
    }

    data := struct {
        FirstName   string
        RedirectUri string
    }{
        FirstName:   firstName,
        RedirectUri: redirectUri,
    }
	log.Println(data.RedirectUri,"jjjjjjjjjjjjjjj")
    var htmlContent bytes.Buffer
    if err := t.Execute(&htmlContent, data); err != nil {
        return fmt.Errorf("failed to execute HTML template: %w", err)
    }

    plainTextContent := "Click the link below to verify your account"
    message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent.String())
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


func SendForgetPasswordEmail(firstName string, email string, resetLink string) error {
	from := mail.NewEmail("Telegram CMS", "betekibebe@gmail.com")
	subject := "Forget Password"
	to := mail.NewEmail(firstName, email)
	// Read the HTML template file
	htmlFile, err := ioutil.ReadFile("./statics/forget_password_template.html")
	if err != nil {
		return fmt.Errorf("failed to read HTML file: %w", err)
	}

	// Parse the HTML template
	t, err := template.New("forget_password").Parse(string(htmlFile))
	if err != nil {
		return fmt.Errorf("failed to parse HTML template: %w", err)
	}

	// Data to be passed to the template
	data := struct {
		FirstName string
		ResetLink string
	}{
		FirstName: firstName,
		ResetLink: resetLink,
	}

	// Execute the template with the data
	var htmlContent bytes.Buffer
	if err := t.Execute(&htmlContent, data); err != nil {
		return fmt.Errorf("failed to execute HTML template: %w", err)
	}

	plainTextContent := "Click the link below to reset your password"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent.String())
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
	}	// Data to be passed to the template

	return nil
}
