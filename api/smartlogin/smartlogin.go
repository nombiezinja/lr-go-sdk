package smartlogin

import (
	lrvalidate "bitbucket.org/nombiezinja/lr-go-sdk/internal/validate"

	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
)

// // SmartLoginBool contains data from responses that return a single boolean attribute
// type SmartLoginBool struct {
// 	IsPosted   bool `json:"IsPosted"`
// 	IsVerified bool `json:"IsVerified"`
// }

// // SmartLogin contains the login information received by Smart Login Ping
// type SmartLogin struct {
// 	Profile     AuthProfile `json:"Profile"`
// 	AccessToken string      `json:"access_token"`
// 	ExpiresIn   time.Time   `json:"expires_in"`
// }

// GetSmartLoginByEmail sends a Smart Login link to the user's Email Id.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/smart-login/smart-login-by-email
// Required query parameters: apikey, email, clientguid
// Optional query parameters: smartloginemailtemplate, welcomeemailtemplate, redirecturl
func (lr Loginradius) GetSmartLoginByEmail(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"email": true, "clientguid": true, "smartloginemailtemplate": true, "welcomeemailtemplate": true, "redirecturl": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apikey"] = lr.Client.Context.ApiKey
	req := lr.Client.NewGetReq("/identity/v2/auth/login/smartlogin", validatedQueries)
	delete(req.QueryParams, "apiKey")
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetSmartLoginByUsername sends a Smart Login link to the user's Email Id.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/smart-login/smart-login-by-username
// Required query parameters: apikey, username, clientguid
// Optional query parameters: smartloginemailtemplate, welcomeemailtemplate, redirecturl
func (lr Loginradius) GetSmartLoginByUsername(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"username": true, "clientguid": true, "smartloginemailtemplate": true, "welcomeemailtemplate": true, "redirecturl": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apikey"] = lr.Client.Context.ApiKey
	req := lr.Client.NewGetReq("/identity/v2/auth/login/smartlogin", validatedQueries)
	delete(req.QueryParams, "apiKey")
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetSmartLoginPing is used to check if the Smart Login link has been clicked or not.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/smart-login/smart-login-ping
// Required query parameters: apikey, clientguid
func (lr Loginradius) GetSmartLoginPing(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"clientguid": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apikey"] = lr.Client.Context.ApiKey
	req := lr.Client.NewGetReq("/identity/v2/auth/login/smartlogin/ping", validatedQueries)
	delete(req.QueryParams, "apiKey")
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetSmartLoginVerifyToken verifies the provided token for Smart Login.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/smart-login/smart-login-verify-token
// Required query parameterS: apikey, verificationtoken,
// Optional query parameters: welcommeemailtemplate
func (lr Loginradius) GetSmartLoginVerifyToken(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"clientguid": true, "verificationtoken": true, "welcomeemailtemplate": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apikey"] = lr.Client.Context.ApiKey
	req := lr.Client.NewGetReq("/identity/v2/auth/email/smartlogin", validatedQueries)
	delete(req.QueryParams, "apiKey")
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
	// 	data := new(SmartLoginBool)
	// 	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/auth/email/smartlogin", "")
	// 	if reqErr != nil {
	// 		return *data, reqErr
	// 	}

	// 	q := req.URL.Query()
	// 	q.Add("apikey", os.Getenv("APIKEY"))
	// 	q.Add("verificationtoken", verificationToken)
	// 	q.Add("welcomeemailtemplate", welcomeEmailTemplate)
	// 	req.URL.RawQuery = q.Encode()
	// 	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	// 	err := RunRequest(req, data)
	// 	return *data, err
}
