package loginradius

import (
	"errors"

	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	"bitbucket.org/nombiezinja/lr-go-sdk/lrerror"
)

var domain = "https://api.loginradius.com"

type Loginradius struct {
	Context *Context
	Domain  string
}

type Config struct {
	ApiKey    string
	ApiSecret string
}

type Context struct {
	ApiKey    string
	ApiSecret string

	Request *httprutils.Request
}

func NewLoginradius(cfg *Config, optionalArgs ...map[string]string) (*Loginradius, error) {

	if cfg.ApiKey == "" || cfg.ApiSecret == "" {
		errMsg := "Must initialize Loginradius client with ApiKey and ApiSecret"
		err := lrerror.New("IntializationError", errMsg, errors.New(errMsg))
		return nil, err
	}

	ctx := Context{
		Request: &httprutils.Request{},
	}

	// If an access token is passed on initiation, create Auth Bearer token header
	for _, arg := range optionalArgs {
		if arg["token"] != "" {
			tokenHeader := "Bearer " + arg["token"]
			ctx.Request.Headers = map[string]string{"Authorization": tokenHeader}
		}
	}

	ctx.ApiKey = cfg.ApiKey
	ctx.ApiSecret = cfg.ApiSecret

	return &Loginradius{
		Context: &ctx,
		Domain:  domain,
	}, nil
}
