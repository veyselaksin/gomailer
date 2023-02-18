package mailer

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

// Message is a struct that contains the information for an email message.
type Message struct {
	From        string
	To          []string
	CC          []string
	BCC         []string
	Subject     string
	Body        string
	Attachments map[string][]byte
}

type IMessage interface {
	// ToBytes() []byte
	AttachFile(string) error
}

var _ IMessage = (*Message)(nil)

// SimpleMessage is a method that creates a new message with the given subject and body.
// The attachments are optional.
func NewMessage(subject, body string) *Message {
	return &Message{Subject: subject, Body: body, Attachments: make(map[string][]byte)}
}

// AttachFile is a method that attaches a file to the message. The file is read from the given path.
// The file name is used as the attachment name.
// The file content is base64 encoded.
func (m *Message) AttachFile(src string) error {
	content, err := ioutil.ReadFile(src)
	if err != nil {
		panic(err)
	}

	_, fileName := filepath.Split(src)
	m.Attachments[fileName] = content
	return nil
}

// ToBytes is a method that converts the message to a byte array.
// The byte array is used to send the message to the SMTP server. The byte array is in the format
// of an email message.
func (m *Message) ToBytes() []byte {
	// If there are attachments it will return true otherwise false.
	hasAttachment := len(m.Attachments) > 0

	// The message is built as a string and then converted to a byte array.
	var message string = ""
	message += "From: " + m.From + "\n"
	message += "To: " + strings.Join(m.To, ";") + "\n"
	message += "Cc: " + strings.Join(m.CC, ";") + "\n"
	message += "Bcc: " + strings.Join(m.BCC, ";") + "\n"
	message += "Subject: " + m.Subject + "\n"
	message += "Body: " + m.Body + "\n"
	message += "MIME-Version: 1.0\n"

	// If there are attachments, the message is in the multipart/mixed format.
	// Otherwise, the message is in the text/plain format.
	if hasAttachment {
		message += "Content-Type: multipart/mixed; boundary=frontier\n"
		message += "This is a multi-part message in MIME format.\n"
		message += "--frontier\n"
		message += "Content-Type: text/plain; charset=us-ascii\n"
		message += "Content-Transfer-Encoding: 7bit\n"
		message += m.Body + "\n"
		message += "--frontier\n"
		for k, v := range m.Attachments {
			message += "Content-Type: " + http.DetectContentType(v) + "; name=\"" + k + "\"\n"
			message += "Content-Transfer-Encoding: base64\n"
			message += "Content-Disposition: attachment; filename=\"" + k + "\"\n"
			message += base64.StdEncoding.EncodeToString(v) + "\n"
			message += "--frontier\n"
		}
	} else {
		message += "Content-Type: text/plain; charset=utf-8\n"
		message += "Content-Transfer-Encoding: 7bit\n"
		message += m.Body + "\n"
	}

	return []byte(message)
}
