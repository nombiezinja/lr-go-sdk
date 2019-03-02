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
