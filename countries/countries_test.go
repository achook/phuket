package countries

import (
	"testing"
)

func TestCountry_Code(t *testing.T) {
	code := "PL"
	country := NewCountry(code, "Polska", "Poland")

	if country.Code() != code {
		t.Errorf("Country.Code() = %s; want %s", country.Code(), code)
	}
}

func TestCountry_EnglishName(t *testing.T) {
	englishName := "United States of America"
	country := NewCountry("US", "Stany Zjednoczone", englishName)

	if country.EnglishName() != englishName {
		t.Errorf("Country.EnglishName() = %s; want %s", country.EnglishName(), englishName)
	}
}

func TestCountry_PolishName(t *testing.T) {
	polishName := "Zjednoczone Emiraty Arabskie"
	country := NewCountry("AE", polishName, "United Arab Emirates")

	if country.PolishName() != polishName {
		t.Errorf("Country.PolishName() = %s; want %s", country.PolishName(), polishName)
	}
}
