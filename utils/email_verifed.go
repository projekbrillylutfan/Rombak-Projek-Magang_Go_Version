package utils

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

// pengen pake email ini


func SendVerificationEmail(email string, token string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "mencobaakun531@gmail.com")
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Email Verification")
	mailer.SetBody("text/plain", fmt.Sprintf("Please verify your email using this link: http://localhost:3000/verify-email?token=%s", token))

	dialer := gomail.NewDialer("smtp.example.com", 587, "mencobaakun531@gmail.com", "Ayam.KauDuaRibu534")

	return dialer.DialAndSend(mailer)
}