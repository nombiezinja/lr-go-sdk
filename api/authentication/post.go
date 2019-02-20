package lrauthentication

import (
	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	"bitbucket.org/nombiezinja/lr-go-sdk/internal/sott"
	lrvalidate "bitbucket.org/nombiezinja/lr-go-sdk/internal/validate"
)

// import (
// 	"fmt"
// 	"os"

// 	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
// 	"bitbucket.org/nombiezinja/lr-go-sdk/internal/sott"
// )

// // PostAuthAddEmail is used to add additional emails to a user's account.
// // Pass data in struct lrbody.AddEmail as body to help ensure parameters satisfy API requirements
// func PostAuthAddEmail(verificationURL, emailTemplate, token string, body interface{}) (*httprutils.Response, error) {
// 	tokenHeader := "Bearer " + token

// 	requestBody, error := httprutils.EncodeBody(body)
// 	if error != nil {
// 		return nil, error
// 	}

// 	request := httprutils.Request{
// 		Method: httprutils.Post,
// 		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/email",
// 		Headers: map[string]string{
// 			"content-Type":  "application/json",
// 			"Authorization": tokenHeader,
// 		},
// 		QueryParams: map[string]string{
// 			"apiKey": os.Getenv("APIKEY"),
// 		},
// 		Body: requestBody,
// 	}

// 	response, err := httprutils.TimeoutClient.Send(request)
// 	return response, err
// }

// // PostAuthForgotPassword is used to send the reset password url to a specified account.
// // Note: If you have the UserName workflow enabled, you may replace the 'email' parameter with 'username'
// // Post parameter is email: string
// // Pass data in struct lrbody.EmailStr as body to help ensure parameters satisfy API requirements
// func PostAuthForgotPassword(resetPasswordURL, emailTemplate string, body interface{}) (*httprutils.Response, error) {
// 	requestBody, error := httprutils.EncodeBody(body)
// 	if error != nil {
// 		return nil, error
// 	}

// 	request := httprutils.Request{
// 		Method: httprutils.Post,
// 		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/password",
// 		Headers: map[string]string{
// 			"content-Type": "application/json",
// 		},
// 		QueryParams: map[string]string{
// 			"apiKey":           os.Getenv("APIKEY"),
// 			"resetpasswordurl": resetPasswordURL,
// 			"emailTemplate":    emailTemplate,
// 		},
// 		Body: requestBody,
// 	}

// 	response, err := httprutils.TimeoutClient.Send(request)
// 	fmt.Println(response.Body)
// 	fmt.Println(err)

// 	return response, err
// }

// PostAuthUserRegistrationByEmail creates a user in the database as well as sends a verification email to the user.
// Post parameters are an array of email objects (Check docs for more info) and password: string
// Pass data in struct lrbody.RegistrationUser as body to help ensure parameters satisfy API requirements
func (lr Loginradius) PostAuthUserRegistrationByEmail(queries interface{}, body interface{}) (*httprutils.Response, error) {

	requestBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	sott := sott.Generate(lr.Context.ApiKey, lr.Context.ApiSecret)
	allowedQueries := map[string]bool{
		"verificationurl": true, "emailtemplate": true, "options": true,
	}
	validatedParams, err := lrvalidate.Validate(allowedQueries, queries)

	if err != nil {
		return nil, error
	}

	validatedParams["apiKey"] = lr.Context.ApiKey

	request := httprutils.Request{
		Method: httprutils.Post,
		URL:    lr.Domain + "/identity/v2/auth/register",
		Headers: map[string]string{
			"X-LoginRadius-Sott": sott,
			"content-Type":       "application/json",
		},
		QueryParams: validatedParams,
		Body:        requestBody,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// PostAuthLoginByEmail retrieves a copy of the user data based on the Email after verifying
// the validity of submitted credentials
// Pass data in struct lrbody.EmailLogin as body to help ensure parameters satisfy API requirements
// func (lr Loginradius) PostAuthLoginByEmail(verificationURL, loginURL, emailTemplate,
// 	gRecaptchaResponse, options string, body interface{}) (*httprutils.Response, error) {
func (lr Loginradius) PostAuthLoginByEmail(queries interface{}, body interface{}) (*httprutils.Response, error) {
	requestBody, error := httprutils.EncodeBody(body)

	if error != nil {
		return nil, error
	}

	allowed := map[string]bool{
		"verificationURL": true, "loginURL": true, "emailTemplate": true, "gRecaptchaResponse": true, "options": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowed, queries)

	if error != nil {
		return nil, error
	}

	validatedQueries["apiKey"] = lr.Context.ApiKey

	request := httprutils.Request{
		Method: httprutils.Post,
		URL:    lr.Domain + "/identity/v2/auth/login",
		Headers: map[string]string{
			"content-Type": "application/json",
		},
		QueryParams: validatedQueries,
		Body:        requestBody,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// // PostAuthLoginByUsername retrieves a copy of the user data based on the Username after verifying
// // the validity of submitted credentials
// // Post parameters are username: string, password: string and optional securityanswer: string
// // Pass data in struct lrbody.UsernameLogin as body to help ensure parameters satisfy API requirements
// func PostAuthLoginByUsername(verificationURL, loginURL, emailTemplate,
// 	gRecaptchaResponse, options string, body interface{}) (*httprutils.Response, error) {

// 	requestBody, error := httprutils.EncodeBody(body)
// 	if error != nil {
// 		return nil, error
// 	}

// 	request := httprutils.Request{
// 		Method: httprutils.Post,
// 		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/login",
// 		Headers: map[string]string{
// 			"content-Type": "application/json",
// 		},
// 		QueryParams: map[string]string{
// 			"apiKey":               os.Getenv("APIKEY"),
// 			"loginurl":             loginURL,
// 			"verificationurl":      verificationURL,
// 			"g-recaptcha-response": gRecaptchaResponse,
// 			"emailtemplate":        emailTemplate,
// 			"options":              options,
// 		},
// 		Body: requestBody,
// 	}

// 	response, err := httprutils.TimeoutClient.Send(request)
// 	return response, err
// }
