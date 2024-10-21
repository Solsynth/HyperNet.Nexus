package sec

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

type JwtReader struct {
	key *rsa.PublicKey
}

func NewJwtReader(fp string) (*JwtReader, error) {
	privateKeyBytes, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privateKeyBytes)
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

func ReadJwt[T jwt.Claims](v *JwtReader, in string) (T, error) {
	var out T
	token, err := jwt.Parse(in, func(token *jwt.Token) (interface{}, error) {
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
