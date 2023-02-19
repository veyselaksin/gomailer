## **Introduction**

This is a simple mail sender library for Go. It is based on the [net/smtp](http://golang.org/pkg/net/smtp/) package.

This library is not yet a full featured mailer. It is just a simple wrapper for the net/smtp package.

It is very easy to use and can work integrated with all mail providers (Gmail, Hotmail, Yahoo, etc.) and can be used with any mail server.

You can send mail with attachments, html body, plain text body, etc. You can also send mail with multiple recipients.

## **Documentation**
[![Go Reference](https://pkg.go.dev/badge/github.com/veyselaksin/gomailer.svg)](https://pkg.go.dev/github.com/veyselaksin/gomailer)


# **Important Note**
Why you need to use this library? Because you can not send mail with attachments using the net/smtp package. You can only send plain text mail with this package. So, if you want to send mail with attachments, you need to use this library. Also it is not complicated to use like other mail libraries.

## **Features**
* Send mail with multiple attachments
* Send mail with html body or plain text body
* Send mail with multiple recipients
* Send mail with multiple cc recipients
* Send mail with multiple bcc recipients
* Support for TLS and SSL connections
* Also you can use your own connection (net.Conn) for sending mail
* Easy to use for EXCHANGE mail servers


## **Installation**

    go get github.com/veyselaksin/gomailer

## **Usage**
```golang
auth := mailer.Authentication{
    Username: "blabla@gmail.com",
    Password: "password",
    Host:     "smtp.gmail.com",
    Port:     "587",
}
sender := mailer.NewPlainAuth(&auth)

message := mailer.NewMessage("Hello World", "Hello World")
message.SetTo([]string{"blabla@hotmail.com", "blabla@gmail.com", "blabla@company.com"})
message.SetAttachFiles("./src/file")

sender.SendMail(message)
```
Well done! You have sent your first mail with gomailer. You can also send mail with html body, multiple recipients, multiple attachments, etc. You can also use your own connection for sending mail. You can also use this library for sending mail with EXCHANGE mail servers.

## **Examples**
```golang
auth := mailer.Authentication{
    Username: "blabla@gmail.com",
    Password: "password",
    Host:     "smtp.gmail.com",
    Port:     "587",
}
sender := mailer.NewPlainAuth(&auth)

bodyHTML := `
    <html>
        <body>
            <h1>Hello World</h1>
        </body>
    </html>
`
subject := "Hello World"

message := mailer.NewMessage(subject, bodyHTML)
message.SetTo([]string{"blabla@hotmail.com", "blabla@gmail.com", "blabla@company.com"})
message.SetAttachFiles("./src/file")

sender.SendMail(message)
```

If you want to send multiple recipients, you can use the SetTo method. You can also use the SetCc and SetBcc methods for sending mail with multiple cc and bcc recipients.

```golang
message.SetTo([]string{"blabla@gmail.com", "blabla@company.com"})
message.SetCc([]string{"blabla@gmail.com", "blabla@company.com"})
message.SetBcc([]string{"blabla@gmail.com", "blabla@company.com"})
```
If you want to send mail with multiple attachments, you can use the SetAttachFiles method. Unfortuntely, you need to duplicate the SetAttachFiles method for each attachment. I will fix this issue in the future.

```golang
message.SetAttachFiles("./src/file1")
message.SetAttachFiles("./src/file2")
message.SetAttachFiles("./src/file3")
```


