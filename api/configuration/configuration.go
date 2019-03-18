package lrconfiguration

import "bitbucket.org/nombiezinja/lr-go-sdk/httprutils"

// GetConfiguration is used to get the configurations which are set in the
// LoginRadius Dashboard for a particular LoginRadius site/environment.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/configuration/get-configurations
// Required parameter: apikey
func (lr Loginradius) GetConfiguration() (*httprutils.Response, error) {
	req := lr.Client.NewGetReq("")
	req.URL = "https://config.lrcontent.com/ciam/appinfo"
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
	// 	data := new(Configurations)
	// 	req, reqErr := CreateRequest("GET", "https://config.lrcontent.com/ciam/appinfo", "")
	// 	if reqErr != nil {
	// 		return *data, reqErr
	// 	}

	// 	q := req.URL.Query()
	// 	q.Add("apikey", os.Getenv("APIKEY"))
	// 	req.URL.RawQuery = q.Encode()
	// 	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	// 	err := RunRequest(req, data)
	// 	return *data, err
}

// // GetServerTime allows you to query your LoginRadius account for basic server information
// // and server time information which is useful when generating an SOTT token.
// func GetServerTime(timeDifference string) (ServerTime, error) {
// 	data := new(ServerTime)
// 	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/serverinfo", "")
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	q := req.URL.Query()
// 	q.Add("apikey", os.Getenv("APIKEY"))
// 	q.Add("timedifference", timeDifference)
// 	req.URL.RawQuery = q.Encode()
// 	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

// 	err := RunRequest(req, data)
// 	return *data, err
// }

// // GetGenerateSottAPI allows you to generate SOTT with a given expiration time.
// func GetGenerateSottAPI(timeDifference string) (SOTT, error) {
// 	data := new(SOTT)
// 	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/manage/account/sott", "")
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	q := req.URL.Query()
// 	q.Add("timedifference", timeDifference)
// 	req.URL.RawQuery = q.Encode()
// 	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
// 	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
// 	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

// 	err := RunRequest(req, data)
// 	return *data, err
// }

// // GetActiveSessionDetails is used to get all active sessions by Access Token.
// func GetActiveSessionDetails(accessToken string) (ActiveSession, error) {
// 	data := new(ActiveSession)
// 	req, reqErr := CreateRequest("GET", "http://api.loginradius.com/api/v2/access_token/activesession", "")
// 	if reqErr != nil {
// 		return *data, reqErr
// 	}

// 	q := req.URL.Query()
// 	q.Add("token", accessToken)
// 	q.Add("key", os.Getenv("APIKEY"))
// 	q.Add("secret", os.Getenv("APISECRET"))
// 	req.URL.RawQuery = q.Encode()
// 	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

// 	err := RunRequest(req, data)
// 	return *data, err
// }
