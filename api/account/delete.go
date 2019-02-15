package lraccount

import (
	"os"

	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
)

// DeleteManageAccount is used to delete the Users account and allows them to re-register for a new account.
func DeleteManageAccount(uid string) (*httprutils.Response, error) {
	request := httprutils.Request{
		Method: httprutils.Delete,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/manage/account/" + uid,
		Headers: map[string]string{
			"content-Type":            "application/x-www-form-urlencoded",
			"X-LoginRadius-ApiKey":    os.Getenv("APIKEY"),
			"X-LoginRadius-ApiSecret": os.Getenv("APISECRET"),
		},
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}
