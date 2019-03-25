package lrintegrationtest

import (
	"os"
	"testing"

	lr "bitbucket.org/nombiezinja/lr-go-sdk"
	"bitbucket.org/nombiezinja/lr-go-sdk/api/mfa"
	"bitbucket.org/nombiezinja/lr-go-sdk/lrerror"
	lrjson "bitbucket.org/nombiezinja/lr-go-sdk/lrjson"
)

// func setupMFALogin(t *testing.T) (string, string, string, string, func(t *testing.T)) {
// 	_, _, testuid, testEmail, teardownTestCase := setupAccount(t)
// 	testLogin := TestEmailLogin{testEmail, testEmail}
// 	session, err := PostMFAEmailLogin("", "", "", "", testLogin)
// 	accessToken := session.AccessToken
// 	multiToken := session.SecondFactorAuthentication.SecondFactorAuthenticationToken
// 	if err != nil || accessToken == "" {
// 		t.Errorf("Error logging in")
// 		fmt.Println(err)
// 	}
// 	return multiToken, testuid, testEmail, accessToken, func(t *testing.T) {
// 		defer teardownTestCase(t)
// 	}
// }

func TestPostMFAEmailLogin(t *testing.T) {
	_, _, _, testEmail, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	testLogin := TestEmailLogin{testEmail, testEmail}
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(testLogin)
	if err != nil {
		t.Errorf("Error making PostMFAEmailLogin call: %v", err)
	}
	session, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostMFAEmailLogin call: %v", err)
	}

	res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(testLogin, map[string]string{"emailtemplate": "hello"})

	if err != nil {
		t.Errorf("Error making PostMFAEmailLogin call with optional queries: %v", err)
	}
	session, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostMFAEmailLogin call with optional queries: %v", err)
	}
}

func TestPostMFAEmailLoginInvalidBody(t *testing.T) {
	_, _, _, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := struct{ foo string }{"bar"}
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(invalid)
	if err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("PostMFAEmailLogin should fail with LoginradiusRespondedWithError but did not: %v", res.Body)
	}
}

func TestPostMFAEmailLoginInvalidQuery(t *testing.T) {
	_, _, _, email, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	user := TestEmailLogin{email, email}
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(user, map[string]string{"invalidparam": "value"})
	if err.(lrerror.Error).Code() != "ValidationError" {
		t.Errorf("PostMFAEmailLogin should fail with ValidationError but did not :%v, %+v", res.Body, err)
	}
}

func TestPostMFAUsernameLogin(t *testing.T) {
	_, username, _, password, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAUsernameLogin(
		map[string]string{"username": username, "password": password},
	)
	if err != nil {
		t.Errorf("Error making PostMFAUsernameLogin call: %v", err)
	}
	session, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostMFAUsernameLogin call: %v", err)
	}

	res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAUsernameLogin(
		map[string]string{"username": username, "password": password},
		map[string]string{"emailtemplate": "hello"},
	)

	if err != nil {
		t.Errorf("Error making PostMFAUsernameLogin call with optional queries: %v", err)
	}
	session, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostMFAUsernameLogin call with optional queries: %v", err)
	}
}

func TestPostMFAUsernameLoginInvalidBody(t *testing.T) {
	_, _, _, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := struct{ foo string }{"bar"}
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAUsernameLogin(invalid)
	if err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("PostMFAUsernameLogin should fail with LoginradiusRespondedWithError but did not: %v", res.Body)
	}
}

func TestPostMFAUsernameLoginInvalidQuery(t *testing.T) {
	_, username, _, password, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAUsernameLogin(
		map[string]string{"username": username, "password": password},
		map[string]string{"invalidparam": "value"},
	)
	if err.(lrerror.Error).Code() != "ValidationError" {
		t.Errorf("PostMFAUsernameLogin should fail with ValidationError but did not :%v, %+v", res.Body, err)
	}
}

func TestPostMFAPhoneLogin(t *testing.T) {
	phone, _, _, password, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAPhoneLogin(
		map[string]string{"phone": phone, "password": password},
	)
	if err != nil {
		t.Errorf("Error making PostMFAPhoneLogin call: %v", err)
	}
	session, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostMFAPhoneLogin call: %v", err)
	}

	res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAPhoneLogin(
		map[string]string{"phone": phone, "password": password},
		map[string]string{"emailtemplate": "hello"},
	)

	if err != nil {
		t.Errorf("Error making PostMFAPhoneLogin call with optional queries: %v", err)
	}
	session, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostMFAPhoneLogin call with optional queries: %v", err)
	}
}

func TestPostMFAPhoneLoginInvalidBody(t *testing.T) {
	_, _, _, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := struct{ foo string }{"bar"}
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAPhoneLogin(invalid)
	if err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("PostMFAPhoneLogin should fail with LoginradiusRespondedWithError but did not: %v", res.Body)
	}
}

func TestPostMFAPhoneLoginInvalidQuery(t *testing.T) {
	phone, _, _, password, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAPhoneLogin(
		map[string]string{"phone": phone, "password": password},
		map[string]string{"invalidparam": "value"},
	)
	if err.(lrerror.Error).Code() != "ValidationError" {
		t.Errorf("PostMFAPhoneLogin should fail with ValidationError but did not :%v, %+v", res.Body, err)
	}
}

func TestGetMFAValidateAccessToken(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFAValidateAccessToken()
	if err != nil {
		t.Errorf("Error making call to MFAValidateAccessToken: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || data["QRCode"].(string) == "" {
		t.Errorf("Error returned from MFAValidateAccessToken: %v", err)
	}

	res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFAValidateAccessToken(map[string]string{"smstemplate2fa": "hello"})
	if err != nil {
		t.Errorf("Error making call to MFAValidateAccessToken with optional query params: %v", err)
	}
	data, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil || data["QRCode"].(string) == "" {
		t.Errorf("Error returned from MFAValidateAccessToken with optional query params: %v", err)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// and a Google authenticator added, enter the google authenticator code in this test.
func TestPutMFAValidateGoogleAuthCode(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(
		// Set user credentials here
		map[string]string{"email": "", "password": ""},
	)
	if err != nil {
		t.Errorf("Error making PostMFAEmailLogin call for PutMFAValidateGoogleAuthCode: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil {
		t.Errorf("Error returned from PostMFAEmailLogin call for PutMFAValidateGoogleAuthCode: %v", err)
	}

	code, ok := data["SecondFactorAuthentication"].(map[string]interface{})["SecondFactorAuthenticationToken"].(string)
	if !ok {
		t.Errorf("Returned response from SecondFactorAuthentication does not contain SecondFactorAuthenticationToken")
	}

	res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAValidateGoogleAuthCode(
		map[string]string{"secondfactorauthenticationtoken": code},
		// Set otp from Google Authenticator here
		map[string]string{"googleauthenticatorcode": "246803"},
	)
	if err != nil {
		t.Errorf("Error making call to PutMFAValidateGoogleAuthCode: %v", err)
	}
	data, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil || data["access_token"].(string) == "" {
		t.Errorf("Error returned from PutMFAValidateGoogleAuthCode: %v", err)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// this test tests for the ability to submit a valid request to the LoginRadius end point
// and will pass if a ""The OTP code is invalid, please request for a new OTP" error is returned
// from Loginradius
func TestPutMFAValidateOTP(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(
		// Set user credentials here
		map[string]string{"email": "blueberries@mailinator.com", "password": "password"},
	)
	if err != nil {
		t.Errorf("Error making PostMFAEmailLogin call for PutMFAValidateOTP: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil {
		t.Errorf("Error returned from PostMFAEmailLogin call for PutMFAValidateOTP: %v", err)
	}

	code, ok := data["SecondFactorAuthentication"].(map[string]interface{})["SecondFactorAuthenticationToken"].(string)
	if !ok {
		t.Errorf("Returned response from PutMFAValidateOTP does not contain SecondFactorAuthenticationToken")
	}

	_, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAValidateOTP(
		map[string]string{"secondfactorauthenticationtoken": code},
		map[string]string{"otp": "123456"},
	)

	errMsg, _ := lrjson.DynamicUnmarshal(err.(lrerror.Error).OrigErr().Error())

	if errMsg["Description"].(string) != "The OTP code is invalid, please request for a new OTP." {
		t.Errorf("PutMFAValidateOTP was supposed to return invalid OTP error, but did not: %v", errMsg)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// then obtain a valid secondfactorauthenticationtoken through completing a mfa login attempt
// set the secondfactorauthenticationtoken and a phone number here
func TestPutMFAUpdatePhoneNumber(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(
		// Set user credentials here
		map[string]string{"email": "blueberries@mailinator.com", "password": "password"},
	)
	if err != nil {
		t.Errorf("Error making PostMFAEmailLogin call for PutMFAUpdatePhoneNumber: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil {
		t.Errorf("Error returned from PostMFAEmailLogin call for PutMFAUpdatePhoneNumber: %v", err)
	}

	code, ok := data["SecondFactorAuthentication"].(map[string]interface{})["SecondFactorAuthenticationToken"].(string)
	if !ok {
		t.Errorf("Returned response from SecondFactorAuthentication does not contain SecondFactorAuthenticationToken")
	}

	res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAUpdatePhoneNumber(
		// Set user here
		map[string]string{"secondfactorauthenticationtoken": code},
		map[string]string{"phoneno2fa": ""},
	)
	if err != nil {
		t.Errorf("Error making call to PutMFAUpdatePhoneNumber: %v", err)
	}
	data, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil {
		t.Errorf("Error returned from PutMFAUpdatePhoneNumber: %v", err)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// then obtain a valid access_token through completing a mfa login attempt
// set the access_token and a phone number here
func TestPutMFAUpdatePhoneNumberByToken(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// set valid access_token here
	lrclient.Context.Token = "7f875c92-b7fe-4f55-8658-58b24387ed64"
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAUpdatePhoneNumberByToken(
		// Set user here
		map[string]string{"phoneno2fa": "16047711536"},
	)
	if err != nil {
		t.Errorf("Error making call to PutMFAUpdatePhoneNumber: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || data["Sid"] == "" {
		t.Errorf("Error returned from PutMFAUpdatePhoneNumber: %v", err)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// then obtain a valid access_token through completing a mfa login attempt
func TestGetMFABackUpCodeByAccessToken(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// set valid access_token here
	lrclient.Context.Token = "77aa9464-815c-4dbe-8eec-c6c9e28e43b2"
	_, err := mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFABackUpCodeByAccessToken()
	if err != nil {
		t.Errorf("Error making call to GetMFABackUpCodeByAccessToken: %v", err)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// then obtain a valid access_token through completing a mfa login attempt
func TestGetMFAResetBackUpCodeByAccessToken(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// set valid access_token here
	lrclient.Context.Token = "77aa9464-815c-4dbe-8eec-c6c9e28e43b2"
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFAResetBackUpCodeByAccessToken()
	if err != nil {
		t.Errorf("Error making call to GetMFAResetBackUpCodeByAccessToken: %v", err)
	}

	codes, err := lrjson.DynamicUnmarshal(res.Body)
	_, ok := codes["BackUpCodes"].([]interface{})
	if err != nil || !ok {
		t.Errorf("Error returned from :%v, %v", err, codes)
	}
}
