package lrunittest

import (
	"net/http"
	"net/http/httptest"

	lr "bitbucket.org/nombiezinja/lr-go-sdk"
	lrauth "bitbucket.org/nombiezinja/lr-go-sdk/api/authentication"
	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
)

func initLr() lrauth.Loginradius {
	cfg := lr.Config{
		ApiKey:    "abcd1234",
		ApiSecret: "abcd1234",
	}

	initLr, _ := lr.NewLoginradius(&cfg)
	loginradius := lrauth.Loginradius{initLr}
	return loginradius
}

func initTestServer(path string, resp httprutils.Response) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(resp.StatusCode)
		w.Write([]byte(resp.Body))
	}))
}
