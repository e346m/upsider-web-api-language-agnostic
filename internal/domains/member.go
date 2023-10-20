package domains

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Member struct {
	ID        string
	FullName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	*Organization
}

func generateFromPassword(plainPassword string) ([]byte, error) {
	b := []byte(plainPassword)
	return bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
}

func (m *Member) SetGeneratePassword(plainPassword string) error {
	hashed, err := generateFromPassword(plainPassword)
	if errors.Is(err, bcrypt.ErrPasswordTooLong) {
		return &DomainError{
			Kind:    Validation,
			Message: bcrypt.ErrPasswordTooLong.Error(),
		}
	}

	m.Password = string(hashed)

	return nil
}

func (m *Member) CheckMemberWithPassword(plainPassword string) error {
	b := []byte(plainPassword)
	err := bcrypt.CompareHashAndPassword([]byte(m.Password), b)
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return &DomainError{
			Kind:    NotFound,
			Message: "Member does not exist",
		}
	}

	return nil
}
