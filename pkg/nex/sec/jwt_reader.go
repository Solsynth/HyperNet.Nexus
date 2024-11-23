package sec

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math/big"
	"os"
)

type JwtReader struct {
	key *rsa.PublicKey
}

func NewJwtReader(fp string) (*JwtReader, error) {
	rawKey, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(rawKey)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	anyPk, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pk, ok := anyPk.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return &JwtReader{
		key: pk,
	}, nil
}

// ReadJwt is the helper method to help me validate and parse jwt.
// To use it, pass the initialized jwt reader which contains a public key.
// And pass the token string and a pointer struct (you must initialize it, which it cannot be nil) of your claims
func ReadJwt[T jwt.Claims](v *JwtReader, in string, out T) (T, error) {
	token, err := jwt.ParseWithClaims(in, out, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return v.key, nil
	})
	if err != nil {
		return out, err
	} else if !token.Valid {
		return out, fmt.Errorf("token is not valid")
	}

	if claims, ok := token.Claims.(T); ok {
		return claims, nil
	} else {
		return out, err
	}
}

func (v *JwtReader) BuildJwk(kid string) map[string]any {
	encodeBigInt := func(i *big.Int) string {
		return base64.RawURLEncoding.EncodeToString(i.Bytes())
	}

	return map[string]any{
		"kid": kid,
		"kty": "RSA",
		"use": "sig",
		"alg": "RS256",
		"n":   encodeBigInt(v.key.N),
		"e":   encodeBigInt(big.NewInt(int64(v.key.E))),
	}
}
