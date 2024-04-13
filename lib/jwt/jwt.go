package jwt

import "context"

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope    string `json:"scope"`
	Nickname string `json:"nickname"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}
