// gosendmail
//
// Author: Ondrej Sika <ondrej@ondrejsika.com>
// Source: https://github.com/ondrejsika/gosendmail
//

package lib

import "net/smtp"

// GoSendMail send an email from Go using SMTP
func GoSendMail(smtpHost string, smtpPort string, from string, password string, to string, subject string, message string) error {
	emailMessage := []byte("To: " + to + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		message + "\r\n")

	auth := smtp.PlainAuth(from, from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, emailMessage)
	if err != nil {
		return err
	}
	return nil
}
