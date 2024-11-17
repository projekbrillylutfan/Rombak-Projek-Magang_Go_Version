package utils

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendVerificationEmail(email string, token string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "your_email@example.com")
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Email Verification")
	mailer.SetBody("text/plain", fmt.Sprintf("Please verify your email using this link: http://localhost:3000/verify-email?token=%s", token))

	dialer := gomail.NewDialer("smtp.example.com", 587, "your_email@example.com", "your_password")

	return dialer.DialAndSend(mailer)
}