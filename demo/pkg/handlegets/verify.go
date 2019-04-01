package handlegets

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	lr "github.com/nombiezinja/lr-go-sdk"
	lrauthentication "github.com/nombiezinja/lr-go-sdk/api/authentication"
	"github.com/nombiezinja/lr-go-sdk/demo/pkg/template"
	"github.com/nombiezinja/lr-go-sdk/lrerror"
)

func Verify(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token := r.URL.Query().Get("verification_token")
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

	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthVerifyEmail(
		map[string]string{"verificationtoken": token},
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

func RenderVerify(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := template.Render(w, r, "emailverification.page", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
