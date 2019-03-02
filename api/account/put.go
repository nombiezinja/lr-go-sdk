package lraccount

import (
	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
)

// PutManageAccountUpdateSecurityQuestionConfig is used to update security questions configuration on an existing account.
// Required post parameter: security question answer - object.
// Required template parameter: uid
// For more information regarding security questions, refer to this document: https://docs.loginradius.com/api/v2/dashboard/platform-security/password-policy
// Pass data in struct lrbody.AccountSecurityQuestion as body to help ensure parameters satisfy API requirements
func (lr Loginradius) PutManageAccountUpdateSecurityQuestionConfig(uid string, body interface{}) (*httprutils.Response, error) {

	request, err := lr.Client.NewPutReq("/identity/v2/manage/account/"+uid, body)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(request)

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PutManageAccountSetPassword is used to set the password of an account in Cloud Storage.
// Required post parameter: password - string
// Required template parameter: uid
// Pass data in struct lrbody.AccountPassword as body to help ensure parameters satisfy API requirements
func (lr Loginradius) PutManageAccountSetPassword(uid string, body interface{}) (*httprutils.Response, error) {

	request, err := lr.Client.NewPutReq("/identity/v2/manage/account/"+uid+"/password", body)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(request)

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PutManageAccountUpdate is used to update the information of existing accounts in your Cloud Storage.
// See our Advanced API Usage section for more capabilities.
// Post parameters: profile data that needs to be updated.
// Pass data in struct lrbody.UpdateProfile as body to help ensure parameters satisfy API requirements
// modify struct fields based on need
func (lr Loginradius) PutManageAccountUpdate(uid string, body interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReq("/identity/v2/manage/account/"+uid, body)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(request)

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}
