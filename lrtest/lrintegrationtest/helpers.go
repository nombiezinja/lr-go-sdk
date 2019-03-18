package lrintegrationtest

import (
	"encoding/base64"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	lr "bitbucket.org/nombiezinja/lr-go-sdk"
	lraccount "bitbucket.org/nombiezinja/lr-go-sdk/api/account"
	lrauthentication "bitbucket.org/nombiezinja/lr-go-sdk/api/authentication"
	"bitbucket.org/nombiezinja/lr-go-sdk/api/customobject"
	"bitbucket.org/nombiezinja/lr-go-sdk/lrjson"
)

func setupAccount(t *testing.T) (string, string, string, string, *lr.Loginradius, func(t *testing.T)) {
	t.Log("Setting up test case")

	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	loginradius, _ := lr.NewLoginradius(&cfg)
	authlr := lraccount.Loginradius{loginradius}

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	testEmail := "lrtest" + timeStamp + "@mailinator.com"
	testEmails := TestEmailArr{{"Primary", testEmail}, {"Secondary", "1" + testEmail}}
	username := "lrtest" + timeStamp
	phoneID := "+1" + timeStamp
	testAccount := AccountSetup{true, true, testEmails, testEmail, username, phoneID}

	response, err := lraccount.Loginradius(authlr).PostManageAccountCreate(testAccount)
	if err != nil {
		t.Errorf("Error calling PostManageAccountCreate from setupAccount: %v", err)
	}
	user, err := lrjson.DynamicUnmarshal(response.Body)
	uid := user["Uid"].(string)
	if err != nil || uid == "" {
		t.Errorf("Error creating account: %+v", err)
	}

	return phoneID, username, uid, testEmail, loginradius, func(t *testing.T) {
		t.Log("Tearing down test case")
		_, err := lraccount.Loginradius(lraccount.Loginradius{loginradius}).DeleteManageAccount(uid)
		if err != nil {
			t.Errorf("Error cleaning up account: %+v", err)
		}
	}
}

func setupEmailVerificationAccount(t *testing.T) (string, string, string, *lr.Loginradius, func(t *testing.T)) {
	t.Log("Setting up test case")

	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	loginradius, _ := lr.NewLoginradius(&cfg)
	authlr := lrauthentication.Loginradius{Client: loginradius}

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	testEmail := "lrtest" + timeStamp + "@mailinator.com"
	testEmails := TestEmailArr{{"Primary", testEmail}}
	username := "lrtest" + timeStamp

	phoneID := "+" + timeStamp
	testAccount := AccountSetup{false, false, testEmails, testEmail, username, phoneID}
	response, err := lraccount.Loginradius(authlr).PostManageAccountCreate(testAccount)
	user, _ := lrjson.DynamicUnmarshal(response.Body)
	uid := user["Uid"].(string)
	if err != nil || uid == "" {
		t.Errorf("Error creating account: %+v", err)
	}

	tokenGen := TestEmail{testEmail}
	response, err = lraccount.Loginradius(authlr).PostManageEmailVerificationToken(tokenGen)
	data, _ := lrjson.DynamicUnmarshal(response.Body)
	token := data["VerificationToken"].(string)
	if err != nil {
		t.Errorf("Error generating token: %+v", err)
	}

	return phoneID, testEmail, token, loginradius, func(t *testing.T) {
		t.Log("Tearing down test case")
		_, err := lraccount.Loginradius(lraccount.Loginradius{loginradius}).DeleteManageAccount(uid)
		if err != nil {
			t.Errorf("Error cleaning up account: %+v", err)
		}
	}
}

func setupLogin(t *testing.T) (string, string, string, string, string, *lr.Loginradius, func(t *testing.T)) {
	phoneID, username, testuid, testEmail, loginradius, teardownTestCase := setupAccount(t)
	authlr := lrauthentication.Loginradius{loginradius}
	testLogin := TestEmailLogin{testEmail, testEmail}
	response, err := lrauthentication.Loginradius(authlr).PostAuthLoginByEmail(testLogin)
	session, _ := lrjson.DynamicUnmarshal(response.Body)
	accessToken := session["access_token"].(string)
	if err != nil || accessToken == "" {
		t.Errorf("Error logging in: %+v", err)
	}
	loginradius.Context.Token = accessToken
	return phoneID, username, testuid, testEmail, accessToken, loginradius, func(t *testing.T) {
		defer teardownTestCase(t)
	}
}

func setupCustomObject(t *testing.T) (string, string, string, string, *lr.Loginradius, func(t *testing.T)) {
	t.Log("Setting up test case")
	_, _, uid, _, accessToken, lrclient, teardownTestCase := setupLogin(t)

	objName := os.Getenv("CUSTOMOBJECTNAME")
	time := time.Now()
	timestamp := time.Format("20060102150405")
	customObj := map[string]string{
		"custom1": timestamp + "0",
		"custom2": timestamp + "1",
	}
	resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).PostCustomObjectCreateByUID(uid, map[string]string{"objectname": objName}, customObj)
	if err != nil {
		t.Errorf("Error creating custom object: %v", err)
	}
	obj, err := lrjson.DynamicUnmarshal(resp.Body)
	objId := obj["Id"].(string)
	return accessToken, uid, objId, objName, lrclient, func(t *testing.T) {
		t.Log("Tearing down test case")
		defer teardownTestCase(t)
		_, err = customobject.Loginradius(customobject.Loginradius{lrclient}).DeleteCustomObjectByObjectRecordIDAndUID(uid, objId, map[string]string{"objectname": objName})
		if err != nil {
			t.Errorf("Error deleting custom object")
		}
	}
}

func genGUID() string {
	rand.Seed(time.Now().UnixNano())
	buff := make([]byte, 64)
	rand.Read(buff)
	return base64.StdEncoding.EncodeToString(buff)
}
