package countries

// Country represents a single country with its name in Polish and English
// Also it contains an ISO 3166-1 alfa-2 code.
type Country struct {
	code        string
	polishName  string
	englishName string
}

// NewCountry creates a new Country object.
func NewCountry(code string, polishName string, englishName string) Country {
	return Country{code: code, polishName: polishName, englishName: englishName}
}

// Code returns an ISO 3166-1 alfa-2 code of the country.
func (c *Country) Code() string {
	return c.code
}

// PolishName returns a Polish name of the country.
func (c *Country) PolishName() string {
	return c.polishName
}

// EnglishName returns an English name of the country.
func (c *Country) EnglishName() string {
	return c.englishName
}
