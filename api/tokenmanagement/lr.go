package tokenmanagement

import (
	lr "bitbucket.org/nombiezinja/lr-go-sdk"
)

// Embed Loginradius struct
type Loginradius struct {
	Client *lr.Loginradius
}