package handlegets

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Verify(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// token := r.URL.Query().Get("vtoken")
	// response, err := loginradius.GetAuthVerifyEmail(token, "", "")

	// successMsg := "Email verification successful"
	// errMessage := ""
	// if err != nil {
	// 	if lrerr, ok := err.(lrerror.Error); ok {
	// 		log.Println(lrerr.Error())

	// 		errorResponse := lrerror.ErrorResponse{}
	// 		error := json.Unmarshal([]byte(lrerr.OrigErr().Error()), &errorResponse)
	// 		if error != nil {
	// 			w.WriteHeader(500)
	// 			errMessage = "Something went wrong on our end"
	// 		} else {
	// 			w.WriteHeader(errorResponse.ErrorCode)
	// 			errMessage = errorResponse.Description
	// 		}
	// 	} else {
	// 		log.Println(err)
	// 		w.WriteHeader(500)
	// 		w.Write([]byte("Unknown error"))
	// 	}

	// 	err = template.Render(w, r, "index.page", map[string]interface{}{"Error": errMessage})
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		http.Error(w, "Internal Server Error", 500)
	// 	}
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// data, _ := lrjson.DynamicUnmarshal(response.Body)

	// if !data["IsPosted"].(bool) {
	// 	successMsg = "Not posted!"
	// }

	// err = template.Render(w, r, "index.page", map[string]interface{}{"Success": successMsg})
	// if err != nil {
	// 	log.Println(err.Error())
	// 	http.Error(w, "Internal Server Error", 500)
	// }

}
