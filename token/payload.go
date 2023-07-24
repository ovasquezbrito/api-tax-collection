package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token in invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(email string, duration time.Duration) (*Payload, error) {
	fmt.Println(duration)
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	fmt.Println(payload.ExpiredAt)
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
