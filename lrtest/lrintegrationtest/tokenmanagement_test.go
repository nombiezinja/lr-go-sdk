package lrintegrationtest

import (
	"os"
	"testing"

	lr "bitbucket.org/nombiezinja/lr-go-sdk"
	"bitbucket.org/nombiezinja/lr-go-sdk/api/tokenmanagement"
	lrjson "bitbucket.org/nombiezinja/lr-go-sdk/lrjson"
)

func TestGetAccessTokenViaFacebook(t *testing.T) {
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, err := lr.NewLoginradius(&cfg)

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	res, err := tokenmanagement.Loginradius(tokenmanagement.Loginradius{lrclient}).GetAccessTokenViaFacebook(
		map[string]string{"fb_access_token": "abcd1234abcd"},
	)

	if err != nil {
		t.Errorf("Error calling GetAccessTokenViaFacebook: %v", err)
	}

	tokens, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || tokens["access_token"].(string) == "" {
		t.Errorf("ERror returned from GetAccessTokenViaFacebook: %v, %v", err, tokens)
	}
}

// func TestGetAccessTokenViaTwitter(t *testing.T) {
// 	fmt.Println("Starting test TestGetAccessTokenViaTwitter")
// 	PresetLoginRadiusTestEnv()
// 	session, err := GetAccessTokenViaTwitter(os.Getenv("TWITTERTOKEN"), os.Getenv("TWITTERSECRET"))
// 	if err != nil || session.AccessToken == "" {
// 		t.Errorf("Error retrieving twitter token")
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Test complete")
// }

// func TestGetAccessTokenViaVkontakte(t *testing.T) {
// 	fmt.Println("Starting test TestGetAccessTokenViaVkontakte")
// 	PresetLoginRadiusTestEnv()
// 	session, err := GetAccessTokenViaVkontakte(os.Getenv("VKONTAKTETOKEN"))
// 	if err != nil || session.AccessToken == "" {
// 		t.Errorf("Error retrieving facebook token")
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Test complete")
// }

// func TestGetRefreshUserProfile(t *testing.T) {
// 	fmt.Println("Starting test TestGetRefreshUserProfile")
// 	PresetLoginRadiusTestEnv()
// 	session, err := GetAccessTokenViaFacebook(os.Getenv("FACEBOOKTOKEN"))
// 	if err != nil || session.AccessToken == "" {
// 		t.Errorf("Error retrieving facebook token")
// 		fmt.Println(err)
// 	}
// 	_, err2 := GetRefreshUserProfile(session.AccessToken)
// 	if err2 != nil {
// 		t.Errorf("Error refreshing profile")
// 		fmt.Println(err2)
// 	}
// 	fmt.Println("Test complete")
// }

// func TestGetRefreshToken(t *testing.T) {
// 	fmt.Println("Starting test TestGetRefreshToken")
// 	PresetLoginRadiusTestEnv()
// 	session, err := GetAccessTokenViaTwitter(os.Getenv("TWITTERTOKEN"), os.Getenv("TWITTERSECRET"))
// 	if err != nil || session.AccessToken == "" {
// 		t.Errorf("Error retrieving twitter token")
// 		fmt.Println(err)
// 	}
// 	_, err2 := GetRefreshToken(session.AccessToken)
// 	if err2 != nil {
// 		t.Errorf("Error refreshing token")
// 		fmt.Println(err2)
// 	}
// 	fmt.Println("Test complete")
// }
