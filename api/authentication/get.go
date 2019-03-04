package lrauthentication

import (
	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	lrvalidate "bitbucket.org/nombiezinja/lr-go-sdk/internal/validate"
)

// GetAuthVerifyEmail is used to verify the email of user.
// Note: This API will only return the full profile if you have'Enable auto login after email verification' set in your
// LoginRadius Dashboard's Email Workflow settings under 'Verification Email'
// Required queries: apiKey, verificationtoken;  Optional queries: url
func (lr Loginradius) GetAuthVerifyEmail(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"url": true, "verificationtoken": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Client.Context.ApiKey

	request := lr.Client.NewGetReq("/identity/v2/auth/email", validatedQueries)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthCheckEmailAvailability is used to check whether an email exists or not on your site.
// Required queries: apiKey, email
func (lr Loginradius) GetAuthCheckEmailAvailability(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"email": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Client.Context.ApiKey

	request := lr.Client.NewGetReq("/identity/v2/auth/email", validatedQueries)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthCheckUsernameAvailability is used to check the UserName exists or not on your site.
// Required queries: apiKey, username
func (lr Loginradius) GetAuthCheckUsernameAvailability(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"username": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Client.Context.ApiKey

	request := lr.Client.NewGetReq("/identity/v2/auth/username", validatedQueries)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthReadProfilesByToken retrieves a copy of the user data based on the access token.
// Required queries: apiKey
func (lr Loginradius) GetAuthReadProfilesByToken() (*httprutils.Response, error) {
	request, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/account")
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthPrivatePolicyAccept is used update the privacy policy stored in the user's profile based on user's access token
func (lr Loginradius) GetAuthPrivatePolicyAccept() (*httprutils.Response, error) {
	request, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/privacypolicy/accept")
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthSendWelcomeEmail sends the welcome email.
// Queries are optional and can be passed as variadic argument
func (lr Loginradius) GetAuthSendWelcomeEmail(queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/account/sendwelcomeemail")
	if err != nil {
		return nil, err
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{"welcomeemailtemplate": true}
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

// GetAuthSocialIdentity is called just before account linking API and it prevents
// the raas profile of the second account from getting created.
// Required queries: apiKey
func (lr Loginradius) GetAuthSocialIdentity() (*httprutils.Response, error) {
	request, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/socialidentity")
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthValidateAccessToken returns an expiry date for the access token if it is valid
// and an error if it is invalid
// Required queries: apiKey
func (lr Loginradius) GetAuthValidateAccessToken() (*httprutils.Response, error) {
	request, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/access_token/validate")
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthDeleteAccount is used to delete an account by passing it a delete token.
// Required queries: apiKey, deletetoken
func (lr Loginradius) GetAuthDeleteAccount(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"deletetoken": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Client.Context.ApiKey
	request := lr.Client.NewGetReq("/identity/v2/auth/account/delete", validatedQueries)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthInvalidateAccessToken invalidates the active access_token or expires an access token's validity.
// Required queries: apiKey
func (lr Loginradius) GetAuthInvalidateAccessToken() (*httprutils.Response, error) {
	request, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/access_token/invalidate")
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthSecurityQuestionByAccessToken is used to retrieve the
// list of questions that are configured on the respective LoginRadius site for the user.
// Will return error unless security question is enabled
// Refer to this document: https://docs.loginradius.com/api/v2/dashboard/platform-security/password-policy
// Endpoint returns an array rather than JSON, use lrjson.UnmarshalGetAuthQuestion rather than lrjson.DynamicUnmarshal
// to unmarshal into struct.
// Required queries: apiKey
func (lr Loginradius) GetAuthSecurityQuestionByAccessToken() (*httprutils.Response, error) {
	request, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/securityquestion/accesstoken")
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthSecurityQuestionByEmail is used to retrieve the
// list of questions that are configured on the respective LoginRadius site for the user.
// Will return error unless security question feature is enabled
// Follow instructions in this document: https://docs.loginradius.com/api/v2/dashboard/platform-security/password-policy
// Endpoint returns an array rather than JSON, use lrjson.UnmarshalGetAuthQuestion rather than lrjson.DynamicUnmarshal
// to unmarshal into struct.
// Required queries: apiKey
func (lr Loginradius) GetAuthSecurityQuestionByEmail(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"email": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	request, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/securityquestion/email", validatedQueries)
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthSecurityQuestionByUsername is used to retrieve the
// list of questions that are configured on the respective LoginRadius site for the user.
// Will return error unless security question feature is enabled.
// Follow instructions in this document: https://docs.loginradius.com/api/v2/dashboard/platform-security/password-policy
// Endpoint returns an array rather than JSON, use lrjson.UnmarshalGetAuthQuestion rather than lrjson.DynamicUnmarshal
// to unmarshal into struct.
// Required queries: apikey
func (lr Loginradius) GetAuthSecurityQuestionByUsername(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"username": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	request, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/securityquestion/username", validatedQueries)
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetAuthSecurityQuestionByPhone is used to retrieve the
// list of questions that are configured on the respective LoginRadius site for the user.
// Will return error unless security question feature is enabled
// Follow instructions in this document: https://docs.loginradius.com/api/v2/dashboard/platform-security/password-policy
// Endpoint returns an array rather than JSON, use lrjson.UnmarshalGetAuthQuestion rather than lrjson.DynamicUnmarshal
// to unmarshal into struct.
// Required queries: phone
func (lr Loginradius) GetAuthSecurityQuestionByPhone(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"phone": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	request, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/securityquestion/phone", validatedQueries)
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetPasswordlessLoginByEmail is used to send a Passwordless Login verification link to the provided Email ID.
// Required queries: email, apiKey; optional queries: passwordlesslogintemplate, verificationurl
func (lr Loginradius) GetPasswordlessLoginByEmail(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"email": true, "passwordlesslogintemplate": true, "verificationurl": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Client.Context.ApiKey

	request := lr.Client.NewGetReq("/identity/v2/auth/login/passwordlesslogin/email", validatedQueries)

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetPasswordlessLoginByUsername is used to send a Passwordless Login verification link to the provided Username.
// Required queries: username, apiKey; optional queries: passwordlesslogintemplate, verificationurl
func (lr Loginradius) GetPasswordlessLoginByUsername(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"username": true, "passwordlesslogintemplate": true, "verificationurl": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Client.Context.ApiKey
	request := lr.Client.NewGetReq("/identity/v2/auth/login/passwordlesslogin/email", validatedQueries)

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetPasswordlessLoginVerification is used to verify the Passwordless Login verification link.
// Required queries: verificationtoken; optional queries: welcomeemailtemplate
func (lr Loginradius) GetPasswordlessLoginVerification(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"verificationtoken": true, "welcomeemailtemplate": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Client.Context.ApiKey
	request := lr.Client.NewGetReq("/identity/v2/auth/login/passwordlesslogin/email/verify", validatedQueries)

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}
