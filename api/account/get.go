package lraccount

import (
	"os"

	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
)

// GetManageAccountProfilesByEmail is used to retrieve all of the profile data,
// associated with the specified account by email in Cloud Storage.
func GetManageAccountProfilesByEmail(email string) (*httprutils.Response, error) {
	request := httprutils.Request{
		Method: httprutils.Get,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/manage/account",
		Headers: map[string]string{
			"content-Type": "application/x-www-form-urlencoded",
		},
		QueryParams: map[string]string{
			"apikey":    os.Getenv("APIKEY"),
			"apisecret": os.Getenv("APISECRET"),
			"email":     email,
		},
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// GetManageAccountIdentitiesByEmail is used to retrieve all of the identities (UID and Profiles),
// associated with a specified email in Cloud Storage.
// Note: This is intended for specific workflows where an email may be associated to multiple UIDs.
func GetManageAccountIdentitiesByEmail(email string) (*httprutils.Response, error) {
	request := httprutils.Request{
		Method: httprutils.Get,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/manage/account/identities",
		Headers: map[string]string{
			"content-Type":            "application/x-www-form-urlencoded",
			"X-LoginRadius-ApiKey":    os.Getenv("APIKEY"),
			"X-LoginRadius-ApiSecret": os.Getenv("APISECRET"),
			"email":                   email,
		},
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}
