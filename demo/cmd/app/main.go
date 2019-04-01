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
	router.ServeFiles("/assets/*filepath", http.Dir(filepath.Join(cwd, "../../ui/assets")))
	router.GET("/", handlegets.Index)
	router.GET("/emailverification", handlegets.Verify)
	router.GET("/register/verify/email", handleposts.Verify)
	router.GET("/screen", handlegets.Screen)
	router.POST("/register", handleposts.Signup)
	router.POST("/login/email", handleposts.Login)
	router.GET("/profile", handlegets.Profile)
	router.POST("/profile", handleposts.Profile)
	router.GET("/roles/get", handlegets.Roles)
	router.GET("/roles", handlegets.Roles)

	http.ListenAndServe(":3000", router)
}
