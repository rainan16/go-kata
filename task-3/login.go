package user

import "fmt"

// Loginer is an interface with a login method.
// Yes the name sucks, but it is idiomatic to call interfaces with a -er suffix.
type Loginer interface {
	// Login allows the type implementing this method to login by providing
	// an opaque token.
	// In case of success a nil error and a session token is returned.
	// In case of error the session in not to valid.
	Login(token string) (session string, err error)
}

// TODO how can we make our user implement this interface?
// And what if we have different login mechanisms?
// Can we have more types implementing the same interface and assign one at user creation?


type LoginBasic struct {
	username 	string
	password	string
}

func (login LoginBasic) Login(token string) (session string, err error) {
	
	if login.username == "" {
		return "", fmt.Errorf("empty username")
	}

	if login.password == "" {
		return "", fmt.Errorf("empty password")
	}

	if len(login.username) < 3 {
		return "", fmt.Errorf("username too short")
	}

	return "OK", nil
} 