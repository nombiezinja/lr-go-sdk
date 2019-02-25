package loginradius

import (
	"errors"

	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	"bitbucket.org/nombiezinja/lr-go-sdk/lrerror"
)

// NewAuthGetRequest constructs the request for Auth api end points requiring
// Access token in the header and ApiKey in query param
func (lr Loginradius) NewAuthGetReqWithAccessToken(path string, queries ...map[string]string) (*httprutils.Request, error) {

	if lr.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	request := &httprutils.Request{
		Method: httprutils.Get,
		URL:    lr.Domain + path,
		Headers: map[string]string{
			"content-Type":  "application/x-www-form-urlencoded",
			"Authorization": "Bearer " + lr.Context.Token,
		},
		QueryParams: map[string]string{
			"apiKey": lr.Context.ApiKey,
		},
	}

	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}

	return request, nil
}

func (lr Loginradius) NewAuthGetReq(path string, queries map[string]string) *httprutils.Request {
	return &httprutils.Request{
		Method:      httprutils.Get,
		URL:         lr.Domain + path,
		Headers:     httprutils.URLEncodedHeader,
		QueryParams: queries,
	}
}

func (lr Loginradius) NewAuthPostReqWithToken(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {

	if lr.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	encodedBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := &httprutils.Request{
		Method: httprutils.Post,
		URL:    lr.Domain + path,
		Headers: map[string]string{
			"content-Type":  "application/json",
			"Authorization": "Bearer " + lr.Context.Token,
		},
		QueryParams: map[string]string{
			"apiKey": lr.Context.ApiKey,
		},
		Body: encodedBody,
	}

	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}

	return request, nil
}

func (lr Loginradius) NewAuthPostReq(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {

	encodedBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := &httprutils.Request{
		Method: httprutils.Post,
		URL:    lr.Domain + path,
		Headers: map[string]string{
			"content-Type": "application/json",
		},
		QueryParams: map[string]string{
			"apiKey": lr.Context.ApiKey,
		},
		Body: encodedBody,
	}

	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}

	return request, nil
}
