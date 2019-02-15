package integrationtest

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	loginradius "bitbucket.org/nombiezinja/lr-go-sdk"
	lrjson "bitbucket.org/nombiezinja/lr-go-sdk/json"
)

func setupAccount(t *testing.T) (string, string, string, string, func(t *testing.T)) {
	t.Log("Setting up test case")
	SetTestCredentials()

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	testEmail := "lrtest" + timeStamp + "@mailinator.com"
	testEmails := TestEmailArr{{"Primary", testEmail}, {"Secondary", "1" + testEmail}}
	username := "lrtest" + timeStamp
	phoneID := "+1" + timeStamp
	testAccount := AccountSetup{true, true, testEmails, testEmail, username, phoneID}

	response, err := loginradius.PostManageAccountCreate(testAccount)
	if err != nil {
		t.Errorf("Error calling PostManageAccountCreate from setupAccount: %v", err)
	}
	user, err := lrjson.DynamicUnmarshal(response.Body)
	uid := user["Uid"].(string)
	if err != nil || uid == "" {
		t.Errorf("Error creating account")
		fmt.Println(err)
	}

	return phoneID, username, uid, testEmail, func(t *testing.T) {
		t.Log("Tearing down test case")
		_, err = loginradius.DeleteManageAccount(uid)
		if err != nil {
			t.Errorf("Error cleaning up account")
			fmt.Println(err)
		}
	}
}

func setupEmailVerificationAccount(t *testing.T) (string, string, string, func(t *testing.T)) {
	t.Log("Setting up test case")
	SetTestCredentials()
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	testEmail := "lrtest" + timeStamp + "@mailinator.com"
	testEmails := TestEmailArr{{"Primary", testEmail}}
	username := "lrtest" + timeStamp

	phoneID := "+" + timeStamp
	testAccount := AccountSetup{false, false, testEmails, testEmail, username, phoneID}
	response, err := loginradius.PostManageAccountCreate(testAccount)
	user, _ := lrjson.DynamicUnmarshal(response.Body)
	uid := user["Uid"].(string)
	if err != nil || uid == "" {
		t.Errorf("Error creating account")
		fmt.Println(err)
	}

	tokenGen := TestEmail{testEmail}
	response, err = loginradius.PostManageEmailVerificationToken(tokenGen)
	data, _ := lrjson.DynamicUnmarshal(response.Body)
	token := data["VerificationToken"].(string)
	if err != nil {
		t.Errorf("Error generating token")
		fmt.Println(err)
	}

	return phoneID, testEmail, token, func(t *testing.T) {
		t.Log("Tearing down test case")
		_, err2 := loginradius.DeleteManageAccount(uid)
		if err2 != nil {
			t.Errorf("Error cleaning up account")
			fmt.Println(err2)
		}
	}
}

func setupLogin(t *testing.T) (string, string, string, string, string, func(t *testing.T)) {
	// SetTestCredentials()
	phoneID, username, testuid, testEmail, teardownTestCase := setupAccount(t)
	testLogin := TestEmailLogin{testEmail, testEmail}
	response, err := loginradius.PostAuthLoginByEmail("", "", "", "", "", testLogin)
	session, _ := lrjson.DynamicUnmarshal(response.Body)
	accessToken := session["access_token"].(string)
	if err != nil || accessToken == "" {
		t.Errorf("Error logging in")
		fmt.Println(err)
	}
	return phoneID, username, testuid, testEmail, accessToken, func(t *testing.T) {
		defer teardownTestCase(t)
	}
}
