// Package sendmail provides the functionnality to send emails thru smtp server
package sendmail

import (
	"net/smtp"

	"github.com/softinnov/email"
)

// SMTPServer defines all information to send mail via smtp server
type SMTPServer struct {
	Expeditor string // from as "John Doe <john@doe.com>"
	URL       string // server url with port if neccessary server:999
	ID        string // login
	PW        string // password
	Host      string // server address
	Disclaim  string // disclaimer text at the end of all messages
}

// SendMail sends text or html with attachments
func (s SMTPServer) SendMail(dest []string, subject string, text []byte, html []byte, attachfile []string) error {
	e := email.NewEmail()
	e.From = s.Expeditor
	e.To = dest
	e.Subject = subject
	e.Text = append(text, []byte(s.Disclaim)...)
	e.HTML = append(html, []byte(s.Disclaim)...)
	for _, p := range attachfile {
		_, err := e.AttachFile(p)
		if err != nil {
			return err
		}
	}
	return e.Send(s.URL, smtp.PlainAuth("", s.ID, s.PW, s.Host))
}
