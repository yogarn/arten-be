package config

import "os"

func LoadSmtpCredentials() map[string]string {
	credentials := make(map[string]string)
	credentials["host"] = os.Getenv("SMTP_HOST")
	credentials["port"] = os.Getenv("SMTP_PORT")
	credentials["name"] = os.Getenv("SENDER_NAME")
	credentials["email"] = os.Getenv("AUTH_EMAIL")
	credentials["password"] = os.Getenv("AUTH_PASSWORD")
	return credentials
}
