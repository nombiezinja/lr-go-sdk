package lrintegrationtest

import (
	"os"
	"testing"

	lr "bitbucket.org/nombiezinja/lr-go-sdk"
	lrconfiguration "bitbucket.org/nombiezinja/lr-go-sdk/api/configuration"
)

func TestGetConfiguration(t *testing.T) {
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	_, err := lrconfiguration.Loginradius(lrconfiguration.Loginradius{lrclient}).GetConfiguration()
	if err != nil {
		t.Errorf("Error calling GetConfiguration: %v", err)
	}
}

// func TestGetServerTime(t *testing.T) {
// 	fmt.Println("Starting test TestGetServerTime")
// 	PresetLoginRadiusTestEnv()
// 	_, err := GetServerTime("")
// 	if err != nil {
// 		t.Errorf("Error getting server time")
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Test complete")
// }
// func TestGetGenerateSottAPI(t *testing.T) {
// 	fmt.Println("Starting test TestGetGenerateSottAPI")
// 	PresetLoginRadiusTestEnv()
// 	_, err := GetGenerateSottAPI("")
// 	if err != nil {
// 		t.Errorf("Error generating SOTT")
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Test complete")
// }
// func TestGetActiveSessionDetails(t *testing.T) {
// 	fmt.Println("Starting test TestGetActiveSessionDetails")
// 	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
// 	defer teardownTestCase(t)
// 	_, err := GetActiveSessionDetails(accessToken)
// 	if err != nil {
// 		t.Errorf("Error getting active session details")
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Test complete")
// }
