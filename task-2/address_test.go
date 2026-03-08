package address

import "testing"

// TestAddressIsValid tests the address validation function IsValid()
// The Go standard library includes a testing framework in the testing package.
//
// This test is quite incomplete, can you improve it?
func TestAddressIsValid(t *testing.T) {
	t.Run("isValidCountry", func(t *testing.T) {
		tests := []struct {
			name    string
			country string
			wantErr bool
		}{
			{name: "valid exact", country: "Austria", wantErr: false},
			{name: "valid case-insensitive", country: "austria", wantErr: false},
			{name: "empty", country: "", wantErr: true},
			{name: "unknown", country: "Atlantis", wantErr: true},
		}
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				err := isValidCountry(tc.country)
				if (err != nil) != tc.wantErr {
					t.Fatalf("isValidCountry(%q) err=%v wantErr=%v", tc.country, err, tc.wantErr)
				}
			})
		}
	})

	t.Run("isValidZipCode", func(t *testing.T) {
		tests := []struct {
			name    string
			city    string
			zip     string
			wantErr bool
		}{
			{name: "valid exact", city: "Salzburg", zip: "5020", wantErr: false},
			{name: "valid case-insensitive city", city: "salzburg", zip: "5023", wantErr: false},
			{name: "empty city", city: "", zip: "5020", wantErr: true},
			{name: "empty zip", city: "Salzburg", zip: "", wantErr: true},
			{name: "unknown city", city: "Gotham", zip: "1234", wantErr: true},
			{name: "invalid zip", city: "Vienna", zip: "9999", wantErr: true},
		}
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				err := isValidZipCode(tc.city, tc.zip)
				if (err != nil) != tc.wantErr {
					t.Fatalf("isValidZipCode(%q,%q) err=%v wantErr=%v", tc.city, tc.zip, err, tc.wantErr)
				}
			})
		}
	})

	t.Run("Validate", func(t *testing.T) {
		valid := Address{
			Country: "Austria",
			City:    "Salzburg",
			ZipCode: "5020",
			Street:  "Hauptstrasse 1",
			Name:    "Alice",
		}
		if err := valid.Validate(); err != nil {
			t.Fatalf("valid address should pass: %v", err)
		}

		invalid := []Address{
			{Country: "", City: "Salzburg", ZipCode: "5020", Street: "x", Name: "y"},
			{Country: "Austria", City: "", ZipCode: "5020", Street: "x", Name: "y"},
			{Country: "Austria", City: "Salzburg", ZipCode: "", Street: "x", Name: "y"},
			{Country: "Austria", City: "Salzburg", ZipCode: "9999", Street: "x", Name: "y"},
			{Country: "Austria", City: "Salzburg", ZipCode: "5020", Street: "", Name: "y"},
			{Country: "Austria", City: "Salzburg", ZipCode: "5020", Street: "x", Name: ""},
		}
		for i, addr := range invalid {
			if err := addr.Validate(); err == nil {
				t.Fatalf("invalid address %d should fail", i)
			}
		}
	})
}
