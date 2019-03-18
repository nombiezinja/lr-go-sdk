package multifactor

import (
	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	lrvalidate "bitbucket.org/nombiezinja/lr-go-sdk/internal/validate"
)

// PostMFAEmailLogin can be used to login by emailid on a Multi-factor authentication enabled LoginRadius site.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-email-login
// Required query parameter: apikey
// Optional query parameters: loginurl, verificationurl, emailtemplate, smstemplate2fa
// Required post parameters: email - string; password - string;
func (lr Loginradius) PostMFAEmailLogin(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPostReq("/identity/v2/auth/login/2fa", body)
	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"verificationurl": true, "loginurl": true, "emailtemplate": true, "smstemplate2fa": true,
		}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			request.QueryParams[k] = v
		}
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
	// data := new(MFAPost)
	// req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/auth/login/2fa", body)
	// if reqErr != nil {
	// 	return *data, reqErr
	// }

	// q := req.URL.Query()
	// q.Add("apikey", os.Getenv("APIKEY"))
	// q.Add("loginurl", loginURL)
	// q.Add("verificationurl", verificationURL)
	// q.Add("emailtemplate", emailTemplate)
	// q.Add("smstemplate2fa", smstemplate2fa)
	// req.URL.RawQuery = q.Encode()
	// req.Header.Add("content-Type", "application/json")

	// err := RunRequest(req, data)
	// return *data, err
}

// PostMFAUsernameLogin can be used to login by username on a Multi factor authentication enabled LoginRadius site.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-user-name-login
// Required query parameter: apikey
// Optional query parameters: loginurl, verificationurl, emailtemplate, smstemplate2fa
// Required post parameters: username - string; password - string;
func (lr Loginradius) PostMFAUsernameLogin(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPostReq("/identity/v2/auth/login/2fa", body)
	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"verificationurl": true, "loginurl": true, "emailtemplate": true, "smstemplate2fa": true,
		}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			request.QueryParams[k] = v
		}
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
	// smsTemplate, smstemplate2fa string, body interface{}) (MFAPost, error) {
	// 	data := new(MFAPost)
	// 	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/auth/login/2fa", body)
	// 	if reqErr != nil {
	// 		return *data, reqErr
	// 	}

	// 	q := req.URL.Query()
	// 	q.Add("apikey", os.Getenv("APIKEY"))
	// 	q.Add("loginurl", loginURL)
	// 	q.Add("verificationurl", verificationURL)
	// 	q.Add("emailtemplate", emailTemplate)
	// 	q.Add("smsTemplate", smsTemplate)
	// 	q.Add("smstemplate2fa", smstemplate2fa)
	// 	req.URL.RawQuery = q.Encode()
	// 	req.Header.Add("content-Type", "application/json")

	// 	err := RunRequest(req, data)
	// 	return *data, err
}
