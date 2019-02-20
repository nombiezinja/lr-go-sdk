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

func NewLoginradius(cfg *Config) (*Loginradius, error) {

	if cfg.ApiKey == "" || cfg.ApiSecret == "" {
		errMsg := "Must initialize Loginradius client with ApiKey and ApiSecret"
		err := lrerror.New("IntializationError", errMsg, errors.New(errMsg))
		return nil, err
	}
	ctx := Context{
		Request: &httprutils.Request{},
	}
	copyCfgToCtx(cfg, &ctx)

	return &Loginradius{
		Context: &ctx,
		Domain:  domain,
	}, nil
}

func copyCfgToCtx(cfg *Config, ctx *Context) {
	ctx.ApiKey = cfg.ApiKey
	ctx.ApiSecret = cfg.ApiSecret
}
