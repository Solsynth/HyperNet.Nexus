package sec

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type InternalTokenWriter struct {
	pk ed25519.PrivateKey
}

func NewInternalTokenWriter(fp string) (*InternalTokenWriter, error) {
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

	pk, ok := anyPk.(ed25519.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("not an Ed25519 private key")
	}

	return &InternalTokenWriter{
		pk: pk,
	}, nil
}

func (v *InternalTokenWriter) WriteUserInfoJwt(in UserInfo, audiences ...string) (string, error) {
	rawData := base64.StdEncoding.EncodeToString(in.Encode())
	claims := jwt.RegisteredClaims{
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		Audience:  audiences,
		Issuer:    "nexus",
		Subject:   rawData,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	return token.SignedString(v.pk)
}

type InternalTokenReader struct {
	pk ed25519.PublicKey
}

func NewInternalTokenReader(fp string) (*InternalTokenReader, error) {
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

	pk, ok := anyPk.(ed25519.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an Ed25519 public key")
	}

	return &InternalTokenReader{
		pk: pk,
	}, nil
}

func (v *InternalTokenReader) ReadUserInfoJwt(in string) (*UserInfo, error) {
	token, err := jwt.ParseWithClaims(in, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return v.pk, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}
	rawData, err := base64.StdEncoding.DecodeString(claims.Subject)
	if err != nil {
		return nil, err
	}
	info, err := NewUserInfoFromBytes(rawData)
	return &info, err
}
