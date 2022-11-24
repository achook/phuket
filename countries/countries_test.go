package countries

import (
	"testing"
)

func TestCountry_Code(t *testing.T) {
	code := "PL"
	country := newCountry(code, "Polska", "Poland")

	if country.Code() != code {
		t.Errorf("Country.Code() = %s; want %s", country.Code(), code)
	}
}
