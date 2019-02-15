package lraccount

import (
	"os"

	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
)

// PostManageAccountCreate is used to create an account in Cloud Storage.
// This API bypasses the normal email verification process and manually creates the user.
// In order to use this API, you need to format a JSON request body with all of the mandatory fields
// Required post parameters are email object and password:string. Rest are optional profile parameters.
// Pass data in struct lrbody.AccountCreate as body to help ensure parameters satisfy API requirements
func PostManageAccountCreate(body interface{}) (*httprutils.Response, error) {
	requestBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := httprutils.Request{
		Method: httprutils.Post,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/manage/account",
		Headers: map[string]string{
			"content-Type":            "application/json",
			"X-LoginRadius-ApiKey":    os.Getenv("APIKEY"),
			"X-LoginRadius-ApiSecret": os.Getenv("APISECRET"),
		},
		Body: requestBody,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// PostManageForgotPasswordToken returns a forgot password token. Note: If you have the
// UserName workflow enabled, you may replace the 'email' parameter with 'username'.
// Post parameter is either the username: string or the email: string
// Pass data in struct lrbody.Username or lrbody.Email as body to help ensure parameters satisfy API requirements
func PostManageForgotPasswordToken(body interface{}) (*httprutils.Response, error) {
	encoded, err := httprutils.EncodeBody(body)
	request := httprutils.Request{

		Method: httprutils.Post,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/manage/account/forgot/token",
		Headers: map[string]string{
			"content-Type":            "application/json",
			"X-LoginRadius-ApiKey":    os.Getenv("APIKEY"),
			"X-LoginRadius-ApiSecret": os.Getenv("APISECRET"),
		},
		Body: encoded,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// PostManageEmailVerificationToken Returns an Email Verification token.
// Post parameter is the email: string
// Pass data in struct lrbody.EmailForVToken as body to help ensure parameters satisfy API requirements
func PostManageEmailVerificationToken(body interface{}) (*httprutils.Response, error) {
	encoded, err := httprutils.EncodeBody(body)
	request := httprutils.Request{
		Method: httprutils.Post,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/manage/account/verify/token",
		Headers: map[string]string{
			"content-Type":            "application/json",
			"X-LoginRadius-ApiKey":    os.Getenv("APIKEY"),
			"X-LoginRadius-ApiSecret": os.Getenv("APISECRET"),
		},
		Body: encoded,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}
