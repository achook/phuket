package types

import (
	"phuket/countries"
)

type Souvenir struct {
	id int

	name        string
	description string

	year int

	country *countries.Country
	region  string

	from string
	trip string
}

func NewSouvenir(name string, description string, year int,
	countryCode string, region string, from string, trip string) (Souvenir, error) {

	country, err := countries.FromCode(countryCode)
	if err != nil {
		return Souvenir{}, err
	}

	return Souvenir{
		name:        name,
		description: description,
		year:        year,
		country:     &country,
		region:      region,
		from:        from,
		trip:        trip,
	}, nil
}
