package lrintegrationtest

import (
	"os"
	"testing"

	lr "bitbucket.org/nombiezinja/lr-go-sdk"
	"bitbucket.org/nombiezinja/lr-go-sdk/api/phoneauthentication"
	lrjson "bitbucket.org/nombiezinja/lr-go-sdk/lrjson"
)

// func unverifyPhone(uid string) {
// 	falsePhoneIDVerified := UndoPhoneVerify{false}
// 	PutManageAccountUpdate(uid, falsePhoneIDVerified)
// }

func TestPostPhoneLogin(t *testing.T) {
	phoneID, _, _, password, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PostPhoneLogin(
		map[string]string{"phone": phoneID, "password": password},
	)
	if err != nil {
		t.Errorf("Error calling PostPhoneLogin: %v", err)
	}
	profile, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || profile["Uid"] == "" {
		t.Errorf("Error returned from PostPhoneLogin: %v", err)
	}
}

// To run this test comment out t.SkipNow() and set PHONENUMBER in secret.env
// with a valid phone number of a manually created user profile
func TestPostPhoneForgotPasswordByOTP(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)
	resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PostPhoneForgotPasswordByOTP(
		map[string]string{"phone": os.Getenv("PHONENUMBER")},
	)
	if err != nil {
		t.Errorf("Error calling PostPhoneForgotPasswordByOTP: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(resp.Body)

	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from PostPhoneForgotPasswordByOTP: %v", err)
	}
}

// To run this test comment out t.SkipNow() and set PHONENUMBER in secret.env
// with a valid phone number of a manually created user profile
func TestPostPhoneResendVerificationOTP(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)
	resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PostPhoneResendVerificationOTP(
		map[string]string{"phone": os.Getenv("PHONENUMBER")},
	)
	if err != nil {
		t.Errorf("Error calling ostPhoneResendVerificationOTP: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(resp.Body)

	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from ostPhoneResendVerificationOTP: %v", err)
	}
}

// To run this test comment out t.SkipNow() and set PHONENUMBER and USERTOKEN in secret.env
// with a valid phone number of a manually created user profile
func TestPostPhoneResendVerificationOTPByToken(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg, map[string]string{"token": os.Getenv("USERTOKEN")})
	resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PostPhoneResendVerificationOTPByToken(
		map[string]string{"phone": os.Getenv("PHONENUMBER")},
	)
	if err != nil {
		t.Errorf("Error calling ostPhoneResendVerificationOTP: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(resp.Body)

	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from ostPhoneResendVerificationOTP: %v", err)
	}
}

// func TestPostPhoneUserRegistrationBySMS(t *testing.T) {
// 	PresetLoginRadiusTestEnv()
// 	fmt.Println("Starting test TestPostPhoneUserRegistrationBySMS")
// 	time := time.Now()
// 	timestamp := time.Format("20060102150405")
// 	timestampEmail := "testemail" + timestamp + "@mailinator.com"
// 	testEmails := TestEmailArr{{"Primary", timestampEmail}}
// 	phoneAccount := PhoneRegister{testEmails, "+12016768872", "password"}
// 	session, err := PostPhoneUserRegistrationBySMS("", "", "", phoneAccount)
// 	if err != nil && session.IsPosted != true {
// 		t.Errorf("Error registering phone number")
// 		fmt.Println(err)
// 	}
// 	user, err2 := GetManageAccountProfilesByEmail(timestampEmail)
// 	if err2 != nil {
// 		t.Errorf("Error cleaning up account")
// 		fmt.Println(err2)
// 	}
// 	uid := user.UID
// 	_, err3 := DeleteManageAccount(uid)
// 	if err3 != nil {
// 		t.Errorf("Error cleaning up account")
// 		fmt.Println(err3)
// 	}
// 	fmt.Println("Test complete")
// }

// func TestGetPhoneNumberAvailability(t *testing.T) {
// 	fmt.Println("Starting test TestGetPhoneNumberAvailability")
// 	phoneID, _, _, _, teardownTestCase := setupAccount(t)
// 	defer teardownTestCase(t)
// 	_, err := GetPhoneNumberAvailability(phoneID)
// 	if err != nil {
// 		t.Errorf("Error checking phone number availability")
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Test complete")
// }

// func TestPutPhoneNumberUpdate(t *testing.T) {
// 	fmt.Println("Starting test TestPutPhoneNumberUpdate")
// 	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
// 	defer teardownTestCase(t)
// 	phone := TestPhone{"+12016768874"}
// 	session, err := PutPhoneNumberUpdate("", accessToken, phone)
// 	if err != nil && session.IsPosted != true {
// 		t.Errorf("Error updating phone number")
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Test complete")
// }

// func TestPutResetPhoneIDVerification(t *testing.T) {
// 	fmt.Println("Starting test TestPutResetPhoneIDVerification")
// 	_, _, testuid, _, teardownTestCase := setupAccount(t)
// 	defer teardownTestCase(t)
// 	_, err := PutResetPhoneIDVerification(testuid)
// 	if err != nil {
// 		t.Errorf("Error resetting verification")
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Test complete")
// }

// func TestDeleteRemovePhoneIDByAccessToken(t *testing.T) {
// 	fmt.Println("Starting test TestDeleteRemovePhoneIDByAccessToken")
// 	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
// 	defer teardownTestCase(t)
// 	_, err := DeleteRemovePhoneIDByAccessToken(accessToken)
// 	if err != nil {
// 		t.Errorf("Error removing phone ID")
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Test complete")
// }
