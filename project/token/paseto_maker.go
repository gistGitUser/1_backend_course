package token

import (
	"aidanwoods.dev/go-paseto"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/chacha20poly1305"
	"time"
)

var ErrInvalidPasetoKeySize = fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)

type PasetoMaker struct {
	paseto       *paseto.V4SymmetricKey
	token        *paseto.Token
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (*PasetoMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, ErrInvalidPasetoKeySize
	}

	key := paseto.NewV4SymmetricKey()
	token := paseto.NewToken()
	maker := &PasetoMaker{
		paseto:       &key,
		token:        &token,
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func NewPASETOPayload(username string, duration time.Duration) (*paseto.Token, error) {
	token := paseto.NewToken()

	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	token.SetIssuedAt(time.Now())
	token.SetString("username", username)
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(duration))

	token.SetString("user-id", tokenID.String())
	return &token, err
}

func (maker *PasetoMaker) VerifyToken(token string) (*paseto.Token, error) {
	parser := paseto.NewParser()
	T, err := parser.ParseV4Local(*maker.paseto, token, nil)
	if err != nil {
		
	}
	return T, err
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPASETOPayload(username, duration)
	if err != nil {
		return "", err
	}

	return payload.V4Encrypt(*maker.paseto, nil), err
}
