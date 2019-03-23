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
	// data := new(MFAValidate)
	// req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/account/2fa", "")
	// if reqErr != nil {
	// 	return *data, reqErr
	// }

	// q := req.URL.Query()
	// q.Add("apikey", os.Getenv("APIKEY"))
	// q.Add("smstemplate2fa", smstemplate2fa)
	// req.URL.RawQuery = q.Encode()
	// req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	// req.Header.Add("Authorization", "Bearer "+authorization)

	// err := RunRequest(req, data)
	// return *data, err
}
