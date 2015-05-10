package password

import (
	"testing"
)

func TestPassword(t *testing.T) {
	some_password := "some_password"
	if !Check(some_password, Generate(some_password)) {
		t.Errorf("Can't validate password with generated hash!")
	}
}
