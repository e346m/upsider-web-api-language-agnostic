package domains

import (
	"math/rand"
	"testing"
	"time"
)

func TestSetGeneratePassword(t *testing.T) {
	t.Run("password more than 73 length is expected to fail",
		func(t *testing.T) {
			m := &Member{
				Password: "",
			}
			password := generateRandomString(73)
			err := SetGeneratePassword(password, m)

			if err == nil {
				t.Fatalf("test should fail")
			}
		},
	)
	t.Run(
		"password less than 73 length is expected to success",

		func(t *testing.T) {
			m := &Member{
				Password: "",
			}
			password := generateRandomString(72)
			err := SetGeneratePassword(password, m)
			if err != nil {
				t.Fatalf("test should success but %v happens", err.Error())
			}
		},
	)
}

func TestCheckMemberWithPassword(t *testing.T) {
	t.Run("invalid password is expected to fail",
		func(t *testing.T) {
			password := generateRandomString(7)
			invalid := ""
			hashed, _ := generateFromPassword(password)
			err := CheckMemberWithPassword(invalid, string(hashed))

			if err == nil {
				t.Fatalf("test should fail")
			}
		},
	)
	t.Run(
		"valid password is expected to success",

		func(t *testing.T) {
			password := generateRandomString(7)
			hashed, _ := generateFromPassword(password)
			err := CheckMemberWithPassword(password, string(hashed))
			if err != nil {
				t.Fatalf("test should success but %v happens", err.Error())
			}
		},
	)
}

// TODO move util and test it as well
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}
