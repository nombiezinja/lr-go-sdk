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

	request, err := lr.Client.NewPutReq("/identity/v2/manage/account/", body)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(request)
	request.URL = request.URL + uid

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}
