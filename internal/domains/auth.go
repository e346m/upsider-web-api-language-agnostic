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

func CheckMemberWithPassword(plainPassword, encryptedPassword string) error {
	b := []byte(plainPassword)
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), b)
	if err != nil {
		return &DomainError{
			Kind:    NotFound,
			Message: "Member does not exist",
		}
	}

	return nil
}
