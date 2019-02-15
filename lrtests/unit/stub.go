package unittest

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"bitbucket.org/nombiezinja/lr-go-sdk/httprutils"
)

// type requestMaker interface {
// 	testRequest() (*httprutils.Response, error)
// }

func initTestServer(path string, resp httprutils.Response) *httptest.Server {
	fmt.Println("Run")
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("uri %v, path %v", r.RequestURI, path)
		if r.RequestURI != path {
			fmt.Println("run")
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		fmt.Println(path)
		fmt.Printf("%+v", resp)
		w.WriteHeader(resp.StatusCode)
		w.Write([]byte(resp.Body))
	}))
}

// type requestTest struct {
// 	path     string
// 	response string
// 	expected *httprutils.Response
// }

// // Suite should set up test server, call method with params, then compare output
// func SuiteTest(t *testing.T, testCase requestTest, expected *httprutils.Response, req requestMaker) {
// 	stub := initTestServer(
// 		testCase.path,
// 		testCase.response,
// 	)
// 	defer stub.Close()
// 	res, _ := req.testRequest()
// 	if res != expected {
// 		t.Errorf("Expected %v, received %v", testCase.expected, res)
// 	}
// }
