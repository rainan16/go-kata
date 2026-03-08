package user

import (
	"fmt"
	address "go-kata/task-2"
	"testing"
)

// TestUserLogin is a function that tests the user logins.
func TestUserLogin(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		// variables can be declared and defined with the := operator
		// the type is inferred
		name, email := "Bob", "bob's email.com"
		// we don't need the new user for this test so we can assign it to
		// the blank identifier https://go.dev/doc/effective_go#blank
		user, err := New(name, email, address.Address{}, address.Address{})
		if err != nil {
			t.Logf("user should be valid with name: %s and email: %s", name, email)
			t.Fail()
		}

		user.Loginer = LoginBasic{username: "Herbert", password: "y"}
		session, err := user.Loginer.Login("xtoken")
		if err != nil {
			t.Logf("token invalid: %s", err)
			t.Fail()
		}
		if session != "OK" {
			t.Logf("session invalid: %s", session)
			t.Fail()
		}
	})
}

func TestLoginBasic(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		l := LoginBasic{username: "Herbert", password: "secret"}
		if _, err := l.Login("token"); err != nil {
			t.Fatalf("expected valid login: %v", err)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		tests := []LoginBasic{
			{username: "", password: "x"},
			{username: "ab", password: "x"},
			{username: "abc", password: ""},
		}
		for i, l := range tests {
			if _, err := l.Login("token"); err == nil {
				t.Fatalf("expected error for case %d", i)
			}
		}
	})
}

func TestLoginInterface(t *testing.T) {
	var l Loginer = testLogin{hash: "ok"}
	if _, err := l.Login("bad"); err == nil {
		t.Fatalf("expected error for bad token")
	}
	if session, err := l.Login("ok"); err != nil || session == "" {
		t.Fatalf("expected session on success, got session=%q err=%v", session, err)
	}
}

// testLogin implements the Loginer interface  for tests.
// 💡can you create other implementations?
type testLogin struct {
	hash string
}

func (tl testLogin) Login(token string) (string, error) {
	if token == tl.hash {
		return "session", nil
	}
	return "", fmt.Errorf("login error, check credentials")
}
