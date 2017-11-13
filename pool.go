package mailer

import (
	"net/url"
	"strconv"
	"sync"
)

type Pool struct {
	maxOpen int // <= 0 means unlimited
	maxIdle int // <  0 means unlimited
	opener  func() (*Mailer, error)
	mailers []*Mailer
	mu      sync.Mutex
}

func NewPool(mailerUrl string) (*Pool, error) {
	uri, err := url.Parse(mailerUrl)
	if err != nil {
		return nil, err
	}
	maxOpen, maxIdle, err := poolParams(uri.Query())
	if err != nil {
		return nil, err
	}
	opener, err := getOpener(uri)
	if err != nil {
		return nil, err
	}

	return &Pool{
		maxOpen: maxOpen, maxIdle: maxIdle, opener: opener,
	}, nil
}

func (p *Pool) Send() error {
	return nil
}

func poolParams(q url.Values) (maxOpen, maxIdle int, err error) {
	maxOpen, maxIdle = 10, 3
	if str := q.Get(`maxOpen`); str != `` {
		if maxOpen, err = strconv.Atoi(str); err != nil {
			return
		}
	}
	if str := q.Get(`maxIdle`); str != `` {
		if maxIdle, err = strconv.Atoi(str); err != nil {
			return
		}
	}
	return
}
