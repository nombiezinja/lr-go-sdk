package lraccount

import (
	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
)

// DeleteManageAccount is used to delete the Users account and allows them to re-register for a new account.
// Required template variable: uid
func (lr Loginradius) DeleteManageAccount(uid string) (*httprutils.Response, error) {
	request := lr.Client.NewDeleteReq("/identity/v2/manage/account/")
	lr.Client.AddApiCredentialsToReqHeader(request)
	request.URL = request.URL + uid

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// DeleteManageAccount is used to delete the Users account and allows them to re-register for a new account.
// Required template variable: uid
func (lr Loginradius) DeleteManageAccountEmail(uid string, body interface{}) (*httprutils.Response, error) {
	request := lr.Client.NewDeleteReq("/identity/v2/manage/account/"+uid+"/email", body)
	lr.Client.AddApiCredentialsToReqHeader(request)
	request.URL = request.URL + uid

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// func DeleteManageAccountEmail(uid string, body interface{}) (AuthProfile, error) {
// 	data := new(AuthProfile)
// 	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN")+"/identity/v2/manage/account/"+uid+"/email", body)
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	req.Header.Add("content-Type", "application/json")
// 	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
// 	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

// 	err := RunRequest(req, data)
// 	return *data, err
// }
