package mfa

import (
	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	lrvalidate "bitbucket.org/nombiezinja/lr-go-sdk/internal/validate"
)

// PutMFAValidateGoogleAuthCode is used to login via Multi-factor-authentication by passing the google authenticator code.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-validate-google-auth-code
// Required query parameters: apikey, secondfactorauthenticationtoken
// secondfactorauthenticationtoken can be obtained by successful logins through MFA login routes
// Optional query parameters: smstemplate2fa
// Required post parameter: googleauthenticatorcode: string
func (lr Loginradius) PutMFAValidateGoogleAuthCode(queries, body interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"secondfactorauthenticationtoken": true, "smstemplate2fa": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewPutReq("/identity/v2/auth/login/2fa/verification/googleauthenticatorcode", body, validatedQueries)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutMFAValidateOTP is used to login via Multi-factor authentication by passing the One Time Password received via SMS.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-validate-otp
// Required query parameters: apikey, secondfactorauthenticationtoken
// Optional query parameter: smstemplate2fa
// Required post parameter: otp - string
// Optional query parameters: securityanswer, g-recaptcha-response, qq_captcha_ticket, qq_captcha_randstr
func (lr Loginradius) PutMFAValidateOTP(queries, body interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"secondfactorauthenticationtoken": true, "smstemplate2fa": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewPutReq("/identity/v2/auth/login/2fa/verification/otp", body, validatedQueries)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutMFAUpdatePhoneNumber is used to update (if configured) the phone number used for Multi-factor authentication by sending the verification OTP to the provided phone number.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-update-phone-number
// Required query parameters: apikey, secondfactorauthenticationtoken
// Optional query parameter: smstemplate2fa
// Required post parameter: phoneno2fa - string
func (lr Loginradius) PutMFAUpdatePhoneNumber(queries, body interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"secondfactorauthenticationtoken": true, "smstemplate2fa": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewPutReq("/identity/v2/auth/login/2fa", body, validatedQueries)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutMFAUpdatePhoneNumberByToken is used to update (if configured) the phone number used for Multi-factor authentication by sending the verification OTP to the provided phone number.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-update-phone-number
// Required query parameters: apikey
// Optional query parameter: smstemplate2fa
// Required post parameter: phoneno2fa - string
func (lr Loginradius) PutMFAUpdatePhoneNumberByToken(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/account/2fa", body)
	if err != nil {
		return nil, err
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"smstemplate2fa": true,
		}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			req.QueryParams[k] = v
		}
	}

	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// This API is used to validate the backup code provided by the user and if valid, we return an access_token allowing the user to login incases where Multi-factor authentication (MFA) is enabled and the secondary factor is unavailable. When a user intially downloads the Backup codes, We generate 10 codes, each code can only be consumed once. if any user attempts to go over the number of invalid login attempts configured in the Dashboard then the account gets blocked automatically
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-validate-backup-code
// Required query parameters: apikey, secondfactorauthenticationtoken
// Required body parameter: backupcode
func (lr Loginradius) PutMFAValidateBackupCode(queries interface{}, body interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"secondfactorauthenticationtoken": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewPutReq("/identity/v2/auth/login/2fa/verification/backupcode", body, validatedQueries)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}
