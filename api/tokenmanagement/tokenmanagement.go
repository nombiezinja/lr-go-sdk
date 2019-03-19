package tokenmanagement

import (
	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
	lrvalidate "bitbucket.org/nombiezinja/lr-go-sdk/internal/validate"
)

// GetAccessTokenViaFacebook is used to get a LoginRadius access token by sending Facebook’s access token.
// It will be valid for the specific duration of time specified in the response.
// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/native-social-login-api/access-token-via-facebook-token
// Required query parameter: key, fb_access_token
func (lr Loginradius) GetAccessTokenViaFacebook(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"fb_access_token": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["key"] = lr.Client.Context.ApiKey
	req := lr.Client.NewGetReq("/api/v2/access_token/facebook", validatedQueries)

	delete(req.QueryParams, "apiKey")
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
	// 	data := new(AccessToken)
	// 	req, reqErr := CreateRequest("GET", "http://api.loginradius.com/api/v2/access_token/facebook", "")
	// 	if reqErr != nil {
	// 		return *data, reqErr
	// 	}

	// 	q := req.URL.Query()
	// 	q.Add("key", os.Getenv("APIKEY"))
	// 	q.Add("fb_access_token", fbAccessToken)
	// 	req.URL.RawQuery = q.Encode()
	// 	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	// 	err := RunRequest(req, data)
	// 	return *data, err
}

// // GetAccessTokenViaTwitter is used to get a LoginRadius access token by sending Twitter’s access token.
// // It will be valid for the specific duration of time specified in the response.
// func GetAccessTokenViaTwitter(twAccessToken, twTokenSecret string) (AccessToken, error) {
// 	data := new(AccessToken)
// 	req, reqErr := CreateRequest("GET", "http://api.loginradius.com/api/v2/access_token/twitter", "")
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	q := req.URL.Query()
// 	q.Add("key", os.Getenv("APIKEY"))
// 	q.Add("tw_access_token", twAccessToken)
// 	q.Add("tw_token_secret", twTokenSecret)
// 	req.URL.RawQuery = q.Encode()
// 	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

// 	err := RunRequest(req, data)
// 	return *data, err
// }

// // GetAccessTokenViaVkontakte is used to get a LoginRadius access token by sending Vkontakte’s access token.
// // It will be valid for the specific duration of time specified in the response.
// func GetAccessTokenViaVkontakte(vkAccessToken string) (AccessToken, error) {
// 	data := new(AccessToken)
// 	req, reqErr := CreateRequest("GET", "http://api.loginradius.com/api/v2/access_token/vkontakte", "")
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	q := req.URL.Query()
// 	q.Add("key", os.Getenv("APIKEY"))
// 	q.Add("vk_access_token", vkAccessToken)
// 	req.URL.RawQuery = q.Encode()
// 	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

// 	err := RunRequest(req, data)
// 	return *data, err
// }

// // GetRefreshUserProfile is used to get the latest updated
// // social profile data from the user’s social account after authentication.
// // The social profile will be retrieved via oAuth and OpenID protocols.
// // The data is normalized into LoginRadius’ standard data format.
// // This API should be called using the access token retrieved from the refresh access token API.
// func GetRefreshUserProfile(accessToken string) (AuthProfile, error) {
// 	data := new(AuthProfile)
// 	req, reqErr := CreateRequest("GET", "http://api.loginradius.com/api/v2/userprofile/refresh", "")
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	q := req.URL.Query()
// 	q.Add("access_token", accessToken)
// 	req.URL.RawQuery = q.Encode()
// 	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

// 	err := RunRequest(req, data)
// 	return *data, err
// }

// // GetRefreshToken is used to refresh the provider access token after authentication.
// // It will be valid for up to 60 days on LoginRadius depending on the provider. In order
// // to use the access token in other APIs, always refresh the token using this API.
// // Supported Providers : Facebook,Yahoo,Google,Twitter, Linkedin.
// // Contact LoginRadius support team to enable this API.
// func GetRefreshToken(accessToken string) (AccessToken, error) {
// 	data := new(AccessToken)
// 	req, reqErr := CreateRequest("GET", "http://api.loginradius.com/api/v2/access_token/refresh", "")
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	q := req.URL.Query()
// 	q.Add("secret", os.Getenv("APISECRET"))
// 	q.Add("access_token", accessToken)
// 	req.URL.RawQuery = q.Encode()
// 	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

// 	err := RunRequest(req, data)
// 	return *data, err
// }
// package tokenmanagement
