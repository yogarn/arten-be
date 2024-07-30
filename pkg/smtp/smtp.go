package smtp

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/yogarn/arten/pkg/config"
)

type Interface interface {
	SendMail(to []string, subject, message string) error
}

type smtpClient struct{}

func Init() Interface {
	return &smtpClient{}
}

func (smtpClient *smtpClient) SendMail(to []string, subject, message string) error {
	credentials := config.LoadSmtpCredentials()

	body := "From: " + credentials["name"] + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", credentials["email"], credentials["password"], credentials["host"])
	smtpAddr := fmt.Sprintf("%s:%s", credentials["host"], credentials["port"])

	err := smtp.SendMail(smtpAddr, auth, credentials["email"], to, []byte(body))
	if err != nil {
		return err
	}

	return nil
}
