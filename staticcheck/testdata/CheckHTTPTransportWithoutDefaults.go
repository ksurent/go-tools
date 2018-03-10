package pkg

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

func fn1() {
	tr := &http.Transport{} // MATCH /built new Transport without any of the default settings/
	_ = tr
}

func fn2() {
	tr := http.Transport{} // MATCH /built new Transport without any of the default settings/
	_ = tr
}

func fn3() {
	tr := http.Transport{ // MATCH /built new Transport without any of the default settings/
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	_ = tr
}

func fn4() {
	client := http.Client{
		Transport: &http.Transport{}, // MATCH /built new Transport without any of the default settings/
	}
	_ = client
}

func fn5() {
	var (
		tr1 http.Transport  // MATCH /built new Transport without any of the default settings/
		tr2 *http.Transport // MATCH /built new Transport without any of the default settings/
	)
	_ = tr1
	_ = tr2
}

func fn6() {
	tr := http.Transport{}
	if time.Now().Unix()%2 == 0 {
		tr.TLSHandshakeTimeout = 2 * time.Second
	}

	tr = http.Transport{} // MATCH /built new Transport without any of the default settings/
}

func fn7() {
	tr := http.Transport{} // MATCH /built new Transport without any of the default settings/
	tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func fn8() {
	var tr *http.Transport
	tr = &http.Transport{IdleConnTimeout: 500 * time.Millisecond}
	_ = tr

	_ = http.Transport{ExpectContinueTimeout: 3 * time.Second}
	_ = http.Transport{MaxIdleConns: 42}
	_ = http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return nil, nil
		},
	}
	_ = http.Transport{DialContext: (&net.Dialer{}).DialContext}
}
