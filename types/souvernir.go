package types

import "phuket/countries"

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
