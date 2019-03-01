package lraccount

import (
	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	lrvalidate "bitbucket.org/nombiezinja/lr-go-sdk/internal/validate"
)

// PostManageAccountCreate is used to create an account in Cloud Storage.
// This API bypasses the normal email verification process and manually creates the user.
// In order to use this API, you need to format a JSON request body with all of the mandatory fields
// Required post parameters: email - object;  assword - string. Rest are optional profile parameters.
// Required query parameters: apiKey, apiSecret
// Pass data in struct lrbody.AccountCreate as body to help ensure parameters satisfy API requirements
func (lr Loginradius) PostManageAccountCreate(body interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPostReq("/identity/v2/manage/account", body)
	if err != nil {
		return nil, err
	}

	delete(request.QueryParams, "apiKey")
	lr.Client.AddApiCredentialsToReqHeader(request)

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PostManageForgotPasswordToken returns a forgot password token. Note: If you have the
// UserName workflow enabled, you may replace the 'email' parameter with 'username'.
// Required post parameters: email - string OR username - string
// Pass data in struct lrbody.Username or lrbody.Email as body to help ensure parameters satisfy API requirements
// Optional query parameters: sendemail - string; emailTemplate-string; resetPasswordUrl-string
func (lr Loginradius) PostManageForgotPasswordToken(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPostReq("/identity/v2/manage/account/forgot/token", body)
	if err != nil {
		return nil, err
	}

	delete(request.QueryParams, "apiKey")
	lr.Client.AddApiCredentialsToReqHeader(request)

	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"sendemail": true, "emailTemplate": true, "resetPasswordUrl": true,
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
}

// PostManageEmailVerificationToken Returns an Email Verification token.
// Post parameter is the email: string
// Pass data in struct lrbody.EmailForVToken as body to help ensure parameters satisfy API requirements
func (lr Loginradius) PostManageEmailVerificationToken(body interface{}) (*httprutils.Response, error) {
	encoded, err := httprutils.EncodeBody(body)
	request := httprutils.Request{
		Method: httprutils.Post,
		URL:    lr.Client.Domain + "/identity/v2/manage/account/verify/token",
		Headers: map[string]string{
			"content-Type":            "application/json",
			"X-LoginRadius-ApiKey":    lr.Client.Context.ApiKey,
			"X-LoginRadius-ApiSecret": lr.Client.Context.ApiSecret,
		},
		Body: encoded,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}
