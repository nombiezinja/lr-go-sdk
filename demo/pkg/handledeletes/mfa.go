package handledeletes

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	lr "github.com/nombiezinja/lr-go-sdk"
	"github.com/nombiezinja/lr-go-sdk/api/mfa"
	"github.com/nombiezinja/lr-go-sdk/lrerror"
)

// w.Header().Set("Content-Type", "application/json")

// auth := r.URL.Query().Get("auth")
// var resetAuth ResetGoogleAuth
// b, _ := ioutil.ReadAll(r.Body)
// json.Unmarshal(b, &resetAuth)
// profile, err := loginradius.DeleteMFAResetGoogleAuthenticatorByToken(auth, resetAuth)
// if err != nil {
// 	w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
// 	w.Write([]byte(err.Error()))
// 	return
// }

// w.WriteHeader(http.StatusOK)
// data, _ := json.Marshal(profile)
// w.Write(data)

func MfaGoogleReset(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var errors string
	respCode := 200

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	token := r.Header.Get("Authorization")[7:]
	lrclient, err := lr.NewLoginradius(
		&cfg,
		map[string]string{"token": token},
	)
	if err != nil {
		errors = errors + err.(lrerror.Error).OrigErr().Error()
		respCode = 500
	}

	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).DeleteMFAResetGoogleAuthenticatorByToken()
	if err != nil {
		errors = errors + err.(lrerror.Error).OrigErr().Error()
		respCode = 500
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(respCode)
	if errors != "" {
		log.Printf(errors)
		w.Write([]byte(errors))
		return
	}
	w.Write([]byte(res.Body))
}
