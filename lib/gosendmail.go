// gosendmail
//
// Author: Ondrej Sika <ondrej@ondrejsika.com>
// Source: https://github.com/ondrejsika/gosendmail
//

package lib

import "net/smtp"

// GoRawSendMail send a raw from Go using SMTP
func GoRawSendMail(smtpHost string, smtpPort string, user string, password string, from string, to string, rawMessage []byte) error {
	var auth smtp.Auth
	if password != "" {
		auth = smtp.PlainAuth(from, user, password, smtpHost)
	}

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, rawMessage)
	if err != nil {
		return err
	}
	return nil
}

// GoSendMail send an email from Go using SMTP
func GoSendMail(smtpHost string, smtpPort string, user string, password string, from string, to string, subject string, message string) error {
	rawMessage := []byte("To: " + to + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		message + "\r\n")
	return GoRawSendMail(smtpHost, smtpPort, user, password, from, to, rawMessage)
}
