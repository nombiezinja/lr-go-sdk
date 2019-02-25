package lrauthentication

import (
	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	"bitbucket.org/nombiezinja/lr-go-sdk/internal/sott"
	lrvalidate "bitbucket.org/nombiezinja/lr-go-sdk/internal/validate"
)

// PostAuthAddEmail is used to add additional emails to a user's account.
// Pass data in struct lrbody.AddEmail as body to help ensure parameters satisfy API requirements
// Required queries: apiKey; optional queries: verificationurl, emailtemplate
// Body params: email(string), type(string)
func (lr Loginradius) PostAuthAddEmail(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.NewAuthPostReqWithToken("/identity/v2/auth/email", body)
	if err != nil {
		return nil, err
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"verificationurl": true, "emailtemplate": true,
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

// PostAuthForgotPassword is used to send the reset password url to a specified account.
// Note: If you have the UserName workflow enabled, you may replace the 'email' parameter with 'username'
// Pass data in struct lrbody.EmailStr as body to help ensure parameters satisfy API requirements
// Required queries: apiKey, resetpasswordurl; optional queries: emailtemplate
// Required post parameter:email: string
func (lr Loginradius) PostAuthForgotPassword(body interface{}, queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"resetpasswordurl": true, "emailtemplate": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)

	if err != nil {
		return nil, err
	}

	request, err := lr.NewAuthPostReq("/identity/v2/auth/password", body, validatedQueries)
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PostAuthUserRegistrationByEmail creates a user in the database as well as sends a verification email to the user.
// Post parameters are an array of email objects (Check docs for more info) and password: string
// Pass data in struct lrbody.RegistrationUser as body to help ensure parameters satisfy API requirements
func (lr Loginradius) PostAuthUserRegistrationByEmail(queries interface{}, body interface{}) (*httprutils.Response, error) {
	sott := sott.Generate(lr.Context.ApiKey, lr.Context.ApiSecret)
	allowedQueries := map[string]bool{
		"verificationurl": true, "emailtemplate": true, "options": true,
	}
	validatedParams, err := lrvalidate.Validate(allowedQueries, queries)

	if err != nil {
		return nil, err
	}

	validatedParams["apiKey"] = lr.Context.ApiKey
	request, err := lr.NewAuthPostReq("/identity/v2/auth/register", body, validatedParams)

	request.Headers["X-LoginRadius-Sott"] = sott
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PostAuthLoginByEmail retrieves a copy of the user data based on the Email after verifying
// the validity of submitted credentials
// Pass data in struct lrbody.EmailLogin as body to help ensure parameters satisfy API requirements
// Required queries: apiKey; optional queries: verificationurl, loginurl, emailtemplate, g-recaptcha-response
// Required body param: email, password; optional body param: security answer
func (lr Loginradius) PostAuthLoginByEmail(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.NewAuthPostReq("/identity/v2/auth/login", body)
	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"verificationurl": true, "loginurl": true, "emailtemplate": true, "g-recaptcha-response": true,
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

// PostAuthLoginByUsername retrieves a copy of the user data based on the Username after verifying
// the validity of submitted credentials
// Post parameters are username: string, password: string and optional securityanswer: string
// Pass data in struct lrbody.UsernameLogin as body to help ensure parameters satisfy API requirements
func (lr Loginradius) PostAuthLoginByUsername(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.NewAuthPostReq("/identity/v2/auth/login", body)
	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"verificationurl": true, "loginurl": true, "emailtemplate": true, "g-recaptcha-response": true,
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
