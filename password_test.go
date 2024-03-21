package password

import (
	"testing"
)

func FuzzPasswordChecking(f *testing.F) {
	passwords := []string{"password", " ", "!23%"}
	for _, p := range passwords {
		f.Add(p)
	}
	f.Fuzz(func(t *testing.T, password string) {
		encrypted, err := HashPassword(password)
		if err != nil {
			t.Errorf("HashPassword failed on '%s'", password)
		}
		check := CheckPassword(password, encrypted)
		if !check {
			t.Errorf("CheckPassword failed on '%s'", password)
		}
	})
}
