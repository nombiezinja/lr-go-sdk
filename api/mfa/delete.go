package mfa

import (
	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	lrvalidate "bitbucket.org/nombiezinja/lr-go-sdk/internal/validate"
)

// DeleteMFAResetGoogleAuthenticatorByToken resets the Google Authenticator configurations on a given account via the access_token.
// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-reset-google-authenticator-by-token
// Required query parameter: apikey
// Required body parameter: googleauthenticator - pass true as value
func (lr Loginradius) DeleteMFAResetGoogleAuthenticatorByToken() (*httprutils.Response, error) {
	req, err := lr.Client.NewDeleteReqWithToken(
		"/identity/v2/auth/account/2fa/authenticator",
		map[string]bool{"googleauthenticator": true},
	)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	req.Headers = httprutils.JSONHeader
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// DeleteMFAResetSMSAuthenticatorByToken resets the SMS Authenticator configurations on a given account via the access_token.
// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-reset-sms-authenticator-by-token
// Required query parameter: apikey
// Required body parameter: otpauthenticator - pass true as value
func (lr Loginradius) DeleteMFAResetSMSAuthenticatorByToken() (*httprutils.Response, error) {
	req, err := lr.Client.NewDeleteReqWithToken(
		"/identity/v2/auth/account/2fa/authenticator",
		map[string]bool{"otpauthenticator": true},
	)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	req.Headers = httprutils.JSONHeader
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// DeleteMFAResetSMSAuthenticatorByUid resets the SMS Authenticator configurations on a given account via the access_token.
// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-reset-sms-authenticator-by-uid
// Required query parameter: apikey, apisecret, uid
// Required body parameter: otpauthenticator - pass true as value
func (lr Loginradius) DeleteMFAResetSMSAuthenticatorByUid(queries interface{}) (*httprutils.Response, error) {
	queryParams := map[string]string{}
	uid, ok := queries.(string)
	if ok {
		queryParams["uid"] = uid
	} else {
		allowedQueries := map[string]bool{"uid": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
		if err != nil {
			return nil, err
		}
		queryParams = validatedQueries
	}

	req, err := lr.Client.NewDeleteReqWithToken(
		"/identity/v2/auth/account/2fa/authenticator",
		map[string]bool{"otpauthenticator": true},
		queryParams,
	)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(req)
	req.Headers = httprutils.JSONHeader
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// DeleteMFAResetGoogleAuthenticatorByUid resets the SMS Authenticator configurations on a given account via the access_token.
// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-reset-google-authenticator-by-uid
// Required query parameter: apikey, apisecret, uid
// Required body parameter: otpauthenticator - pass true as value
func (lr Loginradius) DeleteMFAResetGoogleAuthenticatorByUid(queries interface{}) (*httprutils.Response, error) {
	queryParams := map[string]string{}
	uid, ok := queries.(string)
	if ok {
		queryParams["uid"] = uid
	} else {
		allowedQueries := map[string]bool{"uid": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
		if err != nil {
			return nil, err
		}
		queryParams = validatedQueries
	}

	req, err := lr.Client.NewDeleteReqWithToken(
		"/identity/v2/auth/account/2fa/authenticator",
		map[string]bool{"googleauthenticator": true},
		queryParams,
	)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(req)
	req.Headers = httprutils.JSONHeader
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}
