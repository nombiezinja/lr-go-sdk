package lraccount

import (
	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	lrvalidate "bitbucket.org/nombiezinja/lr-go-sdk/internal/validate"
)

// GetManageAccountProfilesByEmail is used to retrieve all of the profile data,
// associated with the specified account by email in Cloud Storage.
// This end point returns a single profile
// Required query param: email - string
func (lr Loginradius) GetManageAccountProfilesByEmail(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"email": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	request := lr.Client.NewGetReq("/identity/v2/manage/account", validatedQueries)
	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetManageAccountProfilesByUsername is used to retrieve all of the profile data,
// associated with the specified account by username in Cloud Storage.
// This end point returns a single profile
// Required query param: username - string
func (lr Loginradius) GetManageAccountProfilesByUsername(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"username": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	request := lr.Client.NewGetReq("/identity/v2/manage/account", validatedQueries)
	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetManageAccountProfilesByPhoneID is used to retrieve all of the profile data,
// associated with the specified account by PhoneID in Cloud Storage.
// This end point returns a single profile
// Required query param: phone - string
func (lr Loginradius) GetManageAccountProfilesByPhoneID(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"phone": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	request := lr.Client.NewGetReq("/identity/v2/manage/account", validatedQueries)
	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetManageAccountProfilesByUid is used to retrieve all of the profile data,
// associated with the specified account by uid in Cloud Storage.
// This end point returns a single profile
// Required template param: uid - string
func (lr Loginradius) GetManageAccountProfilesByUid(uid string) (*httprutils.Response, error) {
	request := lr.Client.NewGetReq("/identity/v2/manage/account/" + uid)
	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetManageAccountIdentitiesByEmail is used to retrieve all of the identities (UID and Profiles),
// associated with a specified email in Cloud Storage.
// Note: This is intended for specific workflows where an email may be associated to multiple UIDs.
// This end point returns data in an array, the response needs to be handled like so:
// 						body, _ := lrjson.DynamicUnmarshal(response.Body) // unmarshals body
// 						profiles := body["Data"].([]interface{}) // type assertion
// 						profile := profiles[0].(map[string]interface{}) // get first profile
// 						uid := profile["Uid"].(string) // get id of first profile
// Required query param: email - string
func (lr Loginradius) GetManageAccountIdentitiesByEmail(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"email": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	request := lr.Client.NewGetReq("/identity/v2/manage/account/identities", validatedQueries)
	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetManageAccessTokenUID is used to get LoginRadius access token based on UID.
// Required query params: uid
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-impersonation-api
func (lr Loginradius) GetManageAccessTokenUID(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"uid": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	request := lr.Client.NewGetReq("/identity/v2/manage/account/access_token", validatedQueries)
	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetManageAccountPassword is used to retrieve the hashed password of a specified account in Cloud Storage.
// Required template parameter: string representing uid
func (lr Loginradius) GetManageAccountPassword(uid string) (*httprutils.Response, error) {
	request := lr.Client.NewGetReq("/identity/v2/manage/account/" + uid + "/password")
	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}
