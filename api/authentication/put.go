package lrauthentication

import (
	"os"

	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	lrvalidate "bitbucket.org/nombiezinja/lr-go-sdk/internal/validate"
)

// PutAuthVerifyEmailByOtp will send the welcome email.
// Post parameters include otp: string, email: string, optional securityanswer: string, optional qq_captcha_ticket: string,
// optional qq_captcha_randstr: string and optional g-recaptcha-response:string
func PutAuthVerifyEmailByOtp(url, welcomeEmailTemplate string, body interface{}) (*httprutils.Response, error) {
	requestBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := httprutils.Request{
		Method: httprutils.Get,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/email",
		Headers: map[string]string{
			"content-Type": "application/json",
		},
		QueryParams: map[string]string{
			"apikey":          os.Getenv("APIKEY"),
			"url":             url,
			"welcomeTemplate": welcomeEmailTemplate,
		},
		Body: requestBody,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// PutAuthChangePassword is used to change the accounts password based on the previous password.
// Post parameters- oldpassword: string;newpassword: string
// Required query paramter: apikey
// Pass data in struct lrbody.ChangePassword as body to help ensure parameters satisfy API requirements
func (lr Loginradius) PutAuthChangePassword(body interface{}) (*httprutils.Response, error) {

	request, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/password/change", body)

	if err != nil {
		return nil, err
	}

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PutAuthLinkSocialIdentities is used to link up a social provider account with the specified
// account based on the access token and the social providers user access token.
// Post parameters- candidatetoken: string
// Pass data in struct lrbody.LinkSocialIds as body to help ensure parameters satisfy API requirements
func PutAuthLinkSocialIdentities(token string, body interface{}) (*httprutils.Response, error) {
	tokenHeader := "Bearer " + token

	requestBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := httprutils.Request{
		Method: httprutils.Put,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/socialidentity",
		Headers: map[string]string{
			"content-Type":  "application/json",
			"Authorization": tokenHeader,
		},
		QueryParams: map[string]string{
			"apikey": os.Getenv("APIKEY"),
		},
		Body: requestBody,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// PutResendEmailVerification resends the verification email to the user.
// Post parameter- email: string
// Required query parameter: apikey ; optional query parameters: emailtemplate, verificationurl
// Pass data in struct lrbody.EmailStr as body to help ensure parameters satisfy API requirements
func (lr Loginradius) PutResendEmailVerification(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReq("/identity/v2/auth/register", body)
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

// PutAuthResetPasswordByResetToken is used to set a new password for the specified account.
// Required post parameters- resettoken: string; password: string
// optional post parameters- welcomeemailtemplate: string; resetpasswordemailtemplate: string
// Required query parameter: apiKey
// Pass data in struct lrbody.ResetPw as body to help ensure parameters satisfy API requirements
func (lr Loginradius) PutAuthResetPasswordByResetToken(body interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReq("/identity/v2/auth/password/reset", body)
	if err != nil {
		return nil, err
	}

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PutAuthResetPasswordByOTP is used to set a new password for the specified account.
// Post parameters are the password: string, otp: string, email: string,
// optional welcomeemailtemplate: string and optional resetpasswordemailtemplate: string
// Pass data in struct lrbody.ResetPwOtp as body to help ensure parameters satisfy API requirements
func PutAuthResetPasswordByOTP(body interface{}) (*httprutils.Response, error) {
	requestBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := httprutils.Request{
		Method: httprutils.Put,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/password/reset",
		Headers: map[string]string{
			"content-Type": "application/json",
		},
		QueryParams: map[string]string{
			"apikey": os.Getenv("APIKEY"),
		},
		Body: requestBody,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// PutAuthResetPasswordBySecurityAnswerAndEmail is used to reset password for the specified account by security question.
// Post parameters are the password: string, email: string, securityanswer: string
// and optional resetpasswordemailtemplate: string
// Pass data in struct lrbody.ResetPwSecurityQuestionEmail as body to help ensure parameters satisfy API requirements
func PutAuthResetPasswordBySecurityAnswerAndEmail(body interface{}) (*httprutils.Response, error) {
	requestBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := httprutils.Request{
		Method: httprutils.Put,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/password/securityanswer",
		Headers: map[string]string{
			"content-Type": "application/json",
		},
		QueryParams: map[string]string{
			"apikey": os.Getenv("APIKEY"),
		},
		Body: requestBody,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// PutAuthResetPasswordBySecurityAnswerAndPhone is used to reset password for the specified account by security question.
// Post parameters are the password: string, phone: string, securityanswer: string
// and optional resetpasswordemailtemplate: string
// Pass data in struct lrbody.ResetPwSecurityQuestionPhone as body to help ensure parameters satisfy API requirements
func PutAuthResetPasswordBySecurityAnswerAndPhone(body interface{}) (*httprutils.Response, error) {
	requestBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := httprutils.Request{
		Method: httprutils.Put,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/password/securityanswer",
		Headers: map[string]string{
			"content-Type": "application/json",
		},
		QueryParams: map[string]string{
			"apikey": os.Getenv("APIKEY"),
		},
		Body: requestBody,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// PutAuthResetPasswordBySecurityAnswerAndUsername is used to reset password for the specified account by security question.
// Post parameters are the password: string, username: string, securityanswer: string
// and optional resetpasswordemailtemplate: string
// Pass data in struct lrbody.ResetPwSecurityQuestionusername as body to help ensure parameters satisfy API requirements
func PutAuthResetPasswordBySecurityAnswerAndUsername(body interface{}) (*httprutils.Response, error) {
	requestBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := httprutils.Request{
		Method: httprutils.Put,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/password/securityanswer",
		Headers: map[string]string{
			"content-Type": "application/json",
		},
		QueryParams: map[string]string{
			"apikey": os.Getenv("APIKEY"),
		},
		Body: requestBody,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// PutAuthSetOrChangeUsername is used to set or change UserName by access token.
// Post parameter is username: string
// Pass data in struct lrbody.AuthUsername as body to help ensure parameters satisfy API requirements
func PutAuthSetOrChangeUsername(token string, body interface{}) (*httprutils.Response, error) {
	tokenHeader := "Bearer " + token

	requestBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := httprutils.Request{
		Method: httprutils.Put,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/username",
		Headers: map[string]string{
			"content-Type":  "application/json",
			"Authorization": tokenHeader,
		},
		QueryParams: map[string]string{
			"apikey": os.Getenv("APIKEY"),
		},
		Body: requestBody,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// PutAuthUpdateProfileByToken is used to update the user's profile by passing the access_token.
// Post parameters are fields in the profile that need to be updated
// Pass data in struct lrbody.UpdateProfile as body to help ensure parameters satisfy API requirements
// modify struct fields based on need
// Required query parameter: apiKey; optional query parameters: smstemplate, emailtemplate, verificationurl
func (lr Loginradius) PutAuthUpdateProfileByToken(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	// tokenHeader := "Bearer " + token

	// requestBody, error := httprutils.EncodeBody(body)
	// if error != nil {
	// 	return nil, error
	// }

	// request := httprutils.Request{
	// 	Method: httprutils.Put,
	// 	URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/account",
	// 	Headers: map[string]string{
	// 		"content-Type":  "application/json",
	// 		"Authorization": tokenHeader,
	// 	},
	// 	QueryParams: map[string]string{
	// 		"apikey":          os.Getenv("APIKEY"),
	// 		"verificationurl": verificationURL,
	// 		"emailtemplate":   emailTemplate,
	// 		"smstemplate":     smsTemplate,
	// 	},
	// 	Body: requestBody,
	// }

	request, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/account", body)
	if err != nil {
		return nil, err
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"verificationurl": true, "emailtemplate": true, "smstemplate": true,
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

// PutAuthUpdateSecurityQuestionByAccessToken is used to update security questions by the access token.
// Body parameter is the securityquestionanswer: string
// Pass data in struct lrbody.SecurityQuestionAnswer as body to help ensure parameters satisfy API requirements
// For more information on this parameter, please see: https://www.loginradius.com/docs/api/v2/dashboard/platform-security/password-policy#securityquestion4
func PutAuthUpdateSecurityQuestionByAccessToken(token string, body interface{}) (*httprutils.Response, error) {

	requestBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	tokenHeader := "Bearer " + token

	request := httprutils.Request{
		Method: httprutils.Put,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/account",
		Headers: map[string]string{
			"content-Type":  "application/json",
			"Authorization": tokenHeader,
		},
		QueryParams: map[string]string{
			"apikey": os.Getenv("APIKEY"),
		},
		Body: requestBody,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}
