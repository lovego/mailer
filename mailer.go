package mailer

import (
	"crypto/tls"
	"net"
	"net/mail"
	"net/smtp"
	"net/url"
)

type Mailer struct {
	client *smtp.Client
	sender string
}

func New(mailerUrl string) (*Mailer, error) {
	uri, err := url.Parse(mailerUrl)
	if err != nil {
		return nil, err
	}
	if opener, err := getOpener(uri); err != nil {
		return nil, err
	} else {
		return opener()
	}
}

func getOpener(uri *url.URL) (func() (*Mailer, error), error) {
	q := uri.Query()
	user, err := mail.ParseAddress(q.Get(`user`))
	if err != nil {
		return nil, err
	}
	auth := smtp.PlainAuth(``, user.Address, q.Get(`pass`), uri.Hostname())

	return func() (*Mailer, error) {
		if client, err := Open(uri.Host, auth); err != nil {
			return nil, err
		} else {
			return &Mailer{client: client, sender: user.String()}, nil
		}
	}, nil
}

func Open(addr string, auth smtp.Auth) (*smtp.Client, error) {
	client, err := smtp.Dial(addr)
	if err != nil {
		return nil, err
	}
	if ok, _ := client.Extension("STARTTLS"); ok {
		host, _, err := net.SplitHostPort(addr)
		if err != nil {
			return nil, err
		}
		if err = client.StartTLS(&tls.Config{ServerName: host}); err != nil {
			return nil, err
		}
	}
	if auth != nil {
		if ok, _ := client.Extension("AUTH"); ok {
			if err = client.Auth(auth); err != nil {
				return nil, err
			}
		}
	}
	return client, nil
}
