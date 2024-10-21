package sec

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtClaims struct {
	jwt.RegisteredClaims

	// Nexus Standard
	Session  int           `json:"sed"`
	CacheTTL time.Duration `json:"ttl,omitempty"`

	// OIDC Standard
	Name  string `json:"name,omitempty"`
	Nick  string `json:"preferred_username,omitempty"`
	Email string `json:"email,omitempty"`

	// OAuth2 Standard
	AuthorizedParties string `json:"azp,omitempty"`
	Nonce             string `json:"nonce,omitempty"`

	// The usage of this token
	// Can be access_token, refresh_token or id_token
	Type string `json:"typ"`
}
