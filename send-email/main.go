package main

import (
	"fmt"
	"net/smtp"
)

func main() {
  // SMTP configuration
  username := "2bbbff83ca70ad"
  password := "6dd01eef01fb01"
  host := "smtp.mailtrap.io"
  port := "25"

  // Subject and body
  subject := "Simple HTML Email"
  body := "Here is a simple plain text."

  // Sender and receiver
  from := "from@example.com"
  to := []string{
    "to@example.com",
  }

  // Build the message
  message := fmt.Sprintf("From: %s\r\n", from)
  message += fmt.Sprintf("To: %s\r\n", to)
  message += fmt.Sprintf("Subject: %s\r\n", subject)
  message += fmt.Sprintf("\r\n%s\r\n", body)

  // Authentication.
  auth := smtp.PlainAuth("", username, password, host)

  // Send email
  err := smtp.SendMail(host+":"+port, auth, from, to, []byte(message))
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println("Email sent successfully.")
}