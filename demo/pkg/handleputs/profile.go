package handleputs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	lr "github.com/nombiezinja/lr-go-sdk"
	account "github.com/nombiezinja/lr-go-sdk/api/account"
	"github.com/nombiezinja/lr-go-sdk/lrerror"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
		FirstName string
		LastName  string
		About     string
	}{}

	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &data)

	res, err := account.Loginradius(account.Loginradius{lrclient}).
		PutManageAccountUpdate(
			r.URL.Query().Get("uid"),
			data,
		)

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
