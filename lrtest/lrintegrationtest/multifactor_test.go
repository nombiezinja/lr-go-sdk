package lrintegrationtest

import (
	"testing"

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

// To run this test, uncomment t.SkipNow() and manually set a manually created user with mfa turned on
func TestPutMFAValidateGoogleAuthCode(t *testing.T) {
	// t.SkipNow()

}
