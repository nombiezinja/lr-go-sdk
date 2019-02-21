package lrauthentication

import (
	"fmt"
	"os"

	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	lrvalidate "bitbucket.org/nombiezinja/lr-go-sdk/internal/validate"
)

// GetAuthVerifyEmail is used to verify the email of user.
// Note: This API will only return the full profile if you have
// 'Enable auto login after email verification' set in your
// LoginRadius Dashboard's Email Workflow settings under 'Verification Email'.
func (lr Loginradius) GetAuthVerifyEmail(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"url": true, "verificationtoken": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Context.ApiKey

	request := lr.NewAuthGetReq("/identity/v2/auth/email", validatedQueries)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthCheckEmailAvailability is used to check whether an email exists or not on your site.
func (lr Loginradius) GetAuthCheckEmailAvailability(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"email": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Context.ApiKey

	request := lr.NewAuthGetReq("/identity/v2/auth/email", validatedQueries)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthCheckUsernameAvailability is used to check the UserName exists or not on your site.
func (lr Loginradius) GetAuthCheckUsernameAvailability(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"username": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Context.ApiKey

	request := lr.NewAuthGetReq("/identity/v2/auth/username", validatedQueries)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthReadProfilesByToken retrieves a copy of the user data based on the access_token.
func (lr Loginradius) GetAuthReadProfilesByToken() (*httprutils.Response, error) {
	request, err := lr.NewAuthGetReqWithAccessToken("/identity/v2/auth/account")
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthPrivatePolicyAccept is used update the privacy policy stored in the user's profile based on user's access token
func (lr Loginradius) GetAuthPrivatePolicyAccept() (*httprutils.Response, error) {
	request, err := lr.NewAuthGetReqWithAccessToken("/identity/v2/auth/privacypolicy/accept")
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthSendWelcomeEmail sends the welcome email.
func (lr Loginradius) GetAuthSendWelcomeEmail(queries ...interface{}) (*httprutils.Response, error) {
	// tokenHeader := "Bearer " + token
	// request := httprutils.Request{
	// 	Method: httprutils.Get,
	// 	URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/account/sendwelcomeemail",
	// 	Headers: map[string]string{
	// 		"content-Type":  "application/x-www-form-urlencoded",
	// 		"Authorization": tokenHeader,
	// 	},
	// 	QueryParams: map[string]string{
	// 		"apiKey":               os.Getenv("APIKEY"),
	// 		"welcomeemailtemplate": welcomeEmailTemplate,
	// 	},
	// }
	request, err := lr.NewAuthGetReqWithAccessToken("/identity/v2/auth/account/sendwelcomeemail")
	if err != nil {
		return nil, err
	}

	allowedQueries := map[string]bool{"welcomeemailtemplate": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	for k, v := range validatedQueries {
		request.QueryParams[k] = v
	}
	fmt.Printf("%+v", request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthSocialIdentity is called just before account linking API and it prevents
// the raas profile of the second account from getting created.
func GetAuthSocialIdentity(token string) (*httprutils.Response, error) {
	tokenHeader := "Bearer " + token
	request := httprutils.Request{
		Method: httprutils.Get,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/socialidentity",
		Headers: map[string]string{
			"content-Type":  "application/x-www-form-urlencoded",
			"Authorization": tokenHeader,
		},
		QueryParams: map[string]string{
			"apiKey": os.Getenv("APIKEY"),
		},
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// GetAuthValidateAccessToken returns an expiry date for the access token if it is valid
// and an error if it is invalid
func GetAuthValidateAccessToken(token string) (*httprutils.Response, error) {
	tokenHeader := "Bearer " + token
	request := httprutils.Request{
		Method: httprutils.Get,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/access_token/validate",
		Headers: map[string]string{
			"content-Type":  "application/x-www-form-urlencoded",
			"Authorization": tokenHeader,
		},
		QueryParams: map[string]string{
			"apiKey": os.Getenv("APIKEY"),
		},
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// GetAuthDeleteAccount is used to delete an account by passing it a delete token.
func GetAuthDeleteAccount(deleteToken string) (*httprutils.Response, error) {
	request := httprutils.Request{
		Method: httprutils.Get,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/access_token/validate",
		Headers: map[string]string{
			"content-Type": "application/x-www-form-urlencoded",
		},
		QueryParams: map[string]string{
			"apiKey":      os.Getenv("APIKEY"),
			"deletetoken": deleteToken,
		},
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// GetAuthInvalidateAccessToken invalidates the active access_token or expires an access token's validity.
func GetAuthInvalidateAccessToken(token string) (*httprutils.Response, error) {
	tokenHeader := "Bearer " + token
	request := httprutils.Request{
		Method: httprutils.Get,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/access_token/invalidate",
		Headers: map[string]string{
			"content-Type":  "application/x-www-form-urlencoded",
			"Authorization": tokenHeader,
		},
		QueryParams: map[string]string{
			"apikey": os.Getenv("APIKEY"),
		},
	}
	fmt.Println("request", request)

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// GetAuthSecurityQuestionByAccessToken is used to retrieve the
// list of questions that are configured on the respective LoginRadius site for the user.
// Will return error unless security question is enabled
// Refer to this document: https://docs.loginradius.com/api/v2/dashboard/platform-security/password-policy
// This endpoint returns an array
func GetAuthSecurityQuestionByAccessToken(token string) (*httprutils.Response, error) {
	tokenHeader := "Bearer " + token
	request := httprutils.Request{
		Method: httprutils.Get,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/securityquestion/accesstoken",
		Headers: map[string]string{
			"content-Type":  "application/x-www-form-urlencoded",
			"Authorization": tokenHeader,
		},
		QueryParams: map[string]string{
			"apikey": os.Getenv("APIKEY"),
		},
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// GetAuthSecurityQuestionByEmail is used to retrieve the
// list of questions that are configured on the respective LoginRadius site for the user.
// Will return error unless security question feature is enabled
// Follow instructions in this document: https://docs.loginradius.com/api/v2/dashboard/platform-security/password-policy
// Endpoint returns an array rather than JSON, use lrjson.UnmarshalGetAuthQuestion rather than lrjson.DynamicUnmarshal
// to unmarshal into struct.
func GetAuthSecurityQuestionByEmail(email string) (*httprutils.Response, error) {
	request := httprutils.Request{
		Method: httprutils.Get,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/securityquestion/email",
		Headers: map[string]string{
			"content-Type": "application/x-www-form-urlencoded",
		},
		QueryParams: map[string]string{
			"apikey": os.Getenv("APIKEY"),
			"email":  email,
		},
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// GetAuthSecurityQuestionByUsername is used to retrieve the
// list of questions that are configured on the respective LoginRadius site for the user.
// Will return error unless security question feature is enabled.
// Follow instructions in this document: https://docs.loginradius.com/api/v2/dashboard/platform-security/password-policy
// Endpoint returns an array rather than JSON, use lrjson.UnmarshalGetAuthQuestion rather than lrjson.DynamicUnmarshal
// to unmarshal into struct.
func GetAuthSecurityQuestionByUsername(username string) (*httprutils.Response, error) {
	request := httprutils.Request{
		Method: httprutils.Get,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/securityquestion/username",
		Headers: map[string]string{
			"content-Type": "application/x-www-form-urlencoded",
		},
		QueryParams: map[string]string{
			"apikey":   os.Getenv("APIKEY"),
			"username": username,
		},
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// GetAuthSecurityQuestionByPhone is used to retrieve the
// list of questions that are configured on the respective LoginRadius site for the user.
// Will return error unless security question feature is enabled
// Follow instructions in this document: https://docs.loginradius.com/api/v2/dashboard/platform-security/password-policy
// Endpoint returns an array rather than JSON, use lrjson.UnmarshalGetAuthQuestion rather than lrjson.DynamicUnmarshal
// to unmarshal into struct.
func GetAuthSecurityQuestionByPhone(phone string) (*httprutils.Response, error) {
	request := httprutils.Request{
		Method: httprutils.Get,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/securityquestion/phone",
		Headers: map[string]string{
			"content-Type": "application/x-www-form-urlencoded",
		},
		QueryParams: map[string]string{
			"apikey": os.Getenv("APIKEY"),
			"phone":  phone,
		},
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// GetPasswordlessLoginByEmail is used to send a Passwordless Login verification link to the provided Email ID.
func GetPasswordlessLoginByEmail(email, passwordlessLoginTemplate, verificationURL string) (*httprutils.Response, error) {
	request := httprutils.Request{
		Method: httprutils.Get,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/login/passwordlesslogin/email",
		Headers: map[string]string{
			"content-Type": "x-www-form-urlencoded",
		},
		QueryParams: map[string]string{
			"apikey":                    os.Getenv("APIKEY"),
			"passwordlesslogintemplate": passwordlessLoginTemplate,
			"verificationurl":           verificationURL,
			"email":                     email,
		},
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// // GetPasswordlessLoginByUsername is used to send a Passwordless Login verification link to the provided Username.
func GetPasswordlessLoginByUsername(username, passwordlessLoginTemplate, verificationURL string) (*httprutils.Response, error) {
	request := httprutils.Request{
		Method: httprutils.Get,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/login/passwordlesslogin/email",
		Headers: map[string]string{
			"content-Type": "x-www-form-urlencoded",
		},
		QueryParams: map[string]string{
			"apikey":                    os.Getenv("APIKEY"),
			"passwordlesslogintemplate": passwordlessLoginTemplate,
			"verificationurl":           verificationURL,
			"username":                  username,
		},
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// GetPasswordlessLoginVerification is used to verify the Passwordless Login verification link.
func GetPasswordlessLoginVerification(verificationToken, welcomeEmailTemplate string) (*httprutils.Response, error) {
	request := httprutils.Request{
		Method: httprutils.Get,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/login/passwordlesslogin/email",
		Headers: map[string]string{
			"content-Type": "x-www-form-urlencoded",
		},
		QueryParams: map[string]string{
			"apikey":               os.Getenv("APIKEY"),
			"welcomeemailtemplate": welcomeEmailTemplate,
			"verificationtoken":    verificationToken,
		},
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}
