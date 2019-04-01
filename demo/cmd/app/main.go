package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/nombiezinja/lr-go-sdk/demo/pkg/handlegets"
	"github.com/nombiezinja/lr-go-sdk/demo/pkg/handleposts"
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

	router := httprouter.New()

	router.ServeFiles("/assets/*filepath", http.Dir(filepath.Join(cwd, "../../ui/assets")))
	router.GET("/", handlegets.Index)
	router.GET("/emailverification", handlegets.RenderVerify)
	router.GET("/register/verify/email", handlegets.Verify)
	router.GET("/screen", handlegets.Screen)
	router.POST("/register", handleposts.Signup)
	router.POST("/login/email", handleposts.Login)
	router.GET("/profile", handleposts.Profile)

	http.ListenAndServe(":3000", router)
}
