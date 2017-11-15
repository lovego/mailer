package mailer

import (
	"net/url"
	"strconv"
	"sync"

	"github.com/lovego/pool"
)

type Pool struct {
	pool *pool.Pool
}

func NewPool(mailerUrl string) (*Pool, error) {
	uri, err := url.Parse(mailerUrl)
	if err != nil {
		return nil, err
	}
	opener, err := getOpener(uri)
	if err != nil {
		return nil, err
	}
	p, err := pool.New(opener, uri.Query())
	if err != nil {
		return nil, err
	}

	return &Pool{pool: p}, nil
}

func (p *Pool) Send() error {
	return nil
}
