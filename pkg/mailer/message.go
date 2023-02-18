package mailer

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

// Message is a struct that contains the information for an email message.
type message struct {
	to          []string
	cc          []string
	bcc         []string
	subject     string
	body        string
	attachments map[string][]byte
}

// IMessage is an interface that defines the methods for an email message.
// It is used to create a new message, set the To, Cc, Bcc, subject, body, and attachments.
// The ToBytes method converts the message to a byte array.
// The SetTo method sets the To field of the message.
// The GetTo method returns the To field of the message.
// The SetCc method sets the Cc field of the message.
// The GetCc method returns the Cc field of the message.
// The SetBcc method sets the Bcc field of the message.
// The GetBcc method returns the Bcc field of the message.
// The SetSubject method sets the subject of the message.
// The GetSubject method returns the subject of the message.
// The SetBody method sets the body of the message.
// The GetBody method returns the body of the message.
// The SetAttachFiles method attaches a file to the message.
// The GetAttachFiles method returns the attached files.
type IMessage interface {
	ToBytes() []byte
	SetTo([]string)
	GetTo() []string
	SetCc([]string)
	GetCc() []string
	SetBcc([]string)
	GetBcc() []string
	SetSubject(string)
	GetSubject() string
	SetBody(string)
	GetBody() string
	SetAttachFiles(string) error
	GetAttachFiles() map[string][]byte
}

var _ IMessage = (*message)(nil)

// NewMessage is a function that creates a new message. The subject and body are required.
// The subject and body are strings. The subject is the subject of the email message.
// The body is the body of the email message. The body can be in HTML format.
func NewMessage(subject, body string) *message {
	return &message{subject: subject, body: body, attachments: make(map[string][]byte)}
}

// SetTo is a method that sets the To field of the message. The To field is a list of email addresses.
func (m *message) SetTo(to []string) {
	m.to = to
}

// GetTo is a method that returns the To field of the message.
func (m *message) GetTo() []string {
	return m.to
}

// SetCc is a method that sets the Cc field of the message. The Cc field is a list of email addresses.
func (m *message) SetCc(cc []string) {
	m.cc = cc
}

// GetCc is a method that returns the Cc field of the message.
func (m *message) GetCc() []string {
	return m.cc
}

// SetBcc is a method that sets the Bcc field of the message. The Bcc field is a list of email addresses.
func (m *message) SetBcc(bcc []string) {
	m.bcc = bcc
}

// GetBcc is a method that returns the Bcc field of the message.
func (m *message) GetBcc() []string {
	return m.bcc
}

// SetSubject is a method that sets the subject of the message. The subject is a string.
// The SetSubject method is optional because the subject is a required field when creating a new message.
// SetSubject is used to change the subject of the message.
func (m *message) SetSubject(subject string) {
	m.subject = subject
}

// GetSubject is a method that returns the subject of the message.
func (m *message) GetSubject() string {
	return m.subject
}

// SetBody is a method that sets the body of the message. The body is a string.
// The SetBody method is optional because the body is a required field when creating a new message.
// SetBody is used to change the body of the message.
func (m *message) SetBody(body string) {
	m.body = body
}

// GetBody is a method that returns the body of the message.
func (m *message) GetBody() string {
	return m.body
}

// AttachFile is a method that attaches a file to the message. The file is read from the given path.
// The file name is used as the attachment name.
// The file content is base64 encoded.
func (m *message) SetAttachFiles(src string) error {
	content, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	_, fileName := filepath.Split(src)
	m.attachments[fileName] = content
	return nil
}

func (m *message) GetAttachFiles() map[string][]byte {
	return m.attachments
}

// ToBytes is a method that converts the message to a byte array.
// The byte array is used to send the message to the SMTP server. The byte array is in the format
// of an email message.
func (m *message) ToBytes() []byte {
	// If there are attachments it will return true otherwise false.
	hasAttachment := len(m.attachments) > 0

	// The message is built as a string and then converted to a byte array.
	var message string = ""
	message += "To: " + strings.Join(m.GetTo(), ";") + "\n"
	message += "Cc: " + strings.Join(m.GetCc(), ";") + "\n"
	message += "Bcc: " + strings.Join(m.GetBcc(), ";") + "\n"
	message += "Subject: " + m.GetSubject() + "\n"
	message += "Body: " + m.GetBody() + "\n"
	message += "MIME-Version: 1.0\n"

	// If there are attachments, the message is in the multipart/mixed format.
	// Otherwise, the message is in the text/plain format.
	if hasAttachment {
		message += "Content-Type: multipart/mixed; boundary=\"BOUNDARY\"\n\n"
		message += "--BOUNDARY\n"
		message += "Content-Transfer-Encoding: base64\n"
		for k, v := range m.GetAttachFiles() {
			message += "--BOUNDARY\n"
			message += "Content-Type: " + http.DetectContentType(v) + "; name=\"" + k + "\"\n"
			message += "Content-Transfer-Encoding: base64\n"
			message += "Content-Disposition: attachment; filename=\"" + k + "\"\n"
			message += base64.StdEncoding.EncodeToString(v) + "\n"
		}

		message += "--BOUNDARY--\n"
	} else {
		message += "Content-Type: text/plain; charset=utf-8\n"
		message += m.GetBody() + "\n"
	}

	return []byte(message)
}
