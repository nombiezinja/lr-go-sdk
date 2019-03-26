// Package lrjson contains functions for unmarshalling JSON responses returned by RESTful API
package lrjson

import (
	"encoding/json"
	"log"

	"github.com/loginradius"
)

// DynamicUnmarshal takes stringified json and unmarshals it into a map with string keys and
// interface{} values. This is recommended for parsing LoginRadius API responses over individual
// functions unmarshalling JSON objects into pre-written structs.

// Golang's strict typing means the latter solution is a better practice in theory, but the former
// solution is recommended for usage with LoginRadius api end points so as to ensure the long-term
// integrity of response handling - i.e. unmarshalling into a prewritten type will throw an error in // case of unexpected field data types, and will quietly do the best it can when there is a mismatch // between incoming JSON object and the destination struct. The former results in unwanted fragility, // and the latter results in data being potentially mis-captured.

// (Note: another alternative is unmarshalling the response as it is read from the body:

// decoder := json.NewDecoder(res.Body)
// decoder.DisallowUnknownFields()
// return decoder.Decode(destinationStruct)

// DisallowUnknownFields was made available in Go 1.10 and will causes the Decoder to return an
// error when the destination is a struct and the input contains object keys which do not match any
// non-ignored, exported fields in the destination. Though this would solve the issue of data not being captured, it would increase the fragility of the API client during new endpoint updates/releases.
//)

// In addition to these reasonings, benchmarks between the two alternatives have been provided in
// benchmark-test.go in this package. The performance difference between the dynamic unmarshalling
// solution and the struct specific unmarshalling solution is about 70000 ns/op
// - insignificant enough to disregard, especially considering the amount of work required to maintain
// pre-written structs and response-specific code.

func DynamicUnmarshal(data string) (map[string]interface{}, error) {
	var unmarshalled = make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &unmarshalled)
	if err != nil {
		log.Println(err)
	}
	return unmarshalled, nil
}

// A sample function for unmarshalling JSON response into pre-written struct, preserved for
// benchmarking purposes
func UnmarshalGetManageAccountProfilesByEmail(data string) (loginradius.AuthProfile, error) {
	authProfile := loginradius.AuthProfile{}

	error := json.Unmarshal([]byte(data), &authProfile)
	if error != nil {
		return authProfile, error
	}
	return authProfile, nil
}
