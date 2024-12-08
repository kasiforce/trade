package util

import (
	"fmt"
	conf "github.com/kasiforce/trade/config"
	"golang.org/x/exp/rand"
	"gopkg.in/mail.v2"
	"mime"
	"strconv"
	"time"
)

type EmailSender struct {
	SmtpHost      string `json:"smtp_host"`
	SmtpEmailFrom string `json:"smtp_email_from"`
	SmtpPass      string `json:"smtp_pass"`
}

func NewEmailSender() *EmailSender {
	eConfig := conf.Config.Email
	return &EmailSender{
		SmtpHost:      eConfig.SmtpHost,
		SmtpEmailFrom: eConfig.SmtpEmail,
		SmtpPass:      eConfig.SmtpPass,
	}
}

// Send 发送邮件
func (s *EmailSender) Send(data, emailTo, subject string) error {
	m := mail.NewMessage()
	encodedName := mime.QEncoding.Encode("UTF-8", "校园交易平台")
	m.SetHeader("From", fmt.Sprintf("%s <%s>", encodedName, s.SmtpEmailFrom))
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", data)
	d := mail.NewDialer(s.SmtpHost, 465, s.SmtpEmailFrom, s.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func GenerateEmailCode() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	return strconv.Itoa(rand.Intn(1000000))
}
