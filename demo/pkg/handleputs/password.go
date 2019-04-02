package handleputs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	lr "github.com/nombiezinja/lr-go-sdk"
	lrauthentication "github.com/nombiezinja/lr-go-sdk/api/authentication"
	"github.com/nombiezinja/lr-go-sdk/lrerror"
)

// var resetPassInfo ResetPasswordRequest
// b, _ := ioutil.ReadAll(r.Body)
// json.Unmarshal(b, &resetPassInfo)
// resetPass := ResetPasswordEmail{resetPassInfo.ResetToken, resetPassInfo.Password, "", ""}
// profile, err := loginradius.PutAuthResetPasswordByResetToken(resetPass)
// if err != nil {
// 	w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
// 	w.Write([]byte(err.Error()))
// 	return
// }

// w.WriteHeader(http.StatusOK)
// data, _ := json.Marshal(profile)
// w.Write(data)
func ResetPassword(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var errors string
	respCode := 200

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, err := lr.NewLoginradius(&cfg)
	if err != nil {
		errors = errors + err.(lrerror.Error).OrigErr().Error()
		respCode = 500
	}

	data := struct {
		ResetToken string
		Password   string
		// 	Email      string
	}{}

	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &data)

	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).
		PutAuthResetPasswordByResetToken(data)

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
