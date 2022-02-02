package logger

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

type Mail struct {
	ToName   string
	ToAddr   string
	FromName string
	FromAddr string
	Subject  string
	Body     string
}

func New(clientName string, body string) Mail {
	return Mail{
		ToAddr:   os.Getenv("ERR_ADDRESS"),
		FromName: "Webinar API",
		FromAddr: os.Getenv("FROM_ADDRESS"),
		Subject:  "Error reporting system for client : " + clientName,
		Body:     body,
	}
}

func (mail Mail) Send() error {

	log.Println("The following receivers will be notified shortly : " + strings.Replace(os.Getenv("ERR_ADDRESS"), ",", ", ", -1))

	body := []byte("From: " + mail.FromName + " <" + mail.FromAddr + ">" + "\r\n" +
		"To: " + mail.ToAddr + "\r\n" +
		"Subject: " + mail.Subject + "\r\n\r\n" +
		mail.Body + "\r\n")

	auth := smtp.PlainAuth("", os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_ENDPOINT"))

	smtpEndpoint := fmt.Sprintf("%s:%s", os.Getenv("SMTP_ENDPOINT"), os.Getenv("SMTP_PORT"))
	return smtp.SendMail(smtpEndpoint, auth, mail.FromAddr, strings.Split(mail.ToAddr, ","), body)

}
