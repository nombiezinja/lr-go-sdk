package unittest

import (
	"os"
	"testing"

	lr "bitbucket.org/nombiezinja/lr-go-sdk"

	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
)

func TestPostAuthUserRegistrationByEmail(t *testing.T) {
	path := "/identity/v2/auth/register?apiKey=&emailtemplate=&options=&verificationurl="
	response := httprutils.Response{
		StatusCode: 200,
		Body:       `{"IsPosted":true,"Data":null}"`,
		Headers:    map[string][]string{},
	}

	stub := initTestServer(path, response)

	defer stub.Close()

	os.Setenv("DOMAIN", stub.URL)
	user := "user struct"
	res, _ := lr.PostAuthUserRegistrationByEmail("", "", "", user)
	if res.StatusCode != response.StatusCode || res.Body != response.Body {
		t.Errorf("PostAuthUserRegistrationByEmailWithStub: Expected %v, received %v", response, res)
	}
}

func TestGetAuthVerifyEmail(t *testing.T) {

	path := "/identity/v2/auth/email"
	response := httprutils.Response{
		StatusCode: 200,
		Body:       `{"IsPosted":true,"Data":null}"`,
		Headers:    map[string][]string{},
	}

	stub := initTestServer(path, response)

	defer stub.Close()

	os.Setenv("DOMAIN", stub.URL)
	res, _ := lr.GetAuthVerifyEmail("", "", "")
	if res.StatusCode != response.StatusCode || res.Body != response.Body {
		t.Errorf("PostAuthUserRegistrationByEmailWithStub: Expected %v, received %v", response, res)
	}
}
