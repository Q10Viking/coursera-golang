package user_test

import (
	"testing"

	"q10viking.github.io/coursera-go/user"
)

func TestUser(t *testing.T) {
	if user.Msg() != "hello World" {
		t.Fatal("Something wrong")
	}
}
