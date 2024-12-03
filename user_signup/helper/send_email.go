package helper

import (
	"log"
	"net/smtp"
)

func SendEmail(to string, subject string, body string) error {
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	form := "divyanshusingh7060@gmail.com"
	password := "Divya@2003"
	auth := smtp.PlainAuth("", form, password, smtpServer)
	msg := "To: " + to + "\r\nSubject: " + subject + "\r\n\r\n" + body
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, form, []string{to}, []byte(msg))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
