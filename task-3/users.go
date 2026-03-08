package user

import (
	"fmt"
	address "go-kata/task-2"
)

// User of the system.
// The User needs to have both a billing and shipping address, and be able to
// login.
type User struct {
	Name  string
	Email string
	// we need to have a billing and a shipping address, here is the Billing one
	// you need to add another one
	Billing address.Address
	Shipping address.Address

	// 💡and a Login method would not go amiss, there should be an interface we
	// can embed... or implement
	Loginer Loginer
}

// New returns a new User. Go has not special functions (constructors) to
// create new values, we just use functions. The name of the function usually
// contains the New word.
func New(name, email string, billig address.Address, shipping address.Address) (User, error) {
	newUser := User{
		Name:    name,
		Email:   email,
		Billing: billig,
		Shipping: shipping,
	}
	return newUser, nil
}

// String implements the Stringer interface for the User type.
// https://pkg.go.dev/fmt#Stringer
func (u User) String() string {
	return fmt.Sprintf("Name: %s, Email: %s", u.Name, u.Email)
}
