// The lrbody package holds the structs to be encoded as the body in POST and PUT calls
// These structs are meant to serve as convenient measures to API calls provided by the Loginradius
// Go SDK; all functions in this SDK takes interface{} as the body, but initiating your
// data in the appropriate struct and passing in place of the body when calling the SDK functions
// will ensure the parameters submitted are correctly formatted and named for the LoginRadius APIs
// The usage of the structs in this package is optional and provided for convenience only
// For examples please see INSERTDOC

package lrbody

// Used by PostManageAccountCreate
type EmailArray []struct {
	Type  string
	Value string
}

// Used by PostManageAccountCreate
type AccountCreate struct {
	Email    EmailArray
	Password string
}

// Used by PostManageEmailVerificationToken, PostManageForgotPasswordToken
type Email struct {
	Email string `json:"Email"`
}

// Used by PostManageForgotPasswordToken
type Username struct {
	Username string `json:"EMail"`
}

// Used by PutManageAccountUpdateSecurityQuestionConfig
type AccountSecurityQuestion struct {
	Securityquestionanswer accountSecurityQA `json:"securityquestionanswer"`
}

// The security question is identified by a random string key in the LoginRadius database
// You can retrieve this key with a call to GetConfiguration, and replace the
// json tag value with your question string
type accountSecurityQA struct {
	QuestionID string `json:"2acec20722394dc3bd6362ef27df824e"`
}
