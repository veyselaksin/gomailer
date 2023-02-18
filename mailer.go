package main

// import (
// 	"gomailer/pkg/message"
// 	"net/smtp"
// )

// // Authentication is a struct that contains the authentication information for the SMTP server.
// // It is used to create a new Sender. The fields are:
// //
// //	Username: The username for the SMTP server.
// //	Password: The password for the SMTP server.
// //	Host: The host for the SMTP server.
// //	Port: The port for the SMTP server.
// type Authentication struct {
// 	Username string
// 	Password string
// 	Host     string
// 	Port     string
// }

// type ISender interface {
// 	PlainAuth(auth *Authentication) *Sender
// 	Send(m *message.Message) error
// }

// type AuthType string

// // Sender is a struct that contains the authentication information for the SMTP server.
// type Sender struct {
// 	AuthType  AuthType
// 	PlainAuth *smtp.Auth
// }

// var _ ISender = (*Sender)(nil)

// // New is a method that creates a new Sender with the given authentication information.
// func New(authType AuthType) *Sender {
// 	return &Sender{AuthType: authType}
// }

// // func PlainAuth(auth *Authentication) *Sender {
// // 	auth_ := smtp.PlainAuth("", auth.Username, auth.Password, auth.Host)
// // 	return &Sender{PlainAuth: &auth_}
// // }

// // func New(authType AuthType) *Sender {
// // 	if authType == "plain" {
// // 		auth := smtp.PlainAuth("", username, password, host)
// // 		return &Sender{authType}
// // 	}

// // 	return &Sender{authType}
// // }

// // func New() *Sender {
// // 	auth := smtp.PlainAuth("", username, password, host)
// // 	return &Sender{auth}
// // }

// // func (s *Sender) Send(m *Message) error {
// // 	return smtp.SendMail(fmt.Sprintf("%s:%s", host, portNumber), s.auth, username, m.To, m.ToBytes())
// // }
