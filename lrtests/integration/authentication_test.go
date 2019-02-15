package integrationtest

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	loginradius "bitbucket.org/nombiezinja/lr-go-sdk"
	lrjson "bitbucket.org/nombiezinja/lr-go-sdk/json"
	"bitbucket.org/nombiezinja/lr-go-sdk/lrerror"
)

type Email struct {
	Type  string `json:"Type"`
	Value string `json:"Value"`
}

type User struct {
	Email    []Email `json:"Email"`
	Password string  `json:"Password"`
}

func TestPostAuthUserRegistrationByEmail(t *testing.T) {

	SetTestCredentials()
	testEmail := "lrtest" + strconv.FormatInt(time.Now().Unix(), 10) + "@mailinator.com"
	user := User{}

	res, err := loginradius.PostAuthUserRegistrationByEmail("", "", "", user)
	if err == nil || err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("PostAuthUserRegistrationByEmail Fail: Expected Error %v, instead received res: %+v, received error: %+v", "LoginradiusRespondedWithError", res, err)
	}

	user = User{
		Email: []Email{
			Email{
				Type:  "Primary",
				Value: testEmail,
			},
		},
		Password: "password",
	}

	res, err = loginradius.PostAuthUserRegistrationByEmail("", "", "", user)
	if res.StatusCode != 200 {
		t.Errorf("PostAuthUserRegistrationByEmail Success: Expected StatusCode %v, received %v", 200, res)
	}

	res, err = loginradius.PostAuthUserRegistrationByEmail("", "", "", user)
	if err == nil || err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("PostAuthUserRegistrationByEmail Fail: Expected Error %v, instead received res: %+v, received error: %+v", "LoginradiusRespondedWithError", res, err)
	}

	res, err = loginradius.GetManageAccountProfilesByEmail(testEmail)
	if err != nil {
		t.Errorf("Error retrieving uid of account to clean up.")
		fmt.Println(err)
	}

	profile, _ := lrjson.DynamicUnmarshal(res.Body)
	uid := profile["Uid"].(string)
	_, err = loginradius.DeleteManageAccount(uid)
	if err != nil {
		t.Errorf("Error cleaning up account.")
		fmt.Println(err)
	}
}

func TestPostAuthAddEmail(t *testing.T) {
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)

	testEmail := "lrtest" + strconv.FormatInt(time.Now().Unix(), 10) + "@mailinator.com"
	testAddEmail := TestEmailCreator{testEmail, "secondary"}
	res, err := loginradius.PostAuthAddEmail("", "", accessToken, testAddEmail)
	if err != nil {
		t.Errorf("Error making PostAuthAddEmail call: %v", err)
	}
	success, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !success["IsPosted"].(bool) {
		t.Errorf("Error returned from PostAuthAddEmail call: %v", err)
	}
}

func TestPostAuthAddEmailInvalid(t *testing.T) {
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	invalid := struct{ foo string }{"bar"}
	response, err := loginradius.PostAuthAddEmail("", "", accessToken, invalid)
	if err == nil {
		t.Errorf("Should fail but did not :%v", response.Body)
	}
}

func TestPostAuthForgotPassword(t *testing.T) {
	_, _, _, testEmail, _, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	testForgotPass := TestEmail{testEmail}
	res, err := loginradius.PostAuthForgotPassword("resetpassword.com", "", testForgotPass)
	if err != nil {
		t.Errorf("Error making PostAuthForgotPassword call: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from PostAuthForgotPassword call: %v", err)
	}
}

func TestPostAuthForgotPasswordInvalid(t *testing.T) {
	_, _, _, _, _, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	invalid := struct{ foo string }{"bar"}
	response, err := loginradius.PostAuthForgotPassword("resetpassword.com", "", invalid)
	if err == nil {
		t.Errorf("Should fail but did not: %v", response.Body)
	}
}

func TestPostAuthLoginByEmail(t *testing.T) {
	_, _, _, testEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	testLogin := TestEmailLogin{testEmail, testEmail}
	res, err := loginradius.PostAuthLoginByEmail("", "", "", "", "", testLogin)
	if err != nil {
		t.Errorf("Error making PostAuthLoginByEmail call: %v", err)
	}
	session, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostAuthLoginByEmail call: %v", err)
	}
}

func TestPostAuthLoginByEmailInvalid(t *testing.T) {
	_, _, _, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := struct{ foo string }{"bar"}
	response, err := loginradius.PostAuthLoginByEmail("", "", "", "", "", invalid)
	if err == nil {
		t.Errorf("Should fail but did not: %v", response.Body)
	}
}

func TestPostAuthLoginByUsername(t *testing.T) {
	_, userName, _, testEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)

	testLogin := TestUsernameLogin{userName, testEmail}
	res, err := loginradius.PostAuthLoginByUsername("", "", "", "", "", testLogin)
	if err != nil {
		t.Errorf("Error making PostAuthLoginByUsername call: %v", err)
	}
	session, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostAuthLoginByUsername call: %v", err)
	}
}

func TestPostAuthLoginByUsernameInvalid(t *testing.T) {
	_, _, _, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := struct{ foo string }{"bar"}
	response, err := loginradius.PostAuthLoginByUsername("", "", "", "", "", invalid)
	if err == nil {
		t.Errorf("Should fail but did not: %v", response.Body)
	}
}

func TestGetAuthCheckEmailAvailability(t *testing.T) {
	_, _, _, testEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	res, err := loginradius.GetAuthCheckEmailAvailability(testEmail)
	if err != nil {
		t.Errorf("Error making GetAuthCheckEmailAvailability call: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsExist"].(bool) {
		t.Errorf("Error returned from GetAuthCheckEmailAvailability call: %v", err)
	}
}

func TestGetAuthCheckUsernameAvailability(t *testing.T) {
	_, username, _, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	res, err := loginradius.GetAuthCheckUsernameAvailability(username)
	if err != nil {
		t.Errorf("Error making GetAuthCheckUsernameAvailability call: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsExist"].(bool) {
		t.Errorf("Error returned from GetAuthCheckUsernameAvailability call: %v", err)
	}
}

func TestGetAuthReadProfilesByToken(t *testing.T) {
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := loginradius.GetAuthReadProfilesByToken(accessToken)
	if err != nil {
		t.Errorf("Error making GetAuthReadProfilesByToken call: %v", err)
	}
	profile, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || profile["Uid"].(string) == "" {
		t.Errorf("Error returned from GetAuthReadProfilesByToken call: %v", err)
	}
}

// Test will fail if the feature Privacy Policy Versioning is not enabled through the dashboard
// To run test, comment out t.SkipNow()
func TestGetAuthPrivatePolicyAccept(t *testing.T) {
	t.SkipNow()
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := loginradius.GetAuthPrivatePolicyAccept(accessToken)
	if err != nil {
		t.Errorf("Error making GetAuthPrivatePolicyAccept call: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || data["Uid"].(string) == "" {
		t.Errorf("Error returned from GetAuthPrivatePolicyAccept call: %v", err)
	}
}

func TestGetAuthSendWelcomeEmail(t *testing.T) {
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := loginradius.GetAuthSendWelcomeEmail("", accessToken)
	if err != nil {
		t.Errorf("Error making GetAuthSendWelcomeEmail call: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from GetAuthSendWelcomeEmail call: %v", err)
	}
}

func TestGetAuthSocialIdentity(t *testing.T) {
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := loginradius.GetAuthSocialIdentity(accessToken)
	if err != nil {
		fmt.Println(err)
		t.Errorf("Error making GetAuthSocialIdentity call")
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || data["Uid"].(string) == "" {
		t.Errorf("Error returned from GetAuthSocialIdentity call")
		fmt.Println(err)
	}
}

func TestGetAuthSocialIdentityFail(t *testing.T) {
	SetTestCredentials()
	_, err := loginradius.GetAuthSocialIdentity("invalidtoken")
	if err == nil {
		t.Errorf("Should fail but did not.")
		fmt.Println(err)
	}
}

func TestGetAuthValidateAccessToken(t *testing.T) {
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := loginradius.GetAuthValidateAccessToken(accessToken)
	if err != nil {
		t.Errorf("Error making GetAuthValidateAccessToken call, %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || data["access_token"].(string) == "" {
		t.Errorf("Error returned from GetAuthValidateAccessToken call %v", err)
	}
}

func TestGetAuthVerifyEmail(t *testing.T) {
	_, _, verificationToken, teardownTestCase := setupEmailVerificationAccount(t)
	defer teardownTestCase(t)
	res, err := loginradius.GetAuthVerifyEmail(verificationToken, "", "")
	if err != nil {
		t.Errorf("Error making TestAuthVerifyEmail call, %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from TestAuthVerifyEmail call, %v", err)
	}
}

func TestGetAuthInvalidateAccessToken(t *testing.T) {
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := loginradius.GetAuthInvalidateAccessToken(accessToken)
	if err != nil {
		fmt.Println(err)
		t.Errorf("Error making GetAuthInvalidateAccessToken call")
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsPosted"].(bool) {
		fmt.Println(err)
		t.Errorf("Error returned from GetAuthInvalidateAccessToken call")
	}
}

// func TestGetAuthDeleteAccount(t *testing.T) {

// }

// Will return error unless security question feature is enabled
// Follow instructions in this document: https://docs.loginradius.com/api/v2/dashboard/platform-security/password-policy
func TestGetAuthSecurityQuestionByAccessToken(t *testing.T) {
	_, _, uid, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := loginradius.PutManageAccountUpdateSecurityQuestionConfig(uid, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question: %v", err)
	}
	response, err := loginradius.GetAuthSecurityQuestionByAccessToken(accessToken)
	if err != nil {
		t.Errorf("Error making GetAuthSecurityQuestionByAccessToken call: %v", err)
	}
	question := loginradius.AuthSecurityQuestion{}
	err = json.Unmarshal([]byte(response.Body), &question)
	if err != nil || (question[0].QuestionID == "") {
		t.Errorf("Error returned from GetAuthSecurityQuestionByUsername call: %v", err)
	}
}

func TestGetAuthSecurityQuestionByEmail(t *testing.T) {
	_, _, uid, email, _, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := loginradius.PutManageAccountUpdateSecurityQuestionConfig(uid, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question: %v", err)
	}
	response, err := loginradius.GetAuthSecurityQuestionByEmail(email)
	if err != nil {
		t.Errorf("Error making GetAuthSecurityQuestionByUsername call: %v", err)
	}
	question := loginradius.AuthSecurityQuestion{}
	err = json.Unmarshal([]byte(response.Body), &question)
	if err != nil || (question[0].QuestionID == "") {
		t.Errorf("Error returned from GetAuthSecurityQuestionByUsername call: %v", err)
	}
}

func TestGetAuthSecurityQuestionByUsername(t *testing.T) {
	_, username, uid, _, _, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := loginradius.PutManageAccountUpdateSecurityQuestionConfig(uid, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question: %v", err)
	}
	response, err := loginradius.GetAuthSecurityQuestionByUsername(username)
	if err != nil {
		t.Errorf("Error making GetAuthSecurityQuestionByUsername call: %v", err)
	}
	question := loginradius.AuthSecurityQuestion{}
	err = json.Unmarshal([]byte(response.Body), &question)
	if err != nil || (question[0].QuestionID == "") {
		t.Errorf("Error returned from GetAuthSecurityQuestionByUsername call: %v", err)
	}
}

func TestGetAuthSecurityQuestionByPhone(t *testing.T) {
	phone, _, uid, _, _, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := loginradius.PutManageAccountUpdateSecurityQuestionConfig(uid, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question: %v", err)
	}
	response, err := loginradius.GetAuthSecurityQuestionByPhone(phone)
	if err != nil {
		t.Errorf("Error making GetAuthSecurityQuestionByPhone call: %v", err)
	}
	question := loginradius.AuthSecurityQuestion{}
	err = json.Unmarshal([]byte(response.Body), &question)
	if err != nil || (question[0].QuestionID == "") {
		t.Errorf("Error returned from GetAuthSecurityQuestionByPhone call: %v", err)
	}
}

func TestPutAuthChangePassword(t *testing.T) {
	_, _, _, email, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	passwords := PassChange{email, email + "1"}
	resp, err := loginradius.PutAuthChangePassword(accessToken, passwords)
	if err != nil {
		t.Errorf("Error calling PutAuthChangePassword: %+v", err)
	}
	posted, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || !posted["IsPosted"].(bool) {
		t.Errorf("Error returned from PutAuthChangePassword: %+v", err)
	}
}

func TestPutResendEmailVerification(t *testing.T) {
	_, retEmail, _, teardownTestCase := setupEmailVerificationAccount(t)
	defer teardownTestCase(t)
	emailRef := TestEmail{retEmail}
	resp, err := loginradius.PutResendEmailVerification("", "", emailRef)
	if err != nil {
		t.Errorf("Error calling PutResendEmailVerification: %v", err)
	}
	posted, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || !posted["IsPosted"].(bool) {
		t.Errorf("Error returned for PutResendEmailVerification: %v", err)
	}
}

func TestPutAuthResetPasswordByResetToken(t *testing.T) {
	_, _, _, email, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)

	resetEmail := TestEmail{email}
	response, err := loginradius.PostManageForgotPasswordToken(resetEmail)
	if err != nil {
		t.Errorf(
			"Error calling PostManageForgotPasswordToken for PutAuthResetPasswordByResetToken: %v",
			err,
		)
	}
	data, _ := lrjson.DynamicUnmarshal(response.Body)
	req := PasswordReset{data["ForgotToken"].(string), email + "1"}
	response, err = loginradius.PutAuthResetPasswordByResetToken(req)
	if err != nil {
		t.Errorf("Error calling PutAuthResetPasswordByResetToken: %v", err)
	}
	data, err = lrjson.DynamicUnmarshal(response.Body)
	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from PutAuthResetPasswordByResetToken: %+v", err)
	}
}

// func TestPutAuthResetPasswordByOTP(t *testing.T) {

// }

func TestPutAuthResetPasswordBySecurityAnswerAndEmail(t *testing.T) {
	_, _, uid, email, _, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)

	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := loginradius.PutManageAccountUpdateSecurityQuestionConfig(uid, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question: %v", err)
	}

	request := ResetWithEmailSecurity{securityQuestion, email, email + "1", ""}
	response, err := loginradius.PutAuthResetPasswordBySecurityAnswerAndEmail(request)
	if err != nil {
		t.Errorf("Error making call to PutAuthResetPasswordBySecurityAnswerAndEmail: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from call to PutAuthResetPasswordBySecurityAnswerAndEmail: %+v", err)
	}
}

func TestPutAuthResetPasswordBySecurityAnswerAndUsername(t *testing.T) {
	_, username, uid, email, _, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)

	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := loginradius.PutManageAccountUpdateSecurityQuestionConfig(uid, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question: %v", err)
	}

	request := ResetWithUsernameSecurity{securityQuestion, username, email + "1", ""}
	response, err := loginradius.PutAuthResetPasswordBySecurityAnswerAndUsername(request)
	if err != nil {
		t.Errorf("Error making call to PutAuthResetPasswordBySecurityAnswerAndUsername: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from PutAuthResetPasswordBySecurityAnswerAndUsername: %+v", err)
	}
}

func TestPutAuthSetOrChangeUsername(t *testing.T) {
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	newName := TestUsername{"newusername"}
	_, err := loginradius.PutAuthSetOrChangeUsername(accessToken, newName)
	if err != nil {
		t.Errorf("Error making call to PutAuthSetOrChangeUsername: %+v", err)
	}
	response, err := loginradius.GetAuthReadProfilesByToken(accessToken)
	if err != nil {
		t.Errorf("Error making call to GetAuthReadProfilesByToken for PutAuthSetOrChangeUsername: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil {
		t.Errorf("Error returned from GetAuthReadProfilesByToken for PutAuthSetOrChangeUsername: %+v", err)
	}
	if data["UserName"].(string) != "newusername" {
		t.Errorf("PutAuthSetOrChangeUsername failed, expected username NewUserName, but instead got: %v", data["UserName"].(string))
	}
}

func TestPutAuthUpdateProfileByToken(t *testing.T) {
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	request := TestUsername{"newname"}
	_, err := loginradius.PutAuthUpdateProfileByToken("", "", "", accessToken, request)
	if err != nil {
		t.Errorf("Error making call to PutAuthUpdateProfileByToken: %+v", err)
	}
	response, err := loginradius.GetAuthReadProfilesByToken(accessToken)
	if err != nil {
		t.Errorf("Error making call to GetAuthReadProfilesByToken for PutAuthUpdateProfileByToken: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil {
		t.Errorf("Error returned from GetAuthReadProfilesByToken for PutAuthUpdateProfileByToken: %+v", err)
	}
	if data["UserName"].(string) != "newname" {
		t.Errorf("PutAuthSetOrChangeUsername failed, expected username NewUserName, but instead got: %v", data["UserName"].(string))
	}
}

func TestPutAuthUpdateSecurityQuestionByAccessToken(t *testing.T) {
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := loginradius.PutAuthUpdateSecurityQuestionByAccessToken(accessToken, securityTest)
	if err != nil {
		t.Errorf("Error making PutAuthUpdateSecurityQuestionByAccessToken call: %v", err)
	}
}

func TestDeleteAuthDeleteAccountEmailConfirmation(t *testing.T) {
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	resp, err := loginradius.DeleteAuthDeleteAccountEmailConfirmation("", "", accessToken)
	if err != nil {
		t.Errorf("Error making call to DeleteAuthDeleteAccountEmailConfirmation: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || !data["IsDeleteRequestAccepted"].(bool) {
		t.Errorf("Error returned from DeleteAuthDeleteAccountEmailConfirmation: %+v", err)
	}
}

func TestDeleteAuthRemoveEmail(t *testing.T) {
	_, _, _, testEmail, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	removeEmail := TestEmail{testEmail}
	resp, err := loginradius.DeleteAuthRemoveEmail(accessToken, removeEmail)
	if err != nil {
		t.Errorf("Error making call to DeleteAuthRemoveEmail: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || !data["IsDeleted"].(bool) {
		t.Errorf("Error returned from call to DeleteAuthRemoveEmail: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), and configure secret.env with valid user access token
// Pre-create the user used for this test and link an account of a social provider; configure the
// string of this social provider in the secret.env with lower case names
// e.g.PROVIDER=google, PROVIDER=facebook
func TestDeleteAuthUnlinkSocialIdentities(t *testing.T) {
	t.SkipNow()
	SetTestCredentials()
	accessToken := os.Getenv("USERTOKEN")
	response, err := loginradius.GetAuthReadProfilesByToken(accessToken)
	if err != nil {
		t.Errorf("Error making call to GetAuthReadProfilesByToken: %+v", err)
	}

	data, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil {
		t.Errorf("Error parsing response from GetAuthReadProfilesByToken: %+v", err)
	}
	identities, ok := data["Identities"].([]interface{})
	if !ok {
		fmt.Println("Identities returned is null, not array")
		return
	}

	var id string
	providerstr := os.Getenv("PROVIDER")
	for _, v := range identities {
		asserted := v.(map[string]interface{})
		if asserted["Provider"] == providerstr {
			id = asserted["ID"].(string)
		}
	}

	provider := Provider{providerstr, id}

	response, err = loginradius.DeleteAuthUnlinkSocialIdentities(accessToken, provider)
	if err != nil {
		t.Errorf("Error making call to DeleteAuthUnlinkSocialIdentities: %+v", err)
	}

	deleted, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || !deleted["IsDeleted"].(bool) {
		t.Errorf("Error returned from DeleteAuthUnlinkSocialIdentities: %+v", err)
	}
}

func TestGetPasswordlessLoginByEmail(t *testing.T) {
	_, _, _, email, _, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	response, err := loginradius.GetPasswordlessLoginByEmail(email, "", "")
	if err != nil {
		t.Errorf("Error making call to GetPasswordlessLoginByEmail: %+v", err)
	}
	posted, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || !posted["IsPosted"].(bool) {
		t.Errorf("Error returned from GetPasswordlessLoginByEmail: %+v", err)
	}
}

func TestGetPasswordlessLoginByUsername(t *testing.T) {
	_, username, _, _, _, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	response, err := loginradius.GetPasswordlessLoginByUsername(username, "", "")
	if err != nil {
		t.Errorf("Error making call to GetPasswordlessLoginByUsername: %+v", err)
	}
	posted, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || !posted["IsPosted"].(bool) {
		t.Errorf("Error returned from GetPasswordlessLoginByUsername: %+v", err)
	}
}
