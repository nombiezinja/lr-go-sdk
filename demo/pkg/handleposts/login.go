package handleposts

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// w.Header().Set("Content-Type", "application/json")

	// var credentials lrbody.EmailLogin
	// b, _ := ioutil.ReadAll(r.Body)
	// json.Unmarshal(b, &credentials)
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
