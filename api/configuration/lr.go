// lrconfiguration package contains methods for calling the Configuration and Infrastructure APIs,
// which are used to view configurations and information around the customer account.
package lrconfiguration

import (
	lr "github.com/nombiezinja/lr-go-sdk"
)

// Embed Loginradius struct
type Loginradius struct {
	Client *lr.Loginradius
}
