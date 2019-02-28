package lrunittest

import (
	"net/http"
	"net/http/httptest"

	lr "bitbucket.org/nombiezinja/lr-go-sdk"
	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
)

func initLr() lr.Loginradius {
	cfg := lr.Config{
		ApiKey:    "abcd1234",
		ApiSecret: "abcd1234",
	}

	lrclient, _ := lr.NewLoginradius(&cfg)
	return *lrclient
}

func initTestServer(path string, resp httprutils.Response) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(resp.StatusCode)
		w.Write([]byte(resp.Body))
	}))
}
