// Package phoneauthentication contains methods for calling the LoginRadius phone authentication APIs
package phoneauthentication

import (
	lr "github.com/nombiezinja/lr-go-sdk"
)

// Embed Loginradius struct
type Loginradius struct {
	Client *lr.Loginradius
}
