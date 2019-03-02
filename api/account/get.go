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
// Required template parameter: uid
// func GetManageAccountPassword(uid string) (AccountPassword, error) {
// 	data := new(AccountPassword)
// 	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/manage/account/"+uid+"/password", "")
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
// 	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

// 	err := RunRequest(req, data)
// 	return *data, err
// }
func (lr Loginradius) GetManageAccountPassword(uid string) (*httprutils.Response, error) {
	request := lr.Client.NewGetReq("/identity/v2/manage/account/" + uid + "/password")
	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}
