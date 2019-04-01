package handlegets

import (
	"log"
	"net/http"

	"bitbucket.org/nombiezinja/lr-go-sdk-demo/pkg/template"
	"github.com/julienschmidt/httprouter"
)

func Profile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := template.Render(w, r, "profile.page", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
