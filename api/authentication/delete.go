package lrauthentication

import (
	"os"

	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
)

// DeleteAuthDeleteAccountEmailConfirmation deletes a user account by passing the user's access token.
func DeleteAuthDeleteAccountEmailConfirmation(deleteURL, emailTemplate, token string) (*httprutils.Response, error) {
	tokenHeader := "Bearer " + token

	request := httprutils.Request{
		Method: httprutils.Delete,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/account",
		Headers: map[string]string{
			"content-Type":  "application/x-www-form-urlencoded",
			"Authorization": tokenHeader,
		},
		QueryParams: map[string]string{
			"apikey": os.Getenv("APIKEY"),
		},
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// DeleteAuthRemoveEmail is used to remove additional emails from a user's account.
// Post parameter is the e-mail that is to be removed.
// Pass data in struct lrbody.AuthUsername as body to help ensure parameters satisfy API requirements
func DeleteAuthRemoveEmail(token string, body interface{}) (*httprutils.Response, error) {
	requestBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	tokenHeader := "Bearer " + token

	request := httprutils.Request{
		Method: httprutils.Delete,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/email",
		Headers: map[string]string{
			"content-Type":  "application/json",
			"Authorization": tokenHeader,
		},
		QueryParams: map[string]string{
			"apikey": os.Getenv("APIKEY"),
		},
		Body: requestBody,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}

// DeleteAuthUnlinkSocialIdentities is used to unlink up a social provider account with the specified account
// based on the access token and the social providers user id, the latter is returned with any API call that returns the full
// user profile
// The unlinked account will automatically get removed from your database.
func DeleteAuthUnlinkSocialIdentities(token string, body interface{}) (*httprutils.Response, error) {
	requestBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	tokenHeader := "Bearer " + token

	request := httprutils.Request{
		Method: httprutils.Delete,
		URL:    os.Getenv("DOMAIN") + "/identity/v2/auth/socialidentity",
		Headers: map[string]string{
			"content-Type":  "application/json",
			"Authorization": tokenHeader,
		},
		QueryParams: map[string]string{
			"apikey": os.Getenv("APIKEY"),
		},
		Body: requestBody,
	}

	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}
