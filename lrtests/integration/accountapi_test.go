package integrationtest

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	lrjson "bitbucket.org/nombiezinja/lr-go-sdk/json"

	loginradius "bitbucket.org/nombiezinja/lr-go-sdk"
)

func TestGetManageAccountProfilesByEmail(t *testing.T) {
	fmt.Println("Starting test TestGetManageAccountProfilesByEmail")
	_, _, testuid, testEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	response, err := loginradius.GetManageAccountProfilesByEmail(testEmail)
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
	_, _, testuid, _, _ := setupAccount(t)
	_, err := loginradius.DeleteManageAccount(testuid)
	if err != nil {
		t.Errorf("Error deleting account")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostManageAccountCreate(t *testing.T) {
	fmt.Println("Starting test TestPostManageAccountCreate")
	SetTestCredentials()
	testEmail := "lrtest" + strconv.FormatInt(time.Now().Unix(), 10) + "@mailinator.com"
	testEmails := TestEmailArr{{"Primary", testEmail}}
	testAccount := TestAccount{true, testEmails, testEmail}

	response, err := loginradius.PostManageAccountCreate(testAccount)
	if err != nil {
		t.Errorf("Error calling PostManageAccountCreate: %v", err)
	}
	user, err := lrjson.DynamicUnmarshal(response.Body)
	uid := user["Uid"].(string)
	if err != nil || uid == "" {
		t.Errorf("Error returned from PostManageAccountCreate: %v", err)
	}
	_, err = loginradius.DeleteManageAccount(uid)
	if err != nil {
		t.Errorf("Error cleaning up account: %v", err)
	}
	fmt.Println("Test complete")
}

func TestPostManageEmailVerificationToken(t *testing.T) {
	fmt.Println("Starting test TestPostManageEmailVerificationToken")
	_, testEmail, _, teardownTestCase := setupEmailVerificationAccount(t)
	defer teardownTestCase(t)
	emailObj := TestEmail{testEmail}
	response, err := loginradius.PostManageEmailVerificationToken(emailObj)
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
	_, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	response, err := loginradius.PutManageAccountUpdateSecurityQuestionConfig(testuid, securityTest)
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
	_, _, _, testEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	email := TestEmail{testEmail}
	response, err := loginradius.PostManageForgotPasswordToken(email)
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
