package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	lr "bitbucket.org/nombiezinja/lr-go-sdk"
	lrauthentication "bitbucket.org/nombiezinja/lr-go-sdk/api/authentication"

	"github.com/joho/godotenv"
)

func main() {
	cwd, _ := os.Getwd()

	err := godotenv.Load(
		filepath.Join(cwd, "../../config/secret.env"),
		filepath.Join(cwd, "../../config/public.env"),
	)

	if err != nil {
		log.Fatal("Error loading env files, please configure your secret.env and public.env.")
	}

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}
	lrclient, _ := lr.NewLoginradius(&cfg)

	// lrclient, _ := lr.NewLoginradius(&cfg, map[string]string{"token": "9c3208ae-2848-4ac5-baef-41dd4103e263"})
	// loginradius := lrauth.Loginradius{initLr}

	// queries := map[string]string{}
	// body := lrbody.RegistrationUser{}
	// response, err := lrauth.Loginradius(loginradius).PostAuthUserRegistrationByEmail(queries, body)
	// res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthReadProfilesByToken()
	// res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthSendWelcomeEmail(map[string]string{})
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthSocialIdentity()

	if err != nil {
		fmt.Printf("%+v", err)
	}

	fmt.Printf("%+v", res)
	// fmt.Printf("%+v", lrclient.Context)

}
