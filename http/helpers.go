package http

import (
	"bytes"
	"html/template"
	"io"
	"panth3rWaitlistBackend/internal/env"
	"panth3rWaitlistBackend/templ"

	gomail "gopkg.in/mail.v2"
)

type templData struct {
	Name string
}

func sendConfirmationMail(name, email string) error {

	var (
		from = env.GetEnvVar().Mail.EmailAcc
		tpl  bytes.Buffer
		s    = templData{Name: name}
		data []byte
		err  error
	)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Panth3r Confirmation mail")

	if err = prepareMail(s, &tpl); err != nil {
		return err
	}

	if data, err = templ.Panth3rFrame.ReadFile("Panth3rFrame.png"); err != nil {
		return err
	}
	m.Attach("Panth3rFrame.png", gomail.SetHeader(map[string][]string{
		"Content-ID": {"Panth3rFrame"},
	}), gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(data)
		return err
	}))

	m.SetBody("text/html", tpl.String())

	d := gomail.NewDialer(env.GetEnvVar().Mail.EmailSmtpServerHost, env.GetEnvVar().Mail.EmailSmtpServerPort, from, env.GetEnvVar().Mail.EmailPswd)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func prepareMail(data templData, mailBuf *bytes.Buffer) error {

	t, err := template.ParseFS(templ.EmailHTML, "mail.html")
	if err != nil {
		return err
	}

	if err := t.Execute(mailBuf, data); err != nil {
		return err
	}

	return nil
}
