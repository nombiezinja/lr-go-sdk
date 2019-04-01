package handlegets

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nombiezinja/lr-go-sdk/demo/pkg/template"
)

func Verify(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := template.Render(w, r, "emailverification.page", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
