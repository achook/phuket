package countries

type Country struct {
	code        string
	polishName  string
	englishName string
}

func newCountry(code string, polishName string, englishName string) Country {
	return Country{code: code, polishName: polishName, englishName: englishName}
}

func (c *Country) Code() string {
	return c.code
}

func (c *Country) PolishName() string {
	return c.polishName
}

func (c *Country) EnglishName() string {
	return c.englishName
}
