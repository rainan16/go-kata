// The address package defines the address type and some validation functions
// In Go the symbols starting with a capital letter are exported, visible from
// outside the package scope. In other languages the equivalent is to use a
// keyword like `public` or `export`.

package address

import (
	"fmt"
	"strings"
)

// Address is user defined struct, in Go there are no classes but any user
// defined type can have methods
type Address struct {
	Country string
	City    string
	ZipCode string
	Street  string
	Name    string
}

// Validate returns an error in case the address is not valid.
// This is an example of a method for a type in Go.
//
// There are not enough validations! Can you improve it?
func (addr Address) Validate() error {
	// TODO implement the validation, some is already started can you finish it?
	err := isValidCountry(addr.Country)
	if err != nil {
		return fmt.Errorf("invalid country %q: %w", addr.Country, err)
	}
	if strings.TrimSpace(addr.City) == "" {
		return fmt.Errorf("invalid city %q: empty city", addr.City)
	}
	if strings.TrimSpace(addr.Street) == "" {
		return fmt.Errorf("invalid street %q: empty street", addr.Street)
	}
	if strings.TrimSpace(addr.Name) == "" {
		return fmt.Errorf("invalid name %q: empty name", addr.Name)
	}
	err = isValidZipCode(addr.City, addr.ZipCode)
	if err != nil {
		return fmt.Errorf("invalid zip code %q: %w", addr.ZipCode, err)
	}

	return nil
}

// isValidCountry is not exported and not visible from outside the address package.
func isValidCountry(country string) error {
	// countries don't change very often, maybe a simple way to validate this
	// field is to have a list of countries and check if it is present?
	// let's have a for loop "range" over a slice of strings ([]string)
	// maybe there is some explanation here https://go.dev/doc/effective_go#for
	if strings.TrimSpace(country) == "" {
		return fmt.Errorf("empty country")
	}
	for _, c := range countryList {
		if strings.EqualFold(c, country) {
			return nil
		}
	}
	return fmt.Errorf("unknown country")
}

// countryList is a list of some country for this Kata also not exported
var countryList = []string{
	"Austria",
	"Italy",
	"Switzerland",
	"Germany",
	"Hungary",
	"Slovenia",
	"Slovakia",
	"Czech Republic",
	"Croatia",
}

// isValidZipCode return an error if the zipcode is not valid for the city
func isValidZipCode(city, zipcode string) error {
	// in this case a map (dictionary) can be very useful to see all the
	// possible zipcodes in a city
	if strings.TrimSpace(city) == "" {
		return fmt.Errorf("empty city")
	}
	if strings.TrimSpace(zipcode) == "" {
		return fmt.Errorf("empty zipcode")
	}
	for k, list := range cityZipcode {
		if !strings.EqualFold(k, city) {
			continue
		}
		for _, code := range list {
			if code == zipcode {
				return nil
			}
		}
		return fmt.Errorf("invalid zipcode for city")
	}
	return fmt.Errorf("unknown city")
}

var cityZipcode = map[string][]string{
	"Salzburg": {"5020", "5023"},
	"Vienna":   {"1010", "1020"},
	"Milan":    {"20121", "20122"},
}

// 💡 all these variable are hardcoded, would it be better to load the validation
// data from a file? Maybe a JSON?
