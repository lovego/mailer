package mailer

import (
	"testing"
	"time"

	"github.com/lovego/email"
)

func TestSend(t *testing.T) {
	mailer, err := New(
		`mailer://smtp.qq.com:25/?user=小美<xiaomei-go@qq.com>&pass=zjsbosjlhgugechh`,
	)
	if err != nil {
		panic(err)
	}

	e := &email.Email{
		To:      []string{"张绍兴<shaoxing.zhang@hztl3.com>", "侯志良<bughou@gmail.com>"},
		Subject: "一个非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常非常长的标题",
		Text: []byte(`
        sql: Scan error on column index 17: can't convert =04Q3PA=00=00 to decimal
		<b>超文本!</b>
		a very very very very very very very very very very very very very very very very very very very very long line.
`),
	}
	if err := mailer.Send(e, 10*time.Second); err != nil {
		panic(err)
	}
}
