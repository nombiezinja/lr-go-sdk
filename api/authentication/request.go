package lrauthentication

import (
	"errors"

	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	"bitbucket.org/nombiezinja/lr-go-sdk/lrerror"
)

// NewAuthGetRequest constructs the request for Auth api end points requiring
// Access token in the header and ApiKey in query param
func (lr Loginradius) NewAuthGetReqWithAccessToken(path string) (*httprutils.Request, error) {

	if lr.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	return &httprutils.Request{
		Method: httprutils.Get,
		URL:    lr.Domain + path,
		Headers: map[string]string{
			"content-Type":  "application/x-www-form-urlencoded",
			"Authorization": "Bearer " + lr.Context.Token,
		},
		QueryParams: map[string]string{
			"apiKey": lr.Context.ApiKey,
		},
	}, nil
}

func (lr Loginradius) NewAuthGetReq(path string, queries map[string]string) *httprutils.Request {
	return &httprutils.Request{
		Method:      httprutils.Get,
		URL:         lr.Domain + path,
		Headers:     httprutils.URLEncodedHeader,
		QueryParams: queries,
	}
}
