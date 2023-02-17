package main

import (
	"net/smtp"
)

type AuthType string
type PlainAuth smtp.Auth

// Authentication is a struct that contains the authentication information for the SMTP server.
// It is used to create a new Sender. The fields are:
//
//	Username: The username for the SMTP server.
//	Password: The password for the SMTP server.
//	Host: The host for the SMTP server.
//	Port: The port for the SMTP server.
type Authentication struct {
	Username string
	Password string
	Host     string
	Port     string
}

// Sender is a struct that contains the authentication information for the SMTP server.
// type Sender struct {
// 	authType  AuthType
// 	plainAuth PlainAuth
// }

// func New(authType AuthType) *Sender {
// 	if authType == "plain" {
// 		auth := smtp.PlainAuth("", username, password, host)
// 		return &Sender{authType}
// 	}

// 	return &Sender{authType}
// }

// func New() *Sender {
// 	auth := smtp.PlainAuth("", username, password, host)
// 	return &Sender{auth}
// }

// func (s *Sender) Send(m *Message) error {
// 	return smtp.SendMail(fmt.Sprintf("%s:%s", host, portNumber), s.auth, username, m.To, m.ToBytes())
// }
