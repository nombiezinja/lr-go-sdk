package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/nombiezinja/lr-go-sdk/demo/pkg/handledeletes"
	"github.com/nombiezinja/lr-go-sdk/demo/pkg/handlegets"
	"github.com/nombiezinja/lr-go-sdk/demo/pkg/handleposts"
	"github.com/nombiezinja/lr-go-sdk/demo/pkg/handleputs"
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
	// router.HandleFunc("/login/passwordless", pwlessHandler).Methods("GET")
	// router.HandleFunc("/login/passwordless/auth", verifyLoginHandler).Methods("GET")
	// router.HandleFunc("/customobj", updateCustomObjHandler).Methods("PUT")
	// router.HandleFunc("/customobj", deleteCustomObjHandler).Methods("DELETE")
	// router.HandleFunc("/mfa/google", mfaResetGoogleHandler).Methods("DELETE")
	// router.HandleFunc("/mfa/validate", mfaAccessTokenHandler).Methods("GET")
	// router.HandleFunc("/mfa/google/enable", mfaAccessTokenAuthHandler).Methods("PUT")
	// router.HandleFunc("/roles", createRoleHandler).Methods("POST")
	// router.HandleFunc("/roles", deleteRoleHandler).Methods("DELETE")
	// router.HandleFunc("/roles", assignRoleHandler).Methods("PUT")

	router.POST("/index", handleposts.Index)
	router.GET("/api/register/verify/email", handleposts.Verify)
	router.POST("/api/register", handleposts.Signup)
	router.POST("/api/login/email", handleposts.Login)
	router.POST("/api/profile", handleposts.Profile)
	router.GET("/api/roles/get", handlegets.Roles)
	router.GET("/api/roles", handlegets.Roles)
	router.POST("/api/forgotpassword", handleposts.ForgotPassword)
	router.PUT("/api/login/resetpassword", handleputs.ResetPassword)
	router.POST("/api/mfa/login/email", handleposts.MfaLogin)
	router.PUT("/api/mfa/google/auth", handleputs.MfaGoogleAuth)
	router.PUT("/api/profile/changepassword", handleputs.ChangePassword)
	router.PUT("/api/profile/setpassword", handleputs.SetPassword)
	router.PUT("/api/profile/update", handleputs.UpdateProfile)
	router.POST("/api/customobj", handleposts.CustomObject)
	router.GET("/api/customobj", handlegets.CustomObject)
	router.PUT("/api/customobj", handleputs.CustomObject)
	router.DELETE("/api/customobj", handledeletes.CustomObject)

	// if not found look for a static file
	static := httprouter.New()
	static.ServeFiles("/*filepath", http.Dir(filepath.Join(cwd, "../../ui/assets")))
	router.NotFound = static

	http.ListenAndServe(":3000", router)
}
