package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	lr "github.com/nombiezinja/lr-go-sdk"

	"github.com/joho/godotenv"
	lrauthentication "github.com/nombiezinja/lr-go-sdk/api/authentication"
	"github.com/nombiezinja/lr-go-sdk/lrbody"
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

	user := lrbody.RegistrationUser{
		Email: []lrbody.AuthEmail{
			lrbody.AuthEmail{
				Type:  "Primary",
				Value: "hello123@mailazy.com",
			},
		},
		Password: "password",
	}

	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PostAuthUserRegistrationByEmail(user)
	fmt.Println(res, err)
}
