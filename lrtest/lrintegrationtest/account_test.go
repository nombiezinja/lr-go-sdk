package lrintegrationtest

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	lrjson "bitbucket.org/nombiezinja/lr-go-sdk/lrjson"

	lr "bitbucket.org/nombiezinja/lr-go-sdk"
	lraccount "bitbucket.org/nombiezinja/lr-go-sdk/api/account"
	lrauthentication "bitbucket.org/nombiezinja/lr-go-sdk/api/authentication"
)

func TestGetManageAccountProfilesByEmail(t *testing.T) {
	fmt.Println("Starting test TestGetManageAccountProfilesByEmail")
	_, _, testuid, testEmail, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	response, err := lraccount.GetManageAccountProfilesByEmail(testEmail)
	session, _ := lrjson.DynamicUnmarshal(response.Body)
	uid := session["Uid"].(string)
	if err != nil || uid != testuid {
		t.Errorf("Error retrieving profile associated with email")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestDeleteManageAccount(t *testing.T) {
	fmt.Println("Starting test TestDeleteManageAccount")
	_, _, testuid, _, _, _ := setupAccount(t)
	_, err := lraccount.DeleteManageAccount(testuid)
	if err != nil {
		t.Errorf("Error deleting account")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostManageAccountCreate(t *testing.T) {
	fmt.Println("Starting test TestPostManageAccountCreate")
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrtmp, _ := lr.NewLoginradius(&cfg)
	loginradius := lrauthentication.Loginradius{lrtmp}

	testEmail := "lrtest" + strconv.FormatInt(time.Now().Unix(), 10) + "@mailinator.com"
	testEmails := TestEmailArr{{"Primary", testEmail}}
	testAccount := TestAccount{true, testEmails, testEmail}

	response, err := lraccount.Loginradius(loginradius).PostManageAccountCreate(testAccount)
	if err != nil {
		t.Errorf("Error calling PostManageAccountCreate: %v", err)
	}
	user, err := lrjson.DynamicUnmarshal(response.Body)
	uid := user["Uid"].(string)
	if err != nil || uid == "" {
		t.Errorf("Error returned from PostManageAccountCreate: %v", err)
	}
	_, err = lraccount.DeleteManageAccount(uid)
	if err != nil {
		t.Errorf("Error cleaning up account: %v", err)
	}
	fmt.Println("Test complete")
}

func TestPostManageEmailVerificationToken(t *testing.T) {

	fmt.Println("Starting test TestPostManageEmailVerificationToken")
	_, testEmail, _, loginradius, teardownTestCase := setupEmailVerificationAccount(t)
	defer teardownTestCase(t)
	emailObj := TestEmail{testEmail}
	response, err := lraccount.Loginradius(lraccount.Loginradius{loginradius}).PostManageEmailVerificationToken(emailObj)
	if err != nil {
		t.Errorf(" Error making call to PostManageEmailVerificationToken: %v", err)
	}
	session, _ := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || session["VerificationToken"].(string) == "" {
		t.Errorf("Error returned from PostManageEmailVerificationToken: %v", err)
	}
	fmt.Println("Test complete")
}

func TestPutManageAccountUpdateSecurityQuestionConfig(t *testing.T) {
	fmt.Println("Starting test TestPutManageAccountUpdateSecurityQuestionConfig")
	_, _, testuid, _, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	fmt.Println(securityTest)
	response, err := lraccount.PutManageAccountUpdateSecurityQuestionConfig(testuid, securityTest)
	if err != nil {
		t.Errorf("Error making PutManageAccountUpdateSecurityQuestionConfig call")
	}
	profile, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || profile["Uid"].(string) != testuid {
		t.Errorf("Error returned from PutManageAccountUpdateSecurityQuestionConfig: %v", err)
	}
	fmt.Println("Test complete")
}
func TestPostManageForgotPasswordToken(t *testing.T) {
	fmt.Println("Starting test TestPostManageForgotPasswordToken")
	_, _, _, testEmail, loginradius, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	email := TestEmail{testEmail}
	response, err := lraccount.Loginradius(lraccount.Loginradius{loginradius}).PostManageForgotPasswordToken(email)
	if err != nil {
		t.Errorf("Error making call to PostManageForgotPasswordToken: %+v", err)
	}
	session, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || session["ForgotToken"].(string) == "" {
		t.Errorf("Error creating forgot password token")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}
