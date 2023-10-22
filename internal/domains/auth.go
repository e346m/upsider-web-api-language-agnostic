package domains

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func generateFromPassword(plainPassword string) ([]byte, error) {
	b := []byte(plainPassword)
	return bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
}

func SetGeneratePassword(plainPassword string, m *Member) error {
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

func CheckMemberWithPassword(plainPassword, encriptedPassword string) error {
	b := []byte(plainPassword)
	err := bcrypt.CompareHashAndPassword([]byte(encriptedPassword), b)
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return &DomainError{
			Kind:    NotFound,
			Message: "Member does not exist",
		}
	}

	return nil
}
