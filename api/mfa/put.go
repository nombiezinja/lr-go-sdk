package mfa

import "bitbucket.org/nombiezinja/lr-go-sdk/httprutils"

// PutMFAValidateGoogleAuthCode is used to login via Multi-factor-authentication by passing the google authenticator code.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-validate-google-auth-code
// Required query parameters: apikey, secondfactorauthenticationtoken
// secondfactorauthenticationtoken can be obtained by successful logins through MFA login routes
// Optional query parameters: smstemplate2fa
// Required post parameter: googleauthenticatorcode: string
func (lr Loginradius) PutMFAValidateGoogleAuthCode(queries interface{}, body interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPutReq("/identity/v2/auth/login/2fa/verification/googleauthenticatorcode", body)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
	// data := new(MFALogin)
	// req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/login/2fa/verification/googleauthenticatorcode", body)
	// if reqErr != nil {
	// 	return *data, reqErr
	// }

	// q := req.URL.Query()
	// q.Add("apikey", os.Getenv("APIKEY"))
	// q.Add("secondfactorauthenticationtoken", secondFactorAuthenticationToken)
	// q.Add("smstemplate2fa", smstemplate2fa)
	// req.URL.RawQuery = q.Encode()
	// req.Header.Add("content-Type", "application/json")

	// err := RunRequest(req, data)
	// return *data, err
}
