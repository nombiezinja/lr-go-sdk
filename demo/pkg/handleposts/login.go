package handleposts

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	lr "github.com/nombiezinja/lr-go-sdk"
	lrauthentication "github.com/nombiezinja/lr-go-sdk/api/authentication"
	"github.com/nombiezinja/lr-go-sdk/lrbody"
	"github.com/nombiezinja/lr-go-sdk/lrerror"
)

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	var credentials lrbody.EmailLogin
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &credentials)

	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PostAuthLoginByEmail(
		credentials,
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

	// response, err := loginradius.PostAuthLoginByEmail("", "", "", "", "", credentials)

	// if err != nil {
	// 	log.Println(err)
	// 	if lrerr, ok := err.(lrerror.Error); ok {
	// 		errorResponse := lrerror.ErrorResponse{}
	// 		error := json.Unmarshal([]byte(lrerr.OrigErr().Error()), &errorResponse)
	// 		if error != nil {
	// 			w.WriteHeader(500)
	// 			w.Write([]byte("Something went wrong on our end."))
	// 		} else {
	// 			w.WriteHeader(errorResponse.ErrorCode)
	// 			w.Write([]byte(errorResponse.Description))
	// 		}
	// 	} else {
	// 		w.WriteHeader(500)
	// 		w.Write([]byte("Unknown error"))
	// 	}
	// 	return
	// }
	// // data, err := json.Marshal(response.Body)
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(response.Body))
}
