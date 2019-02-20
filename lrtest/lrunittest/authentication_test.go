package lrunittest

import (
	"testing"

	lrauth "bitbucket.org/nombiezinja/lr-go-sdk/api/authentication"
	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
)

const body = "test body"

func initTest(path string) lrauth.Loginradius {
	response := httprutils.Response{
		StatusCode: 200,
		Body:       body,
		Headers:    map[string][]string{},
	}

	stub := initTestServer(path, response)
	defer stub.Close()

	lr := initLr()
	lr.Domain = stub.URL
	return lr
}

func TestPostAuthUserRegistrationByEmail(t *testing.T) {
	lr := initTest("/identity/v2/auth/register?apiKey=&emailtemplate=&options=&verificationurl=")
	user := "user struct"
	res, _ := lrauth.Loginradius(lr).PostAuthUserRegistrationByEmail(map[string]string{}, user)
	if res.StatusCode != 200 || res.Body != body {
		t.Errorf("Unit TestPostAuthUserRegistrationByEmail: received %v", res)
	}
}

func TestGetAuthVerifyEmail(t *testing.T) {
	lr := initTest("/identity/v2/auth/email")
	res, _ := lrauth.Loginradius(lr).GetAuthVerifyEmail(map[string]string{})

	if res.StatusCode != 200 || res.Body != body {
		t.Errorf("Unit TestPostAuthUserRegistrationByEmail: received %v", res)
	}
}
