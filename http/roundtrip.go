package roundtrip

import (
	"log"
	"net/http"
)

func (t *LogTransport) transport() http.RoundTripper {
	if t.Transport == nil {
		return http.DefaultTransport
	}
	return t.Transport
}

type LogTransport struct {
	Transport http.RoundTripper
	count     int
}

func (t *LogTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	tr := t.transport()
	log.Printf("logging transport: %+v", tr)
	return tr.RoundTrip(req)
}

/*
	## RoundTripperの実装

	* interface
	RoundTrip(*Request) (*Response, error)のシグネチャを満たせば実装できる。
	https://github.com/golang/go/blob/release-branch.go1.14/src/net/http/client.go#L115

	* RoundTrip()の実装
	RoundTrip()内で、roundtrip()をcallする。
	https://github.com/golang/go/blob/release-branch.go1.14/src/net/http/roundtrip.go#L9
	https://github.com/golang/go/blob/release-branch.go1.14/src/net/http/transport.go#L483

	RoundTrip()のレシーバーのフィールド内の http.RoundTripper がnilかどうかで差し替えられる
	https://github.com/golang/go/blob/release-branch.go1.14/src/net/http/client.go#L169

	* transportの実装はsend()メソッド内で入れ替えられる
	https://github.com/golang/go/blob/release-branch.go1.14/src/net/http/transport.go#L37
*/
