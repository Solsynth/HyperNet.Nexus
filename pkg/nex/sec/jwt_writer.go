package sec

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

type JwtWriter struct {
	key *rsa.PrivateKey
}

func NewJwtWriter(fp string) (*JwtWriter, error) {
	rawKey, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(rawKey)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	anyPk, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pk, ok := anyPk.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA private key")
	}

	return &JwtWriter{
		key: pk,
	}, nil
}

func WriteJwt[T jwt.Claims](v *JwtWriter, in T) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, in)
	return token.SignedString(v.key)
}
