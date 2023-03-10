package mailer

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

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

// SMTPDialer is a struct that contains the dialer information for the SMTP server.
// SMTPDialer is used for servers that do not require authentication.
// It is used to create a new Sender. The fields are:
//
//	Host: The host for the SMTP server.
//	Port: The port for the SMTP server.
//
// If the SMTP server does not require authentication, then the Authentication struct is not needed.
type SMTPDialer struct {
	Host string
	Port string
}

// ISender is an interface that defines the methods for a sender.
// It is used to create a new sender, and send an email message.
// The SendMail method sends an email message.
type ISender interface {
	SendMail(m *message) error
	SendMailTLS(m *message, tlsconfig *tls.Config) error
}

// AuthType is a type that defines the authentication type for the sender.
type AuthType string

// Sender is a struct that contains the authentication information for the SMTP server.
type sender struct {
	plainAuth *smtp.Auth
	client    *smtp.Client
	authType  AuthType
	auth      *Authentication
}

var _ ISender = (*sender)(nil)

// NewPlainAuth is a function that creates a new sender. The Authentication struct is required.
// The Authentication struct contains the authentication information for the SMTP server.
// The Authentication struct is used to create a new Sender. The fields are:
//
//	Username: The username for the SMTP server.
//	Password: The password for the SMTP server.
//	Host: The host for the SMTP server.
//	Port: The port for the SMTP server.
func NewPlainAuth(auth *Authentication) *sender {
	plainAuth := smtp.PlainAuth("", auth.Username, auth.Password, auth.Host)
	return &sender{plainAuth: &plainAuth, authType: "plain", auth: auth}
}

// NewDialer is a function that creates a new sender. The SMTPDialer struct is required.
// The SMTPDialer struct contains the dialer information for the SMTP server.
// SMTPDialer is used for servers that do not require authentication.
// It is used to create a new Sender. The fields are:
//
//	Host: The host for the SMTP server.
//	Port: The port for the SMTP server.
func NewDialer(smtpDialer *SMTPDialer) *sender {
	client, _ := smtp.Dial(smtpDialer.Host + ":" + smtpDialer.Port)
	return &sender{client: client, authType: "dialer"}
}

// SendMail is a method that sends an email message.
// The message struct is required.
// The message struct contains the email message information.
// The message struct is used to create a new email message.
func (s *sender) SendMail(m *message) error {
	return smtp.SendMail(fmt.Sprintf("%s:%s", s.auth.Host, s.auth.Port), *s.plainAuth, s.auth.Username, m.GetTo(), m.ToBytes())
}

// SendMailTLS is a method that sends an email message. The message struct is required and the tls.Config struct is required.
// The message struct contains the email message information.
// The tls.Config struct contains the TLS configuration information.
func (s *sender) SendMailTLS(m *message, tlsConfig *tls.Config) error {
	s.client.StartTLS(tlsConfig)

	// Set the sender and recipient first before calling Data() to send the email.
	if err := s.client.Mail(m.GetFrom()); err != nil {
		return err
	}

	recipients := append(append(m.GetTo(), m.GetCc()...), m.GetBcc()...)
	for _, k := range recipients {
		s.client.Rcpt(k)
	}

	// Send the email body. The message is sent in the same format as the RFC 822.
	w, err := s.client.Data()
	if err != nil {
		return err
	}

	w.Write(m.ToBytes())
	w.Close()

	return nil
}
