package loginradius

import (
	"errors"

	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	"bitbucket.org/nombiezinja/lr-go-sdk/lrerror"
)

// NewGetRequest constructs the request for Auth api end points requiring
// Access token in the header and ApiKey in query param
func (lr Loginradius) NewGetReqWithToken(path string, queries ...map[string]string) (*httprutils.Request, error) {

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

func (lr Loginradius) NewGetReq(path string, queries ...map[string]string) *httprutils.Request {
	request := &httprutils.Request{
		Method:      httprutils.Get,
		URL:         lr.Domain + path,
		Headers:     httprutils.URLEncodedHeader,
		QueryParams: map[string]string{},
	}
	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}
	return request
}

func (lr Loginradius) NewPostReqWithToken(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {

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

func (lr Loginradius) NewPostReq(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {

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

func (lr Loginradius) NewPutReq(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {

	encodedBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := &httprutils.Request{
		Method: httprutils.Put,
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

func (lr Loginradius) NewPutReqWithToken(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {

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
		Method: httprutils.Put,
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

func (lr Loginradius) NewDeleteReq(path string, body ...interface{}) *httprutils.Request {
	if len(body) != 0 {
		encoded, err := httprutils.EncodeBody(body[0])
		if err != nil {
			return nil
		}
		return &httprutils.Request{
			Method:  httprutils.Delete,
			URL:     lr.Domain + path,
			Headers: httprutils.URLEncodedHeader,
			Body:    encoded,
		}
	} else {
		return &httprutils.Request{
			Method:  httprutils.Delete,
			URL:     lr.Domain + path,
			Headers: httprutils.URLEncodedHeader,
		}
	}
}

func (lr Loginradius) NewDeleteReqWithToken(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {

	if lr.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

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
		Method: httprutils.Delete,
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

func (lr Loginradius) AddApiCredentialsToReqHeader(req *httprutils.Request) {
	delete(req.QueryParams, "apiKey")
	req.Headers["X-LoginRadius-ApiKey"] = lr.Context.ApiKey
	req.Headers["X-LoginRadius-ApiSecret"] = lr.Context.ApiSecret
}

func (lr Loginradius) AddApiKeyToReqHeader(req *httprutils.Request) {
	req.Headers["X-LoginRadius-ApiKey"] = lr.Context.ApiKey
}

func (lr Loginradius) NormalizeApiKey(req *httprutils.Request) {
	delete(req.QueryParams, "apiKey")
	req.QueryParams["apikey"] = lr.Context.ApiKey
}
