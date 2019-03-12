package lrsocial

import (
	"errors"

	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	lrvalidate "bitbucket.org/nombiezinja/lr-go-sdk/internal/validate"
	"bitbucket.org/nombiezinja/lr-go-sdk/lrerror"
)

// GetSocialAccessToken Is used to translate the Request Token
// returned during authentication into an Access Token that can be used with other API calls.
// Required query parameters: token - string (LoginRadius request token); secret - LoginRadius API secret
// For more information on the LoginRadius request token: https://www.loginradius.com/docs/infrastructure-and-security/loginradius-tokens#loginradius-request-token-expiration-15-mins-
func (lr Loginradius) GetSocialAccessToken(requestToken string) (*httprutils.Response, error) {
	req := lr.Client.NewGetReq("/api/v2/access_token", map[string]string{
		"token":  requestToken,
		"secret": lr.Client.Context.ApiSecret,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialTokenValidate validates access_token, if valid then returns a response with its expiry otherwise error.
// Required query params: key - string ; secret - string; access_token - string
func (lr Loginradius) GetSocialTokenValidate() (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}
	req := lr.Client.NewGetReq("/api/v2/access_token/validate", map[string]string{
		"key":          lr.Client.Context.ApiKey,
		"secret":       lr.Client.Context.ApiSecret,
		"access_token": lr.Client.Context.Token,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialTokenInvalidate validates access_token, if valid then returns a response with its expiry otherwise error.
// Required query params: key - string ; secret - string; access_token - string
// Optional Parameters: preventRefresh - string (takes true or false)
func (lr Loginradius) GetSocialTokenInvalidate(queries ...interface{}) (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/access_token/invalidate", map[string]string{
		"key":          lr.Client.Context.ApiKey,
		"secret":       lr.Client.Context.ApiSecret,
		"access_token": lr.Client.Context.Token,
	})

	for _, arg := range queries {
		allowedQueries := map[string]bool{"preventRefresh": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			req.QueryParams[k] = v
		}
	}

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialAlbum returns the photo albums associated with the passed in access tokens Social Profile.
// Supported Providers: Facebook, Google, Live, Vkontakte.
// Required query parameters: access_token - string
// Please ensure your LoginRadius site has requested for access to this end point for the social provider from your users
// Returns an array
func (lr Loginradius) GetSocialAlbum() (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/album", map[string]string{
		"access_token": lr.Client.Context.Token,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialAudio is used to get audio files data from the user’s social account.
// Supported Providers: Live, Vkontakte
// Required query parameters: access_token - string
// Please ensure your LoginRadius site has requested for access to this end point for the social provider from your users
// Returns an array
func (lr Loginradius) GetSocialAudio() (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/audio", map[string]string{
		"access_token": lr.Client.Context.Token,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialCheckin is used to get check Ins data from the user’s social account.
// Supported Providers: Facebook, Foursquare, Vkontakte
// Required query parameters: access_token - string
// Please ensure your LoginRadius site has requested for access to this end point for the social provider from your users
// Returns an array
func (lr Loginradius) GetSocialCheckin() (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/checkin", map[string]string{
		"access_token": lr.Client.Context.Token,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialCompany is used to get the followed companies data from the user’s social account.
// Supported Providers: Facebook, LinkedIn
// Required query parameters: access_token - string
// Please ensure your LoginRadius site has requested for access to this end point for the social provider from your users
func (lr Loginradius) GetSocialCompany() (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/company", map[string]string{
		"access_token": lr.Client.Context.Token,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialContact is used to get contacts/friends/connections data from the user’s social account.
// This is one of the APIs that makes up the LoginRadius Friend Invite System.
// The data will normalized into LoginRadius’ standard data format.
// This API requires setting permissions in your LoginRadius Dashboard.
// Note: Facebook restricts access to the list of friends that is returned.
// When using the Contacts API with Facebook you will only receive friends that have accepted some permissions with your app.
// Supported Providers: Facebook, Foursquare, Google, LinkedIn, Live, Twitter, Vkontakte, Yahoo
// Required query parameter: access_token - string
func (lr Loginradius) GetSocialContact() (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/contact", map[string]string{
		"access_token": lr.Client.Context.Token,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialEvent is used to get the event data from the user’s social account.
// Supported Providers: Facebook, Live
// Required query parameter: access_token - string
func (lr Loginradius) GetSocialEvent() (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/event", map[string]string{
		"access_token": lr.Client.Context.Token,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialFollowing is used to get the following user list from the user’s social account.
// Supported Providers: Twitter
// Required query parameters: access_token - string
// Please ensure your LoginRadius site has requested for access to this end point for the social provider from your users
func (lr Loginradius) GetSocialFollowing() (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/following", map[string]string{
		"access_token": lr.Client.Context.Token,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialGroup is used to get group data from the user’s social account.
// Supported Providers: Facebook, Vkontakte
// Required query parameters: access_token - string
// Please ensure your LoginRadius site has requested for access to this end point for the social provider from your users
func (lr Loginradius) GetSocialGroup() (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/group", map[string]string{
		"access_token": lr.Client.Context.Token,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialLike is used to get likes data from the user’s social account.
// Supported Providers: Facebook
// Required query parameters: access_token - string
// Please ensure your LoginRadius site has requested for access to this end point for the social provider from your users
func (lr Loginradius) GetSocialLike() (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/like", map[string]string{
		"access_token": lr.Client.Context.Token,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialMention is used to get mention data from the user’s social account.
// Supported Providers: Twitter
// Required query parameters: access_token - string
// Please ensure your LoginRadius site has requested for access to this end point for the social provider from your users
func (lr Loginradius) GetSocialMention() (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/mention", map[string]string{
		"access_token": lr.Client.Context.Token,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialStatusPost is used to update the status on the user’s wall.
// Supported Providers: Facebook, Twitter, LinkedIn
// Required query parameters: url - string; title - string; imageurl-string; status-string; caption - string; description - string;
// GET & POST Social Status  API work the same way except the API method is different
func (lr Loginradius) GetSocialStatusPost(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"url": true, "title": true, "imageurl": true, "status": true, "caption": true, "description": true,
	}

	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	validatedQueries["access_token"] = lr.Client.Context.Token

	request := lr.Client.NewGetReq("/api/v2/status/js", validatedQueries)

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetSocialUserProfile is used to get social profile data from the user’s social account after authentication.
// Supported Providers: All
// Required query parameters: access_token - string
func (lr Loginradius) GetSocialUserProfile() (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/userprofile", map[string]string{
		"access_token": lr.Client.Context.Token,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialPage is used to get the page data from the user’s social account.
// Supported Providers: Facebook, LinkedIn
// Required query parameters: access_token - string; pagename - string
// Please ensure that your app has been reviewed and you have permission to request this information from your users
func (lr Loginradius) GetSocialPage(pagename string) (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/page", map[string]string{
		"access_token": lr.Client.Context.Token, "pagename": pagename,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialPhoto is used to get the photo data from the user’s social account.
// Supported Providers: Facebook, Foursquare, Google, Live, Vkontakte
// Required query parameters: access_token - string; albumid - string
// Please ensure that your app has been reviewed and you have permission to request this information from your users
func (lr Loginradius) GetSocialPhoto(albumid string) (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/photo", map[string]string{
		"access_token": lr.Client.Context.Token, "albumid": albumid,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialPost is used to get post message data from the user’s social account.
// Supported Providers: Facebook
// Required query parameters: access_token - string
// Please ensure that your app has been reviewed and you have permission to request this information from your users
func (lr Loginradius) GetSocialPost() (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/post", map[string]string{
		"access_token": lr.Client.Context.Token,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialStatusFetch is used to get the status messages from the user’s social account.
// Supported Providers: Facebook, LinkedIn, Twitter, Vkontakte
// Required query parameters: access_token - string
// Please ensure that your app has been reviewed and you have permission to request this information from your users
func (lr Loginradius) GetSocialStatus() (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/status", map[string]string{
		"access_token": lr.Client.Context.Token,
	})

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetSocialVideo is used to get video files data from the user’s social account.
// Supported Providers: Facebook, Google, Live, Vkontakte
// Required query parameters: next_cursor - string
// Please ensure that your app has been reviewed and you have permission to request this information from your users
func (lr Loginradius) GetSocialVideo(queries ...interface{}) (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/video", map[string]string{
		"access_token": lr.Client.Context.Token,
	})

	for _, arg := range queries {
		allowedQueries := map[string]bool{"nextcursor": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			req.QueryParams[k] = v
		}
	}

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}
