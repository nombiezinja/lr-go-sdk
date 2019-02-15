// The lrbody package holds the structs to be encoded as the body in POST and PUT calls
// These structs are meant to serve as convenient measures assisting API calls provided by the Loginradius
// Go SDK; all functions in this SDK takes interface{} as the body, but initiating your
// data in the appropriate struct and passing in place of the body when calling the SDK functions
// will ensure the parameters submitted are correctly formatted and named for the LoginRadius APIs
// The usage of the structs in this package is optional and provided for convenience only
// For examples please see INSERTDOC

// Note: these structs provide reference only, and do not include optional parameters

package lrbody

// Used by PostAuthUserRegistrationByEmail
type AuthEmail struct {
	Type  string `json:"Type"`
	Value string `json:"Value"`
}

// Used by PostAuthUserRegistrationByEmail
type RegistrationUser struct {
	Email    []AuthEmail `json:"Email"`
	Password string      `json:"Password"`
}

// Used by PostAuthLoginByEmail
type EmailLogin struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

// Used by PostAuthLoginByUsername
type UsernameLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Used by PostAuthAddEmail
type AddEmail struct {
	Type  string `json:"type"`
	Email string `json:"email"`
}

// Used by PostAuthForgotPassword, PutResendEmailVerification
type EmailStr struct {
	Email string `json:"email"`
}

// Used by PutAuthUpdateSecurityQuestionByEmail
type SecurityQuestionAnswer struct {
	SecurityAnswer string `json:"2acec20722394dc3bd6362ef27df824e"`
}

// used by PutAuthChangePassword
type ChangePassword struct {
	OldPw string `json:"oldpassword"`
	NewPw string `json:"newpassword"`
}

// used by PutAuthLinkSocialIdentities
type LinkeSocialIds struct {
	CandidateToken string `json:"candidatetoken"`
}

// used by PutAuthResetPasswordByResetToken
type ResetPw struct {
	ResetToken string `json:"resettoken"`
	Password   string `json:"password"`
}

// used by PutAuthResetPasswordByOTP
type ResetPwOtp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Otp      string `json:"otp"`
}

// used by PutAuthResetPasswordBySecurityAnswerAndEmail
type ResetPwSecurityQuestionEmail struct {
	Email          string `json:"email"`
	SecurityAnswer string `json:"securityanswer"`
	Password       string `json:"Password"`
}

// used by PutAuthResetPasswordBySecurityAnswerAndPhone
type ResetPwSecurityQuestionPhone struct {
	Phone          string `json:"phone"`
	SecurityAnswer string `json:"securityanswer"`
	Password       string `json:"password"`
}

// used by PutAuthResetPasswordBySecurityAnswerAndUsername
type ResetPwSecurityQuestionUsername struct {
	Username       string `json:"username"`
	SecurityAnswer string `json:"securityanswer"`
	Password       string `json:"password"`
}

type AuthUsername struct {
	Username string `json:"username"`
}
