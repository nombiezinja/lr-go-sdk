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
	// router.HandleFunc("/mfa/login/email", mfaLoginHandler).Methods("POST")
	// router.HandleFunc("/mfa/google/auth", mfaLoginAuthHandler).Methods("PUT")
	// router.HandleFunc("/login/passwordless", pwlessHandler).Methods("GET")
	// router.HandleFunc("/login/passwordless/auth", verifyLoginHandler).Methods("GET")
	// router.HandleFunc("/forgotpassword", forgotPasswordHandler).Methods("POST")
	// router.HandleFunc("/login/resetpassword", resetPasswordByEmailHandler).Methods("PUT")
	// router.HandleFunc("/resetpassword", resetPasswordHandler).Methods("PUT")
	// router.HandleFunc("/profile/changepassword", changePasswordHandler).Methods("PUT")
	// router.HandleFunc("/profile/setpassword", setPasswordHandler).Methods("PUT")
	// router.HandleFunc("/profile/update", updateAccountHandler).Methods("PUT")
	// router.HandleFunc("/customobj", createCustomObjHandler).Methods("POST")
	// router.HandleFunc("/customobj", updateCustomObjHandler).Methods("PUT")
	// router.HandleFunc("/customobj", deleteCustomObjHandler).Methods("DELETE")
	// router.HandleFunc("/customobj", getCustomObjHandler).Methods("GET")
	// router.HandleFunc("/mfa/google", mfaResetGoogleHandler).Methods("DELETE")
	// router.HandleFunc("/mfa/validate", mfaAccessTokenHandler).Methods("GET")
	// router.HandleFunc("/mfa/google/enable", mfaAccessTokenAuthHandler).Methods("PUT")
	// router.HandleFunc("/roles", createRoleHandler).Methods("POST")
	// router.HandleFunc("/roles", deleteRoleHandler).Methods("DELETE")
	// router.HandleFunc("/roles", assignRoleHandler).Methods("PUT")
	// router.GET("/", s.index())
	// router.GET("/api/", s.api())
	router.GET("/api/register/verify/email", handleposts.Verify)
	router.POST("/api/register", handleposts.Signup)
	router.POST("/api/login/email", handleposts.Login)
	router.POST("/api/profile", handleposts.Profile)
	router.GET("/api/roles/get", handlegets.Roles)
	router.GET("/api/roles", handlegets.Roles)

	// if not found look for a static file
	static := httprouter.New()
	static.ServeFiles("/*filepath", http.Dir(filepath.Join(cwd, "../../ui/assets")))
	router.NotFound = static

	// router.GET("/", handlegets.Index)
	// router.GET("/emailverification", handlegets.Verify)
	// router.GET("/screen", handlegets.Screen)
	// router.GET("/profile", handlegets.Profile)

	http.ListenAndServe(":3000", router)
}
