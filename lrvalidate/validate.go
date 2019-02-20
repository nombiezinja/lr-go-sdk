package lrvalidate

import (
	"errors"

	"bitbucket.org/nombiezinja/lr-go-sdk/lrerror"
)

// Validate verifies an interface{} only contains keys belonging to the allowed map of keys
// It returns a map[string]string if type assertion is successful and all params are validated
// It returns an error if the submitted params cannot be type asserted into map[string]string, or if the submitted params contain keys that is not included in the allowed map[string]bool
func Validate(allowed map[string]bool, params interface{}) (map[string]string, error) {
	asserted, ok := params.(map[string]string)

	if !ok {
		err := lrerror.New("ValidationError", "Error validating params - params type error", errors.New("Error validating params"))
		return nil, err
	}

	for k, _ := range asserted {
		if !allowed[k] {
			err := lrerror.New("ValidationError", "Error validating params - invalid params submitted, please double check", errors.New("Error validating params"))
			return nil, err
		}
	}
	return asserted, nil
}
