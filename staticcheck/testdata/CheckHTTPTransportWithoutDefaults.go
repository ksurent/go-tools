package pkg

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

func fn1() {
	_ = &http.Transport{} // MATCH /built new Transport without any of the default settings/
	_ = http.Transport{}  // MATCH /built new Transport without any of the default settings/
	_ = http.Transport{   // MATCH /built new Transport without any of the default settings/
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	_ = http.Client{
		Transport: &http.Transport{}, // MATCH /built new Transport without any of the default settings/
	}

	var (
		tr1 http.Transport // MATCH /built new Transport without any of the default settings/
		tr2 http.Transport // MATCH /built new Transport without any of the default settings/
	)
	_ = tr1
	_ = tr2

	tr3 := http.Transport{} // MATCH /built new Transport without any of the default settings/
	tr3.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	_ = tr3
}

func fn2() {
	var tr1 *http.Transport
	tr1 = &http.Transport{IdleConnTimeout: 42 * time.Second}
	_ = tr1

	tr2 := http.Transport{}
	tr2.TLSHandshakeTimeout = 42 * time.Second

	_ = http.Transport{ExpectContinueTimeout: 42 * time.Second}
	_ = http.Transport{MaxIdleConns: 42}
	_ = http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return nil, nil
		},
	}
	_ = http.Transport{DialContext: (&net.Dialer{}).DialContext}
}

// Known issue: no pointer analysis.
func fn3() {
	setDefaultField := func(tr *http.Transport) {
		tr.IdleConnTimeout = 42 * time.Second
	}
	tr := http.Transport{} // MATCH /built new Transport without any of the default settings/
	setDefaultField(&tr)
}

// Known caveat.
func fn4() {
	tr := http.Transport{}
	if time.Now().Unix()%2 == 0 {
		tr.IdleConnTimeout = 42 * time.Second
	}
}
