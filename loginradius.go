package loginradius

import (
	"errors"

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
	Token     string
}

func NewLoginradius(cfg *Config, optionalArgs ...map[string]string) (*Loginradius, error) {

	if cfg.ApiKey == "" || cfg.ApiSecret == "" {
		errMsg := "Must initialize Loginradius client with ApiKey and ApiSecret"
		err := lrerror.New("IntializationError", errMsg, errors.New(errMsg))
		return nil, err
	}

	ctx := Context{
		ApiKey:    cfg.ApiKey,
		ApiSecret: cfg.ApiSecret,
	}

	// If an access token is passed on initiation, create Auth Bearer token header
	for _, arg := range optionalArgs {
		if arg["token"] != "" {
			ctx.Token = arg["token"]
		} else {
			ctx.Token = ""
		}
	}

	return &Loginradius{
		Context: &ctx,
		Domain:  domain,
	}, nil
}
