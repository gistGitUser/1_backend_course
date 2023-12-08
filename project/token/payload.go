package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var ErrExpiredToken = errors.New("token has expired")
var ErrInvalidToken = errors.New("token is invalid")

type Payload struct {
	Username  string    `json:"username"`
	ID        uuid.UUID `json:"id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	//чтобы можно было использовать нужно встроить это поле
	jwt.RegisteredClaims
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	//в целом можно было бы использовать дефолтные значения
	//но лучше сразу указать как использовать кастомные поля
	return &Payload{
		Username: username,
		ID:       tokenID,
		//ExpiredAt: time.Now(),
		//IssuedAt:  time.Now().Add(duration),
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}, err

	//return &jwt.MapClaims{
	//	"username":   username,
	//	"id":         tokenID,
	//	"issued_at":  time.Now(),
	//	"expired_at": time.Now().Add(duration),
	//}, nil

}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
