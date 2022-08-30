package email

import (
	"errors"

	"github.com/aecra/covid/common"
	"gopkg.in/gomail.v2"
)

type Message struct {
	Title   string `json:"title" form:"title"`
	Email   string `json:"email" form:"email"`
	Content string `json:"content" form:"content"`
}

type SendConfig struct {
	Username  string `json:"username" form:"username"`
	EmailAddr string `json:"email_add" form:"email_add"`
	Password  string `json:"password" form:"password"`
	Host      string `json:"host" form:"host"`
	Port      int    `json:"port" form:"port"`
}

func sendEmail(message *Message, sendConfig *SendConfig) error {
	// verify message
	if message.Title == "" {
		return errors.New("title is empty")
	}
	if !common.VerifyEmail(message.Email) {
		return errors.New("email is invalid")
	}
	if message.Content == "" {
		return errors.New("content is empty")
	}
	// verify send config
	if sendConfig.Username == "" {
		return errors.New("server config: username is empty")
	}
	if !common.VerifyEmail(sendConfig.EmailAddr) {
		return errors.New("server config: email address is invalid")
	}
	if sendConfig.Password == "" {
		return errors.New("server config: password is empty")
	}
	if !common.VerifyHost(sendConfig.Host) {
		return errors.New("server config: host is empty")
	}
	if sendConfig.Port == 0 {
		return errors.New("server config: port is empty")
	}
	return send(message, sendConfig)
}

func send(message *Message, sendConfig *SendConfig) error {
	// send email
	m := gomail.NewMessage()
	m.SetHeader("From", sendConfig.Username+"<"+sendConfig.EmailAddr+">")
	m.SetHeader("To", message.Email)
	m.SetHeader("Subject", message.Title)
	m.SetBody("text/html", message.Content)

	d := gomail.NewDialer(
		sendConfig.Host,
		sendConfig.Port,
		sendConfig.EmailAddr,
		sendConfig.Password,
	)

	return d.DialAndSend(m)
}
