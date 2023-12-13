package services

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"golangapi/common"
	"golangapi/config"
	"time"

	"github.com/o1egl/paseto"
)

type ITokenHandler interface {
	CreateToken(payload map[string]string, footer map[string]string) (string, error)
	ReadToken(string) (map[string]string, error)
}

type SymmetricalPasetoTokenHandler struct{
	lifetime time.Duration
	issuer string
	privateKey string
}

func NewDefaultSymmetricalPasetoTokenHandler() ITokenHandler {
	return SymmetricalPasetoTokenHandler{
		lifetime: config.Config.Paseto.Lifetime,
		issuer: config.Config.Paseto.Issuer,
		privateKey: config.Config.Paseto.PrivateKey,
	}
}

func (ptc SymmetricalPasetoTokenHandler) CreateToken(payload map[string]string, footer map[string]string) (token string, err error) {
	// The Paset Sign function can panic internally, for example if the provided key wasn't of correct length
	// That is an unwanted behaviour and we want to avoid that!
	defer func() {
		if panicErr := recover(); panicErr != nil {
			token = ""
			err = common.ConvertRecoverToError(panicErr)
		}
	}()
	
	now := time.Now()
	exp := now.Add(ptc.lifetime)

	pToken := paseto.JSONToken{
		Expiration: exp,
		IssuedAt:   now,
		NotBefore:  now,
		Issuer:     ptc.issuer,
	}

	for k, v := range payload {
		pToken.Set(k, v)
	}

	b, err := hex.DecodeString(ptc.privateKey)

	if err != nil {
		fmt.Println("Decode ERROR!")
		return "", err
	}

	privateKey := ed25519.PrivateKey(b)

	// Not needed for symmetrical encryption
	// privateKey.Public()

	// This line can only fail if Sign function returns an error
	// Sign function can return error only if private key is not of type ed25519.PrivateKey
	// Or if either payload or footer cant be converted to bytes
	// We made sure that neither of these cases can ever happen, the private key will always be the desired type
	// The payload and footer are of type map[string]string and that type can always be converted to bytes (marshaled)
	// So we can ignore this error
	return paseto.NewV2().Encrypt(privateKey, pToken, footer)
}

func (ptr SymmetricalPasetoTokenHandler) ReadToken(token string) (map[string]string, error) {
	b, err := hex.DecodeString(ptr.privateKey)

	if err != nil {
		return nil, err
	}

	publicKey := ed25519.PrivateKey(b)

	var data map[string]string

	err = paseto.NewV2().Decrypt(token, publicKey, &data, nil)

	return data, err
}