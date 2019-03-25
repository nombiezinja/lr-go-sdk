package mfa

import (
	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	lrvalidate "bitbucket.org/nombiezinja/lr-go-sdk/internal/validate"
)

// GetMFAValidateAccessToken is used to configure the Multi-factor authentication
// after login by using the access_token when MFA is set as optional on the LoginRadius site.
// https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-validate-access-token
// Required query parameter: apikey
// Optional query parameter: smstemplate2fa
// Needs Authorization Bearer token header
func (lr Loginradius) GetMFAValidateAccessToken(queries ...interface{}) (*httprutils.Response, error) {
	queryParams := map[string]string{}

	for _, arg := range queries {
		allowedQueries := map[string]bool{"smstemplate2fa": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)
		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			queryParams[k] = v
		}
	}
	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/account/2fa", queryParams)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetMFABackUpCodeByAccessToken is used to get a set of backup codes via access_token to allow the user login on a site that has Multi-factor Authentication enabled in the event that the user does not have a secondary factor available. We generate 10 codes, each code can only be consumed once. If any user attempts to go over the number of invalid login attempts configured in the Dashboard then the account gets blocked automatically
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-backup-code-by-access-token
// Required query parameter: apikey
func (lr Loginradius) GetMFABackUpCodeByAccessToken() (*httprutils.Response, error) {
	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/account/2fa/backupcode")
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

//GetMFABackUpCodeByAccessToken is used to reset the backup codes on a given account via the access_token. This API call will generate 10 new codes, each code can only be consumed once.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-reset-backup-code-by-access-token
// Required query parameter: apikey
func (lr Loginradius) GetMFAResetBackUpCodeByAccessToken() (*httprutils.Response, error) {
	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/account/2fa/backupcode/reset")
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}
