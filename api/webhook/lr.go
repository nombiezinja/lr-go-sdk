// Package webhook contains methods for calling LoginRadius webhook APIs.
package webhook

import (
	lr "github.com/nombiezinja/lr-go-sdk"
)

// Embed Loginradius struct
type Loginradius struct {
	Client *lr.Loginradius
}
