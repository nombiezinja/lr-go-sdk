package handleposts

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Email struct {
	Type  string `json:"Type"`
	Value string `json:"Value"`
}

type User struct {
	Email                []Email `json:"Email"`
	Password             string  `json:"Password"`
	PasswordConfirmation string  `json:"Password_confirmation"`
}

func Signup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// w.Header().Set("Content-Type", "application/json")

	// var user User
	// b, _ := ioutil.ReadAll(r.Body)
	// json.Unmarshal(b, &user)

	// response, err := loginradius.PostAuthUserRegistrationByEmail(os.Getenv("BASEURL")+"verify", "", "", user)
	// if err != nil {
	// 	// Returns generic lrerror
	// 	if lrerr, ok := err.(lrerror.Error); ok {
	// 		log.Println(lrerr.Error())

	// 		// Check to see whether lrerror returned is a loginradius server returned request error
	// 		errorResponse := lrerror.ErrorResponse{}
	// 		error := json.Unmarshal([]byte(lrerr.OrigErr().Error()), &errorResponse)

	// 		// Safely handle non Loginradius-API return related errors by logging the error and display opaque message to front end
	// 		// If error is caused by a LoginRadius API end point returned error response, display the discription to the front end
	// 		// and set the header with the returned error code
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

	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(response)
	// w.Write([]byte(data))
}
