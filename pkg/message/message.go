package message

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

// Message is a struct that contains the information for an email message.
type Message struct {
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
	return &Message{Subject: subject, Body: body}
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

func (m *Message) ToBytes() []byte {
	buf := bytes.NewBuffer(nil)
	template := fmt.Sprintf("Subject: %s\n To: %s\n Cc: %s\n Bcc: %s\n MIME-Version: 1.0\n %s",
		m.Subject, strings.Join(m.To, ","), strings.Join(m.CC, ","), strings.Join(m.BCC, ","), m.Body)
	buf.WriteString(template)

	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()

	hasAttachment := len(m.Attachments) > 0
	if hasAttachment {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
		buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	} else {
		buf.WriteString("Content-Type: text/plain; charset=utf-8\n")
	}

	if hasAttachment {
		for k, v := range m.Attachments {
			buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(v)))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))

			b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
			base64.StdEncoding.Encode(b, v)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", boundary))
		}

		buf.WriteString("--")
	}

	fmt.Println(string(buf.Bytes()))

	return buf.Bytes()
}
