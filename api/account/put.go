package lraccount

import (
	"os"

	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
)

// PutManageAccountUpdateSecurityQuestionConfig is used to update security questions configuration on an existing account.
// Post parameter is the security question answer object.
// Pass data in struct lrbody.AccountSecurityQuestion as body to help ensure parameters satisfy API requirements
func PutManageAccountUpdateSecurityQuestionConfig(uid string, body interface{}) (*httprutils.Response, error) {
	encoded, err := httprutils.EncodeBody(body)
	request := httprutils.Request{

		Method: httprutils.Put,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/manage/account/" + uid,
		Headers: map[string]string{
			"content-Type":            "application/json",
			"X-LoginRadius-ApiKey":    os.Getenv("APIKEY"),
			"X-LoginRadius-ApiSecret": os.Getenv("APISECRET"),
		},
		Body: encoded,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}
